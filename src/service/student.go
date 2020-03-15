/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: student.go
 * Author: ahaoozhang
 * Date: 2020-03-09 22:35:05 (Monday)
 * Describe:
 ******************************************************************/
package service

import (
	"GradeManager/src/context"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func StudentIndexHandler(c *gin.Context) {
	var s context.StudentContext
	// check cookie
	cookie, err := c.Request.Cookie("user_cookie")
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}
	err = s.Detcry(cookie.Value)
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	// login status ok
	c.HTML(http.StatusOK, "student_index.html", gin.H{
		"title": "login",
		// TODO:
	})
}
