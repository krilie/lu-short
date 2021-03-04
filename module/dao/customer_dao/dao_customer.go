package customer_dao

import (
	"lu-short/component/ndb"
	"lu-short/component/nlog"
)

type CustomerDao struct {
	dao *ndb.NDb
	log *nlog.NLog
}

func NewCustomerDao(dao *ndb.NDb, log *nlog.NLog) *CustomerDao {
	return &CustomerDao{
		dao: dao,
		log: log,
	}
}
