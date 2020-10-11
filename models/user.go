package models

import (
	"authentication/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	password string `form:"password"`
}

func (u User) SaveUser()(int64,error) {
	md5Hash :=md5.New()
	md5Hash.Write([]byte(u.password))
	passwordBytes := md5Hash.Sum(nil)
	u.password = hex.EncodeToString(passwordBytes)
	//--------------
	row,err:=db_mysql.Db.Exec("insert into users(phone,password) value (?,?)",u.Phone,u.password)
	if err != nil {
		return -1,err
	}
	id,err :=row.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil
}
