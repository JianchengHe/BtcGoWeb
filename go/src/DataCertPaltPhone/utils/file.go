package utils

import (
	"fmt"
	"io"
	"os"
)

/*
保存一个文件
 */
func SavaFile(fileName string,file io.Reader)(int64,error)  {
	savaFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 777)
	if err != nil {
		return -1,err
	}

	length, err := io.Copy(savaFile, file)
	if err != nil {
		fmt.Println("保存文件：",err.Error())
		return -1,err
	}
	return length,nil
	
}
