/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: init.go
 * Author: ahaoozhang
 * Date: 2020-03-02 17:33:06 (Monday)
 * Describe:
 ******************************************************************/
package service

import "GradeManager/src/config"

func Init() {
	// 初始化数据库
	if config.Config.GradeManagerDB.IsValid() {
		// init sql
	}
}
