/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: admin.go
 * Author: ahaoozhang
 * Date: 2020-03-09 22:34:51 (Monday)
 * Describe:
 ******************************************************************/
package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminIndexHandler(c *gin.Context) {
	//var a context.AdminContext
	// cookies, err := c.Request.Cookie("userinfo")
	// if err != nil {
	// 	log.Error(err)
	// 	c.Redirect(http.StatusMovedPermanently, "/login")
	// 	return
	// }
	// log.Info(cookies)
	c.HTML(http.StatusOK, "admin_index.html", gin.H{
		"title": "login",
	})
}
