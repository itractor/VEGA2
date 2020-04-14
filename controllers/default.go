package controllers

import (
	"github.com/astaxie/beego"
)

type ReturnJsonTest struct {
	S string `json:"s"`
	D string `json:"d"`
}

type MainController struct {
	beego.Controller
}

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	defer c.ServeJSON()
	var responseJson ReturnJsonTest
	responseJson = ReturnJsonTest{
		S: "asdf",
		D: "qwer",
	}

	c.Data["json"] = responseJson
}
