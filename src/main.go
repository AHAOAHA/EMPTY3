/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: main.go
 * Author: ahaoozhang
 * Date: 2020-01-16 00:13:17 (Thursday)
 * Describe:
 ******************************************************************/
package main

import (
	_ "GradeManager/src/config"
	"GradeManager/src/service"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	// 设置输出日志格式
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Static("/static", "./views/static")
	r.LoadHTMLGlob("/home/ahaoozhang/dev_code/GradeManager/views/templates/*")
	r.GET("/login", service.LoginGetHandler)
	r.POST("/login", service.LoginPostHandler)
	r.GET("/", func(g *gin.Context) {
		g.Redirect(http.StatusMovedPermanently, "/login")
	})
	// r.GET("/admin_index", service.AdminIndexHandler)
	// r.GET("/student_index", service.StudentIndexHandler)
	// r.GET("/teacher_index", service.TeacherIndexHandler)
	r.NoRoute(service.NotFoundHandler)
	r.Run(":8080")
}
