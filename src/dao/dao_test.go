/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: dao_test.go
 * Author: ahaoozhang
 * Date: 2020-03-04 19:32:31 (Wednesday)
 * Describe:
 ******************************************************************/
package dao

import (
	"fmt"
	"testing"
)

var db MyDB

func TestQuery(t *testing.T) {
	db.IP = "101.37.175.110"
	db.Port = 3306
	db.UserName = "GradeManager"
	db.PassWord = "19980721"
	db.DatabaseName = "GradeManager"
	fmt.Println("sdfdsdfds")
	err := db.OpenDB()
	if err != nil {
		t.Log(err)
	}
	// TODO: sql Query format
	m, err := db.Query("select password from admin where user='?'", "admin")
	if err != nil {
		t.Log(err)
	}
	for k, v := range m {
		t.Log(k, ":", v)
	}
}
