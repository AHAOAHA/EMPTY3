/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: login.go
 * Author: ahaoozhang
 * Date: 2020-03-02 14:56:44 (Monday)
 * Describe: service.login
 ******************************************************************/
package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Admin struct {
	User       string
	Password   string
	Mail       string
	CreateTime int32
	ExprTime   int32
	LoginIP    string
}

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "login",
	})
}
