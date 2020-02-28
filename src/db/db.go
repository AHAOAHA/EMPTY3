/******************************************************************
 * Copyright(C) 2016-2020. All right reserved.
 * 
 * Filename: db.go
 * Author: ahaoozhang
 * Date: 2020-02-01 15:26:26 (Saturday)
 * Describe: 
 ******************************************************************/
package db

import (
	"database/sql"
	"log"
    _ "github.com/go-sql-driver/mysql"
)

type MyDB struct {
	Db*			 	 sql.DB
	IP				 string
	Port			 uint32
	UserName		 string
	PassWord		 string
	DatabaseName	 string
	// TODO cache 
}

func (f *MyDB) OpenDB() bool {
	if len(f.IP) == 0 || f.Port == 0 || len(f.UserName) == 0 || len(f.DatabaseName) == 0 || len(f.PassWord) == 0 {
		log.Print("Prase fail")
		return false
	}
	Odb, err := sql.Open("mysql", f.UserName + ":" + f.PassWord + "@tcp(" + f.IP + ")/" + f.DatabaseName + "?charset=utf8")
	if err != nil {
		return false
	}
	f.Db = Odb
	return true
}