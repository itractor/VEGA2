package main

import (
	"fmt"

	"github.com/astaxie/beego"
)

func init() {
	fmt.Println("Im initialized")
	beego.Router("/", &MainController{})
	beego.Router("/test", &TestController{})
}
