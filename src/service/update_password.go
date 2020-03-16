/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: update_password.go
 * Author: ahaoozhang
 * Date: 2020-03-16 18:35:21 (Monday)
 * Describe:
 ******************************************************************/
package service

import (
	"GradeManager/src/context"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UpdatePasswordGetHandler(c *gin.Context) {
	var a context.AdminContext
	var s context.StudentContext
	var t context.TeacherContext
	err_a := a.CheckCookies(c, "user_cookie")
	err_s := s.CheckCookies(c, "user_cookie")
	err_t := t.CheckCookies(c, "user_cookie")
	if err_a != nil && err_s != nil && err_t != nil {
		// 鉴权失败
		c.HTML(http.StatusBadRequest, "401.html", "nil")
		return
	}
	var login_name string
	login_type := c.Query("type")
	switch login_type {
	case "admin":
		login_name = a.Info.GetUser()
		break
	case "student":
		login_name = s.Info.GetName()
		break
	case "teacher":
		login_name = t.Info.GetName()
		break
	default:
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}
	c.HTML(http.StatusOK, "update_password.html", gin.H{
		"type":       login_type,
		"login_name": login_name,
	})
}

func UpdatePasswordPostHandler(c *gin.Context) {
	c.Request.ParseForm()
	var loginer Loginer
	login_type := c.Request.PostForm.Get("type")
	for k, v := range c.Request.PostForm {
		log.Info(k, ":", v)
	}
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
		log.Error("update password login type err")
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	if err := loginer.CheckCookies(c, "user_cookie"); err != nil {
		log.Error(err)
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	old_password := c.Request.PostForm.Get("old_password")
	new_password := c.Request.PostForm.Get("new_password")

	if old_password != loginer.GetPassword() {
		c.HTML(http.StatusOK, "update_password_fail.html", gin.H{
			"second": "3",
			"url":    "/update_password?type=" + login_type,
		})
		return
	}

	err := loginer.UpdatePassword(new_password)
	if err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.SetCookie("user_cookie", "out", 10, "/", "", false, true)
	c.Redirect(http.StatusMovedPermanently, "/login")
}
