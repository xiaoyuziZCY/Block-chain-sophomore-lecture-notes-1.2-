package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type CheckController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "register.html"
}
func (c *CheckController) Get() {
	c.TplName = "list_record.html"
}
