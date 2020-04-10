package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     int64
	Mobile string `orm:"size(11)"`
	Type   int    `orm:"size(1)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// 根据手机号获取或注册用户
func GetUserByMobile(mobile string) (id int64, err error) {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	_, id, err = o.ReadOrCreate(&user, "Mobile")
	return id, err
}

func GetUserById(id int64) (user User, err error) {
	o := orm.NewOrm()
	user = User{Id: id}
	err = o.Read(&user)
	return user, err
}

// 级联获取所有文件
func (u *User) GetFiles() (files []*File, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("file").Filter("User", u.Id).RelatedSel().All(&files)
	return files, err
}

// 获取用户所有任务
func (u *User) GetTasks() (tasks []*Task, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("task").Filter("User", u.Id).RelatedSel().OrderBy("-start_time").All(&tasks)
	return tasks, err
}
