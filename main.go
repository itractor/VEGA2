package main

import (
	"github.com/astaxie/beego"

	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	println("Initlazid")

	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "ba:ba@/orm_test?charset=utf8", maxIdle, maxConn)
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // Using default, you can use other database

	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"
	fmt.Println(profile)

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))

	beego.Run()
}
