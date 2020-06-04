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
	"GradeManager/src/common"
	"GradeManager/src/config"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type DB interface {
	IsValid() bool
	OpenDB() error
	Query(rowStruct interface{}, format string, args ...interface{}) error
	GetDBDrive() *sql.DB
}

type MyDB struct {
	db           *sql.DB
	IP           string
	Port         uint32
	UserName     string
	PassWord     string
	DatabaseName string
}

var DataBase *MyDB

func init() {
	go func() { // check db status runtine
		for true {
			if DataBase != nil {
				err := DataBase.db.Ping()
				if err != nil {
					log.Warnf("Database Ping Err")
				} else {
					time.Sleep(time.Second * 3)
					//log.Infof("Database Normal")
					continue
				}
			}
			if config.Config.GradeManagerDB.IsValid() {
				// init sql
				DataBase = new(MyDB)
				DataBase.IP = config.Config.GradeManagerDB.Host
				DataBase.PassWord = config.Config.GradeManagerDB.Password
				DataBase.DatabaseName = config.Config.GradeManagerDB.DataBaseName
				DataBase.Port = config.Config.GradeManagerDB.Port
				DataBase.UserName = config.Config.GradeManagerDB.User
				err := DataBase.OpenDB()
				if err != nil {
					log.Fatal("Databases Init err")
				}
				log.Info("Database init success")
			}
		}
	}()

}

func (f *MyDB) IsValid() bool {
	if len(f.IP) == 0 || f.Port == 0 || len(f.UserName) == 0 || len(f.DatabaseName) == 0 || len(f.PassWord) == 0 {
		return false
	}
	return true
}

func (f *MyDB) OpenDB() error {
	Odb, err := sql.Open("mysql", f.UserName+":"+f.PassWord+"@tcp("+f.IP+")/"+f.DatabaseName+"?charset=utf8")
	if err != nil {
		return err
	}
	f.db = Odb
	return nil
}

// sql Query format
func (f *MyDB) Query(format string, args ...interface{}) ([]map[string]interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			msgSt := struct {
				t   time.Time
				sql string
			}{
				time.Now(),
				format,
			}
			msgJs, _ := json.Marshal(msgSt)
			common.SendTextToWechat("数据库查询引发panic", string(msgJs))
			log.Error("DB Query panic, recovered!")
		}
	}()

	if f == nil {
		return nil, errors.New("Sql MyDB point is nil")
	}
	err := f.db.Ping()
	if err != nil {
		return nil, err
	}
	// 查询语句传入不存在的字段名时间，会引发panic
	log.Infof(format, args...)
	rows, err := f.db.Query(format, args...)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index, _ := range cache {
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{}
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		list = append(list, item)
	}
	return list, nil
}

func (db *MyDB) GetDBDrive() *sql.DB {
	return db.db
}

// use golang format
func (db *MyDB) Queryf(format string, args ...interface{}) ([]map[string]interface{}, error) {
	sql := fmt.Sprintf(format, args...)
	return db.Query(sql)
}

func (db *MyDB) Execf(format string, args ...interface{}) error {
	sql := fmt.Sprintf(format, args...)
	defer func() {
		if err := recover(); err != nil {
			msgSt := struct {
				t   time.Time
				sql string
			}{
				time.Now(),
				sql,
			}
			msgJs, _ := json.Marshal(msgSt)
			common.SendTextToWechat("数据库执行引发panic", string(msgJs))
			log.Error("db Exec Panic, recovered!")
		}
	}()

	log.Info(sql)
	_, err := db.db.Exec(sql)
	if err != nil {
		log.Error(err)
	}
	return err
}
