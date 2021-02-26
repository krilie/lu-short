package ndb

import (
	"context"
	"lu-short/common/errs"
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
	rows, err := db.Queryx(dataSql, dataArgs...)
	if err != nil {
		return 0, errs.NewInternal().WithError(err)
	}
	defer rows.Close()
	err = rows.Scan(data)
	if err != nil {
		return 0, errs.NewInternal().WithError(err)
	}
	return totalCount, nil
}

// 返回结果的执行 只返回一个数值的 如 select count(1) ...
func Count(ctx context.Context, db IDb, sql string, values ...interface{}) (count int64, err error) {
	raw, err := db.Queryx(sql, values...)
	if err != nil {
		return 0, err
	}
	defer raw.Close()
	err = raw.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
