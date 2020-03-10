/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: 404.go
 * Author: ahaoozhang
 * Date: 2020-03-10 00:55:16 (Tuesday)
 * Describe:
 ******************************************************************/
package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"title": "not found!",
	})
}
