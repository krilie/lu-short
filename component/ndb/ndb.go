package ndb

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-gorp/gorp/v3"
	_ "github.com/go-sql-driver/mysql"
	"lu-short/component/dbmigrate"
	"lu-short/component/ncfg"
	"strings"
	"sync"
	"time"
)

// IDb interface that wrap of db operation
type IDb interface {
	gorp.SqlExecutor
}

// NDb struct wrap of sql db config and start close
type NDb struct {
	cfg struct {
		ConnStr         string
		MaxOpenConn     int
		MaxIdleConn     int
		ConnMaxLeftTime int
	}
	onceStartDb sync.Once
	onceStopDb  sync.Once
	sqlxDb      IDb // gorp.SqlExecutor
}

func (n *NDb) AddTable(iTable interface{}, name string) *gorp.TableMap {
	return n.sqlxDb.(*gorp.DbMap).AddTableWithName(iTable, name)
}

func NewNDb(cfg *ncfg.NConfig) *NDb {

	dbCfg := cfg.GetDbCfg()

	db := &NDb{}
	db.cfg.ConnStr = dbCfg.ConnStr
	db.cfg.ConnMaxLeftTime = dbCfg.ConnMaxLeftTime
	db.cfg.MaxIdleConn = dbCfg.MaxIdleConn
	db.cfg.MaxOpenConn = dbCfg.MaxOpenConn
	db.Start()
	return db
}

func (ndb *NDb) Ping() error {
	return ndb.sqlxDb.(*gorp.DbMap).Db.Ping()
}

func (ndb *NDb) Start() {
	ndb.onceStartDb.Do(func() {

		// 数据库迁移
		ndb.MigrationDb()

		db, err := sql.Open("mysql", ndb.cfg.ConnStr)
		if err != nil {
			panic(err)
		}
		db.SetConnMaxIdleTime(time.Second * time.Duration(ndb.cfg.ConnMaxLeftTime))
		db.SetConnMaxLifetime(time.Hour * 6)
		db.SetMaxIdleConns(ndb.cfg.MaxIdleConn)
		db.SetMaxOpenConns(ndb.cfg.MaxOpenConn)
		err = db.Ping()
		if err != nil {
			panic(err)
		}
		ndb.sqlxDb = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	})
}

func (ndb *NDb) CloseDb() {
	ndb.onceStopDb.Do(func() {
		err := ndb.sqlxDb.(*gorp.DbMap).Db.Close()
		if err != nil {
			panic(err)
		}
	})
}

// GetDb get db before use
// if in transaction this function will return tx set on context or. return sqlx.db
func (ndb *NDb) GetDb(ctx context.Context) IDb {
	txDb := GetTxDbFromCtx(ctx)
	if txDb != nil {
		return txDb
	}
	return ndb.sqlxDb
}

func (ndb *NDb) Exec(ctx context.Context, sql string, args ...interface{}) (rowsAffected int64, err error) {
	execContext, err := ndb.GetDb(ctx).WithContext(ctx).Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	return execContext.RowsAffected()
}

func (ndb *NDb) Get(ctx context.Context, data interface{}, sql string, args ...interface{}) error {
	return ndb.GetDb(ctx).SelectOne(data, sql, args...)
}

func (ndb *NDb) Select(ctx context.Context, data interface{}, sql string, args ...interface{}) error {
	_, err := ndb.GetDb(ctx).Select(data, sql, args...)
	return err
}

// GetGlobalDb get db before use
// get ori db of global which can begin a new trans
func (ndb *NDb) GetGlobalDb(ctx context.Context) IDb {
	return ndb.sqlxDb
}

// WithTrans start trans with db on context
func WithTrans(ctx context.Context, trans func(ctx context.Context) error, onNewTrans ...bool) (err error) {
	// 环境变量
	isOnNewTrans := len(onNewTrans) >= 1 && onNewTrans[0]
	globalDb := MustGetGlobalDbFromCtx(ctx)
	oldTransDb := GetTxDbFromCtx(ctx)
	// 执行新事务
	var doTransOnNewSession = func() error {
		// 准备环境
		newTransDb, err := globalDb.(*gorp.DbMap).Begin() // 新的事务对象
		if err != nil {
			return errors.New("begin tx err " + err.Error())
		}
		newTransCtx := ForceNewDbValuesCtx(ctx, globalDb, newTransDb) // 新的上下文 覆盖旧的上下文
		// panic or err 回滚
		panicked := true
		defer func() {
			if panicked || err != nil {
				err = errors.New("err or panic on trans " + err.Error())
				err2 := newTransDb.Rollback()
				if err2 != nil {
					err = errors.New("err on trans rollback " + err.Error() + err2.Error())
				}
			} else {
				err := newTransDb.Commit()
				if err != nil {
					panic(errors.New("err on commit " + err.Error()))
				}
			}
		}()
		// 执行事务代码并返回
		err = trans(newTransCtx)
		panicked = false
		return err
	}
	// 执行代码
	if isOnNewTrans {
		return doTransOnNewSession()
	} else {
		if oldTransDb == nil {
			return doTransOnNewSession()
		} else {
			// panic or err 回滚
			panicked := true
			defer func() {
				if panicked || err != nil {
					err = errors.New("err or panic on trans (inner trans no commit) " + err.Error())
				}
			}()
			// 执行事务代码并返回
			err = trans(ctx)
			panicked = false
			return err
		}
	}
}

func (ndb *NDb) MigrationDb() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	// 数据库迁移 ?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai
	// multiStatements=true
	var dbName = GetDbNameFromConnectStr(ndb.cfg.ConnStr)
	var connectStrForMigration = strings.Replace(ndb.cfg.ConnStr, dbName, "", 1)
	migrationDb, err := sql.Open("mysql", connectStrForMigration+"&multiStatements=true")
	if err != nil {
		panic(err)
	}
	defer migrationDb.Close()
	_, err = migrationDb.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + " DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;")
	if err != nil {
		panic(err)
	}
	_, err = migrationDb.Exec("USE " + dbName)
	if err != nil {
		panic(err)
	}

	// 如果没有则创建数据库
	dbmigrate.Migrate(migrationDb) // 指定数据库版本
}

func GetDbNameFromConnectStr(connectStr string) (dbName string) {
	begin := strings.Index(connectStr, "/")
	if begin == -1 {
		return ""
	}
	end := strings.Index(connectStr, "?")
	if end == -1 {
		return ""
	}
	dbName = connectStr[begin+1 : end]
	return dbName
}
