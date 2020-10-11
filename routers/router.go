package routers

import (
	"authentication/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
}
