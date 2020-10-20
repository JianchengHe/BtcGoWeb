package utils

import (
	"bytes"
	"encoding/binary"
)

/*
将一个int64转化为[]byte类型
 */
func Int64ToByte(num int64) ([]byte,error) {
	//beffer  缓冲区。。增益
	buff := new(bytes.Buffer)
	//buff.Write()//通过一系列write方法向缓冲区写入
	//buff.Bytes()//通过bytes方法从缓冲区中获取数据
	/*
	大端位序排列 binary.BigEndian
	小端位序排列binary.LittleEndian
	 */
	err := binary.Write(buff,binary.BigEndian,num)
	if err != nil {
		return nil,err
	}
	//从缓冲区中读取数据
	return buff.Bytes(),nil
}
/*
将字符串转换为[]byte类型
 */
func StringToByte(data string) []byte {
	return []byte(data)
}