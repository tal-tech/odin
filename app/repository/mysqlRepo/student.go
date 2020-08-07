package mysqlRepo

import (
	"time"

	dbdao "github.com/tal-tech/torm"
)

type Student struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Sex        int       `xorm:"not null default 0 TINYINT(4)"`
	Name       string    `xorm:"not null default '' VARCHAR(256)"`
	City       string    `xorm:"not null default '' CHAR(64)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('插入时间') TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('更新时间') TIMESTAMP"`
}

type StudentDao struct {
	dbdao.DbBaseDao
}

func NewStudentDao(v ...interface{}) *StudentDao {
	this := new(StudentDao)
	if ins := dbdao.GetDbInstance("godemo", "writer"); ins != nil {
		this.UpdateEngine(ins.Engine)
	} else {
		return nil
	}
	if len(v) != 0 {
		this.UpdateEngine(v...)
	}
	return this
}

func (this *StudentDao) Get(mId dbdao.Param) (ret []Student, err error) {
	ret = make([]Student, 0)
	this.InitSession()

	this.BuildQuery(mId, "id")

	err = this.Session.Find(&ret)
	return
}
func (this *StudentDao) GetLimit(mId dbdao.Param, pn, rn int) (ret []Student, err error) {
	ret = make([]Student, 0)
	this.InitSession()

	this.BuildQuery(mId, "id")

	err = this.Session.Limit(rn, pn).Find(&ret)
	return
}
func (this *StudentDao) GetCount(mId dbdao.Param) (ret int64, err error) {
	this.InitSession()

	this.BuildQuery(mId, "id")

	ret, err = this.Session.Count(new(Student))
	return
}
