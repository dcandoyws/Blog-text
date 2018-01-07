package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	// 设置数据库路径
	_DB_NAME = "data/blogtext.db"
	// 设置数据库名称
	_SQLITE3_DRIVER = "sqlite3"
)

type Text struct {
	Id int64

	Title string

	Content string `orm:"size(5000)"`

	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views   int64     `orm:"index"`
}

func InitDb() {
	// 检查数据库文件
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	// 注册默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
	// 注册模型
	orm.RegisterModel(new(Text))
	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)

}

func AddText(title, content string) error {

	o := orm.NewOrm()

	text := &Text{
		Title: title,

		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}
	_, err := o.Insert(text)
	if err != nil {
		return err
	}
	return err
}

func GetAllText() ([]*Text, error) {
	o := orm.NewOrm()

	texts := make([]*Text, 0)
	qs := o.QueryTable("text")
	_, err := qs.All(&texts)
	return texts, err
}

func GetText(id string) (*Text, error) {
	tid, _ := strconv.ParseInt(id, 10, 64)

	o := orm.NewOrm()

	text := &Text{Id: tid}
	if o.Read(text) == nil {
		text.Updated = time.Now()
		text.Views++
		o.Update(text)
	}
	o.Read(text)

	return text, nil
}

func ModifyText(tid, title, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	text := &Text{Id: tidNum}
	if o.Read(text) == nil {
		text.Title = title

		text.Content = content
		text.Updated = time.Now()
		o.Update(text)
	}

	return nil
}

func DeleteText(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	text := &Text{Id: tidNum}

	_, err = o.Delete(text)
	return err
}
