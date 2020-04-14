package routers

import (
	"VEGA2/controllers"
	"fmt"

	"github.com/astaxie/beego"
)

func init() {
	fmt.Println("Im initialized")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})
}
