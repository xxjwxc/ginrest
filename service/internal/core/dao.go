package core

import (
	"runtime"

	"github.com/xxjwxc/public/errors"

	"github.com/xxjwxc/public/mylog"

	"github.com/xxjwxc/gormt/data/config"
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
	runtime.SetFinalizer(Dao, Dao.Destory) //析构时触发
	Dao.InitDao()
}

func GetDBr() *mysqldb.MySqlDB {
	if Dao.dbr == nil {
		mylog.Error(errors.New("dbr is nil."))
	}
	return Dao.dbr
}

func GetDBw() *mysqldb.MySqlDB {
	if Dao.dbw == nil {
		mylog.Error(errors.New("dbw is nil."))
	}
	return Dao.dbw
}

func (dao *DaoCore) InitDao() {
	dao.dbr = mysqldb.OnInitDBOrm(config.GetMysqlConStr())
	dao.dbw = mysqldb.OnInitDBOrm(config.GetMysqlConStr())
}

func (dao *DaoCore) Destory() {
	if dao.dbr != nil {
		dao.dbr.OnDestoryDB()
	}

	if dao.dbw != nil {
		dao.dbw.OnDestoryDB()
	}
}
