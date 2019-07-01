package core

import (
	"runtime"

	"github.com/xxjwxc/ginrest/service/config"
	"github.com/xxjwxc/public/errors"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/mysqldb"
)

//
var Dao DaoCore

// Dao core dao
type DaoCore struct {
	dbr *mysqldb.MySqlDB // 读库
	dbw *mysqldb.MySqlDB // 写库
}

func init() {
	Dao.InitDao()
}

func (dao *DaoCore) GetDB() *mysqldb.MySqlDB {
	if dao.dbr == nil {
		mylog.Error(errors.New("dbr is nil."))
	}
	return dao.dbr
}

func (dao *DaoCore) GetDBr() *mysqldb.MySqlDB {
	if dao.dbr == nil {
		mylog.Error(errors.New("dbr is nil."))
	}
	return dao.dbr
}

func (dao *DaoCore) GetDBw() *mysqldb.MySqlDB {
	if dao.dbw == nil {
		mylog.Error(errors.New("dbw is nil."))
	}
	return dao.dbw
}

func (dao *DaoCore) InitDao() {
	runtime.SetFinalizer(dao, dao.Destory) //析构时触发

	dao.dbr = mysqldb.OnInitDBOrm(config.GetDbUrl())
	dao.dbw = mysqldb.OnInitDBOrm(config.GetDbUrl())
}

func (dao *DaoCore) Destory() {
	if dao.dbr != nil {
		dao.dbr.OnDestoryDB()
	}

	if dao.dbw != nil {
		dao.dbw.OnDestoryDB()
	}
}
