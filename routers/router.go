package routers

import (
	"authentication/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/zcy", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
}
