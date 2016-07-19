package models

import (
	//"github.com/Unknow/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"os"
	//"path"
	"time"
)

const (
	_DB_NAME = "test"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attactment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func init() {
	//orm.RegisterDriver("mysql", orm.DR_MySQL)
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Category), new(Topic))
	// create table
	orm.RunSyncdb("default", false, true)
}
func AddCategory(name string) error {
	o := orm.NewOrm()
	c := &Category{Title: name}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(c)
	if err != nil {
		return err
	}
	_, err = o.Insert(c)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	cs := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cs)
	return cs, err
}

/**
func RegisterDB(){
	if !com.IsExist(_DB_NAME){
		os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category))
    orm.RegisterModel(new(Topic))
    //orm.RegisterDriver("mysql", orm.DR_MySQL)
    orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)
}*/
