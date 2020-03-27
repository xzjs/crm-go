package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	ID     int64  `orm:"pk;auto;column(id)"`
	Mobile string `orm:"size(11);cplumn(mobile)"`
	Type   int    `orm:"size(1);cplumn(type)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func GetUserByMobile(mobile string) (id int64, err error) {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	if _, id, err := o.ReadOrCreate(&user, "Mobile"); err == nil {
		return id, nil
	}
	return 0, err

}

func GetUserById(id int64) (user User, err error) {
	o := orm.NewOrm()
	user = User{ID: id}
	err = o.Read(&user)
	return user, err
}
