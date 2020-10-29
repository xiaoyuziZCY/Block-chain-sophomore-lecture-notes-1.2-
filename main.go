package main

import (
	"authentication/blockchain"
	"authentication/db_mysql"
	_ "authentication/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	bc := blockchain.NewBlockChain()
	fmt.Printf("最新区块的hash值：%x\n",bc.LastHash)
	block,err :=bc.SaveData([]byte("存储数据信息"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("区块的高度:%d\n", block.Height)
	fmt.Printf("区块的PrevHash:%x\n", block.PrevHash)
	db_mysql.ConnectDB()
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

