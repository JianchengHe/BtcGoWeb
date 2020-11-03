package models

import (
	"DataCertPaltPhone/db_mysql"
	"DataCertPaltPhone/utils"
	//"time"
)

type HomeRecord struct {
	Id        int
	UserId    int
	FileName  string
	FileSize  int64
	FileCert  string //认证号：md5值
	FileTitle string
	CertTime  int64
	CertTimeFormat string
}

//把上传的文件作为记录保存到数据库当中
func (u HomeRecord) SavaRecord()(int64 ,error) {
	result,err := db_mysql.Db.Exec("insert into home_record(user_id,file_name,file_size,file_cert,file_title,cert_time)values(?,?,?,?,?,?)",
		u.UserId, u.FileName, u.FileSize, u.FileCert, u.FileTitle, u.CertTime)
	if err != nil{//保存数据遇到错误
		return -1,err
	}
	id,err := result.RowsAffected()
	if err != nil {//保存数据遇到错误
		return -1,err
	}
	return id,nil
}


/*
根据用户id查询符合条件的认证数据记录
 */
func QueryRecordsByUserId(userId int) ([]HomeRecord,error) {
	rows, err := db_mysql.Db.Query("select id,user_id,file_name,file_size,file_cert,file_title,cert_time from home_record where user_id = ?",userId)
	if err != nil {
		return nil,err
	}
	records := make([]HomeRecord, 0)
	for rows.Next() {
		var record HomeRecord
		err = rows.Scan(&record.Id, &record.UserId, &record.FileName, &record.FileSize, &record.FileCert, &record.FileTitle, &record.CertTime)
		if err != nil {
			return nil,err
		}
		//整形-->字符串
		tStr := utils.TimeFormat(record.CertTime,utils.TIME_FORMAT_THREE)
		record.CertTimeFormat = tStr
		records = append(records,record)
	}
	return records,nil
}
