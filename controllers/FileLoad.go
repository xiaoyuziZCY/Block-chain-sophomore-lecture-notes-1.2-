package controllers

import (
	"authentication/models"
	"authentication/util"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"time"
)

type LoadController struct {
	beego.Controller
}
func (u *LoadController) Get(){
	phone := u.GetString("phone")
	u.Data["Phone"] = phone
	u.TplName = "home.html"
}
func (u *LoadController) Post() {
	filetitle :=u.Ctx.Request.PostFormValue("up_title")
	phone := u.Ctx.Request.PostFormValue("phone")
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
	fmt.Println("解析到上传文件，文件名是："+head.Filename)
	LoadDir :="./static/img/" + head.Filename
	savefile,err:=os.OpenFile(LoadDir,os.O_RDWR|os.O_CREATE,777)

	writer:=bufio.NewWriter(savefile)
	fmt.Printf("%t\n",LoadDir)
	One_files,err :=io.Copy(writer,file)
	fmt.Println(One_files)
	if err !=nil {
		u.Ctx.WriteString("对不起,保存文件失败")
		return
	}
	defer savefile.Close()

	hashFile,err:=os.Open(LoadDir)
	defer hashFile.Close()
	hash,err:=util.MD5HashReader(hashFile)
	if err !=nil {
		fmt.Println(err.Error())
		return
	}
	record := models.UploadRecord{}
	record.FileName=head.Filename
	record.FileSize=head.Size
	record.FileTitle=filetitle
	record.CertTime=time.Now().Unix()
	record.FileCert=hash
	record.Phone=phone
	_, err = record.SaveRecord()
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉，数据认证错误, 请重试!")
		return
	}
	records, err := models.QueryRecordByPhone(phone)
	if err != nil {
		u.Ctx.WriteString("抱歉，获取认证数据失败, 请重试!")
		return
	}
fmt.Println(record)
u.Data["Records"]=records
u.Data["Phone"]=phone
u.TplName="cert_detail.html"
}
