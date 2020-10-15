package controllers

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
)

type LoadController struct {
	beego.Controller
}

func (u *LoadController) Post() {
	filetitle :=u.Ctx.Request.PostFormValue("up_title")
	file,head,err :=u.GetFile("up_file")
	if err != nil {
		u.Ctx.WriteString("对不起，文件解析错误")
		return
	}
	fmt.Println(filetitle)
	fmt.Println(head.Header)
	fmt.Println(head.Size)
	fmt.Println(file)
	fmt.Println(head.Filename)
	u.Ctx.WriteString("解析到上传文件，文件名是："+head.Filename)
	LoadDir :="./static/img/" + head.Filename
	savefile,err:=os.OpenFile(LoadDir,os.O_RDWR|os.O_CREATE,777)

	writer:=bufio.NewWriter(savefile)
	fmt.Printf("%t\n",LoadDir)
	One_files,err :=io.Copy(writer,file)
	if err !=nil {
		u.Ctx.WriteString("对不起,保存文件失败")
		return
	}
	fmt.Println("拷贝到的文件为：",One_files)
}
