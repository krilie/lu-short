package ndb

import (
	"context"
)

func Like(val string) string {
	return "%" + val + "%"
}

// 执行一个count查询 一个date查询
func GetPageDataFormSql(ctx context.Context, db IDb, countSql, dataSql string, countArgs, dataArgs []interface{}, data interface{}) (totalCount int64, err error) {
	totalCount, err = Count(ctx, db, countSql, countArgs...)
	if err != nil {
		return 0, err
	}
	if totalCount <= 0 {
		return 0, nil
	}
	_, err = db.WithContext(ctx).Select(data, dataSql, dataArgs...)
	if err != nil {
		return 0, err
	}
	return totalCount, nil
}

// 返回结果的执行 只返回一个数值的 如 select count(1) ...
func Count(ctx context.Context, db IDb, sql string, values ...interface{}) (count int64, err error) {
	return db.WithContext(ctx).SelectInt(sql, values...)
}
