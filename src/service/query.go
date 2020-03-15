/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: query.go
 * Author: ahaoozhang
 * Date: 2020-03-15 19:27:57 (Sunday)
 * Describe:
 ******************************************************************/
package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryOrder struct {
	UserType     string
	QueryMothod  string
	QueryContent string
}

func QueryGetHandler(c *gin.Context) {
	// TODO: 鉴权
	var query_command QueryOrder
	query_command.QueryMothod = c.Query("query_mothod")
	query_command.UserType = c.Query("usertype")
	query_command.QueryContent = c.Query("query_content")
	query_command.RenderHTML(c)
}

func (qc *QueryOrder) RenderHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "query_result.html", gin.H{
		"testList": nil,
	})
}
