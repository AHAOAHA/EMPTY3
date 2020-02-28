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
    _ "github.com/go-sql-driver/mysql"
    "github.com/BurntSushi/toml"
    "log"
)

type GradeManagerDBInfo struct {
    Host string
    Port uint32
    User string
    Password string
    DataBaseName string
}

type MyConfig struct {
    GradeManagerDB GradeManagerDBInfo
}

var Config MyConfig

func init() {
    // prase config
    if _, err := toml.DecodeFile("/home/ahaoozhang/dev_code/GradeManager/config/config.toml", &Config); err != nil {
        panic(err);
    }
    log.Printf("config: %+v\n", Config)
}
