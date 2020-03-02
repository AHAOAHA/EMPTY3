/******************************************************************
 * Copyright(C) 2016-2020. All right reserved.
 *
 * Filename: config.go
 * Author: ahaoozhang
 * Date: 2020-01-27 12:48:00 (Monday)
 * Describe:
 ******************************************************************/
package config

import (
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

// 配置文件必须实现的接口
type ConfigInfo interface {
	IsValid() bool
}

type GradeManagerDBInfo struct {
	Host         string
	Port         uint32
	User         string
	Password     string
	DataBaseName string
}

type MyConfig struct {
	GradeManagerDB GradeManagerDBInfo
}

var Config MyConfig

func init() {
	// prase config
	if _, err := toml.DecodeFile("/home/ahaoozhang/dev_code/GradeManager/config/config.toml", &Config); err != nil {
		panic(err)
	}
	log.Infof("MyConfig: %v", Config)
}

func (db *GradeManagerDBInfo) IsValid() bool {
	if db.Host != "" && db.DataBaseName != "" && db.Password != "" && db.Port != 0 && db.User != "" {
		return false
	}
	return true
}
