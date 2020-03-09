/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: teacher.go
 * Author: ahaoozhang
 * Date: 2020-03-09 22:35:18 (Monday)
 * Describe:
 ******************************************************************/
package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TeacherIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "teacher_index.html", gin.H{
		"title": "login",
	})
}
