package util

import (
	//"authentication/blockchain"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

//字符串哈希
func MD5HashSting(data string)string{
	md5Hash :=md5.New()
	md5Hash.Write([]byte(data))
	passwordBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(passwordBytes)
}
//
func MD5HashReader(reader io.Reader) (string,error) {
	bytes,err:=ioutil.ReadAll(reader)
	if err !=nil {
		fmt.Println(err.Error())
		return "",err
	}
md5Hash:=md5.New()
md5Hash.Write(bytes)
hashBytes:=md5Hash.Sum(nil)
return hex.EncodeToString(hashBytes),nil
}
//
func SHA256HashBlock(block []byte)([]byte){
sha256hash:=sha256.New()
sha256hash.Write([]byte(""))
return sha256hash.Sum(nil)
}
