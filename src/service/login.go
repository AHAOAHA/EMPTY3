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
	"GradeManager/src/context"
	_ "GradeManager/src/dao"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Loginer interface {
	IsValid() error
	Login(string, string) error
	RedirectIndex(*gin.Context) error
	SetCookies(*gin.Context)

	// protobuf marshal, then base64 encode.
	Entcry() string
	Detcry(cookie string) error

	// check cookie format and update.
	CheckCookies(c *gin.Context, key string) error

	GetPassword() string
	UpdatePassword(new_password string) error
}

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "login",
	})
}

func SignUpHandler(c *gin.Context) {
	c.Request.ParseForm()
	login_type := c.Request.PostForm.Get("type")
	var loginer Loginer
	switch login_type {
	case "admin":
		loginer = new(context.AdminContext)
		break
	case "student":
		loginer = new(context.StudentContext)
		break
	case "teacher":
		loginer = new(context.TeacherContext)
		break
	default:
		log.Errorf("login type err, type: %s\n", login_type)
		return
	}
	for k, v := range c.Request.PostForm {
		log.Info(k, ":", v)
	}

	if err := loginer.Login(c.Request.PostForm.Get("username"), c.Request.PostForm.Get("password")); err != nil {
		log.Warn(err)
		c.HTML(http.StatusOK, "login.html", gin.H{
			"err_code": "false",
		})
		return
	}
	loginer.SetCookies(c)
	loginer.RedirectIndex(c)
}

func SignOutHandler(c *gin.Context) {
	c.SetCookie("user_cookie", "out", 10, "/", "", false, true)
	c.Redirect(http.StatusMovedPermanently, "/login")
}
