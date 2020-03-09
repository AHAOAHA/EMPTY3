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
}

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "login",
	})
}

func LoginPostHandler(c *gin.Context) {
	c.Request.ParseForm()
	// session := sessions.Default(c)
	// session.Set("ahaoo", "test_val")
	// session.Save()
	// get type
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
		log.Error("login type err")
		return
	}
	if err := loginer.IsValid(); err != nil {
		log.Error(err)
	}
	for k, v := range c.Request.PostForm {
		log.Info(k, ":", v)
	}
	if err := loginer.Login(c.Request.PostForm.Get("username"), c.Request.PostForm.Get("password")); err != nil {
		log.Warn("login error")
		c.Redirect(http.StatusMovedPermanently, "/debug_loginerr")
		return
	}
	loginer.RedirectIndex(c)
}
