package controllers

import (
	"authentication/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller

}
func (r *RegisterController) Get() {
	var user models.User
	err :=r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("对不起，解析数据错误")
		return
	}
	_,err =user.SaveUser()
	if err !=nil  {
		r.Ctx.WriteString("对不起，用户注册失败")
		return
	}
	r.TplName ="register.html"
}
func (r *RegisterController) Post() {
	var user models.User
	err :=r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("对不起，解析数据错误")
		return
	}
	_,err =user.SaveUser()
	if err !=nil  {
		r.Ctx.WriteString("对不起，用户注册失败")
		return
	}
	r.TplName ="register.html"
}
