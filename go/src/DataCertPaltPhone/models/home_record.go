package models

import "DataCertPaltPhone/db_mysql"

type HomeRecord struct {
	Id int
	UserId int
	FileName string
	FileSize int
	FileCert string
	FileTitle string
	CertTime int
}

func (u HomeRecord)SavaRecord()  {
	db_mysql.Db.Exec("insert into user_id, ")
}