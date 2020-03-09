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
	"net/http"

	"github.com/gin-gonic/gin"
)

func StudentIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "student_index.html", gin.H{
		"title": "login",
	})
}
