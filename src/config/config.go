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

type AlarmInfo struct {
	Url string
}

type LocationToAddrInfo struct {
	Key  string
	SK   string
	Url  string
	Path string
}

type IPToAddrInfo struct {
	Key  string
	SK   string
	Url  string
	Path string
}

type MyConfig struct {
	Alarm          AlarmInfo
	GradeManagerDB GradeManagerDBInfo
	LocationToAddr LocationToAddrInfo
	IPToAddr       IPToAddrInfo
}

var Config MyConfig

func init() {
	// prase config
	if _, err := toml.DecodeFile("/home/ahaoozhang/dev_code/GradeManager/config/config.toml", &Config); err != nil {
		panic(err)
	}
	log.Infof("MyConfig: %+v", Config)
}

func (db *GradeManagerDBInfo) IsValid() bool {
	if db.Host != "" && db.DataBaseName != "" && db.Password != "" && db.Port != 0 && db.User != "" {
		return true
	}
	return false
}

func (alarm *AlarmInfo) IsValid() bool {
	if alarm.Url == "" {
		return false
	}

	return true
}

func (ita *LocationToAddrInfo) IsValid() bool {
	// TODO
	return true
}

func (ita *IPToAddrInfo) IsValid() bool {
	// TODO
	return true
}
