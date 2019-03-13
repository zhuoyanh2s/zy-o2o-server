package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"zy-o2o-server/utils"
)

type User struct {
	Id       int64
	Username string
	Password string
	State    int8
	IsSuper  bool
	LastTime time.Time
	Created  time.Time
	Updated  time.Time
	Profile  *Profile `orm:"rel(fk)"`
}

type Profile struct {
	Id       int
	RealName string
	Age      int
	Email    string
	Mobile   string
	Address  string
	Gender   string
}

func AddUser(u User) (bool, error) {
	o := orm.NewOrm()
	pw := utils.Md5(u.Password)
	u.Password = pw

	_, err := o.Insert(u)
	if err != nil {
		beego.Error("保存用户数据到数据库时失败 =>", err)
		return false, errors.New("保存用户失败")
	}
	return true, nil
}

func DeleteUser(id int64) (bool, error) {
	o := orm.NewOrm()
	user := User{Id: id}
	_, err := o.Delete(&user)
	if err != nil {
		beego.Error("删除用户数据保存到数据库时失败 =>", err)
		return false, errors.New("删除用户失败")
	}
	return true, nil

}

func BatchDeleteUser(ids []int) (bool, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	_, err := qs.Filter("id__in", ids).Delete()

	if err != nil {
		beego.Error("批量删除用户数据保存到数据库时失败 =>", err)
		return false, errors.New("批量删除用户失败")
	}
	return true, nil

}

func UpdateUser(u *User) (user *User, err error) {
	o := orm.NewOrm()
	_, err = o.Update(u)
	if err != nil {
		beego.Error("更新用户数据保存到数据库时失败 =>", err)
		return nil, errors.New("更新用户失败")
	}

	return user, nil
}

func GetUser(id int64) (user *User, err error) {
	o := orm.NewOrm()
	u := User{Id: id}
	err = o.Read(u)
	if err == nil {
		return user, nil
	}
	return nil, errors.New("User not exists")
}

func GetUserByName(username string) (user *User, err error) {
	o := orm.NewOrm()
	u := User{Username: username}
	err = o.Read(u)
	if err == nil {
		return user, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() (users []*User, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(User)).All(users)

	if err != nil {
		beego.Error("批量删除用户数据保存到数据库时失败 =>", err)
		return nil, errors.New("批量删除用户失败")
	}
	return users, nil
}

func Login(username, password string) bool {
	if u, err := GetUserByName(username); err == nil {
		pw := utils.Md5(password)
		println(pw)
		if u.Password == password {
			return true
		}
	}
	return false
}
