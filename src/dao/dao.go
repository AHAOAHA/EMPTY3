/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: db.go
 * Author: ahaoozhang
 * Date: 2020-02-01 15:26:26 (Saturday)
 * Describe:
 ******************************************************************/
package dao

import (
	"database/sql"
	"errors"
	"reflect"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type DB interface {
	IsValid() bool
	OpenDB() bool
	Query(rowStruct interface{}, format string, args ...interface{}) error
}

type MyDB struct {
	DB           *sql.DB
	IP           string
	Port         uint32
	UserName     string
	PassWord     string
	DatabaseName string
}

func (f *MyDB) IsValid() bool {
	if len(f.IP) == 0 || f.Port == 0 || len(f.UserName) == 0 || len(f.DatabaseName) == 0 || len(f.PassWord) == 0 {
		log.Print("Prase fail")
		return false
	}
	return true
}

func (f *MyDB) OpenDB() bool {
	Odb, err := sql.Open("mysql", f.UserName+":"+f.PassWord+"@tcp("+f.IP+")/"+f.DatabaseName+"?charset=utf8")
	if err != nil {
		return false
	}
	f.DB = Odb
	return true
}

func (f *MyDB) Query(rowStruct interface{}, format string, args ...interface{}) error {
	if f == nil {
		return errors.New("Sql MyDB point is nil")
	}
	rows, err := f.DB.Query(format, args)
	defer rows.Close()
	if err != nil {
		return errors.New("DB Query err")
	}
	// 确定Scan函数输入的类型
	s := reflect.ValueOf(rowStruct).Elem()
	onerow := make([]interface{}, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		onerow[i] = s.Field(i).Addr().Interface()
	}
	for rows.Next() {
		if err := rows.Scan(onerow...); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
