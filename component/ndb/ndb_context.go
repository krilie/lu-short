package ndb

import (
	"context"
	"errors"
)

// DbValues 数据库连接上下文
type DbValues struct {
	txDb     IDb
	globalDb IDb
}

func (c *DbValues) GetTxDb() IDb {
	return c.txDb
}

func (c *DbValues) GetGlobalDb() IDb {
	return c.txDb
}

func (c *DbValues) SetTxDb(db IDb) {
	c.txDb = db
}
func (c *DbValues) SetGlobalDb(db IDb) {
	c.globalDb = db
}

// SetOrNewDbValuesCtx set db txDb to context
// if not set it will init and set
// if set before it will reset directly
func SetOrNewDbValuesCtx(ctx context.Context, db IDb, txDb IDb) context.Context {
	dbs, ok := ctx.Value("db_values_on_context").(*DbValues)
	if ok {
		dbs.SetGlobalDb(db)
		dbs.SetTxDb(txDb)
		return ctx
	} else {
		return context.WithValue(ctx, "db_values_on_context", &DbValues{
			txDb:     txDb,
			globalDb: db,
		})
	}
}

// ForceNewDbValuesCtx set db txDb to context
// if not set it will init and set
// if set before it will reset directly
func ForceNewDbValuesCtx(ctx context.Context, db IDb, txDb IDb) context.Context {
	return context.WithValue(ctx, "db_values_on_context", &DbValues{
		txDb:     txDb,
		globalDb: db,
	})
}

func GetDbValuesFromCtx(ctx context.Context) *DbValues {
	dbs, ok := ctx.Value("db_values_on_context").(*DbValues)
	if ok {
		return dbs
	} else {
		return nil
	}
}

func GetTxDbFromCtx(ctx context.Context) IDb {
	dbs := GetDbValuesFromCtx(ctx)
	if dbs != nil {
		return dbs.txDb
	}
	return nil
}
func GetGlobalDbFromCtx(ctx context.Context) IDb {
	dbs := GetDbValuesFromCtx(ctx)
	if dbs != nil {
		return dbs.globalDb
	}
	return nil
}

func MustGetTxDbFromCtx(ctx context.Context) IDb {
	db := GetTxDbFromCtx(ctx)
	if db == nil {
		panic(errors.New("no tx db found on giving context"))
	}
	return db
}

func MustGetGlobalDbFromCtx(ctx context.Context) IDb {
	db := GetGlobalDbFromCtx(ctx)
	if db == nil {
		panic(errors.New("no global db found on giving context"))
	}
	return db
}
