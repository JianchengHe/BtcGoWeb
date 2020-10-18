package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

/*
对一个字符串进行md5计算
 */
func Md5HashString(data string)  string{
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	bytes:= md5Hash.Sum(nil)
	return hex.EncodeToString(bytes)
}
//读取io流中的数据，并对数据进行hash计算,返回sha256，hash值
func Sha256Reader(reader io.Reader) (string,error) {
	sha256Hash := sha256.New()
	readerBytes,err :=ioutil.ReadAll(reader)
	if err !=nil {
		return "",err
	}
	sha256Hash.Write(readerBytes)
	hashBytes := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}

func Md5HashReader(reader io.Reader)(string,error)  {
	Md5Hash := md5.New()
	readerBytes,err := ioutil.ReadAll(reader)
	if err != nil {
		return "",nil
	}
	Md5Hash.Write(readerBytes)
	hashBytes := Md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}