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

	// rounte
	r.Any("/", func(g *gin.Context) {
		g.Redirect(http.StatusMovedPermanently, "/login")
	})

	// login
	r.Any("/login", service.LoginHandler)
	r.POST("/sign_up", service.SignUpHandler)
	r.Any("/sign_out", service.SignOutHandler)

	// admin function
	r.Any("/admin_index", service.AdminIndexHandler)
	r.GET("/add_teacher", service.AdminAddTeacherGetHandler)
	r.GET("/add_student", service.AdminAddStudentGetHandler)
	r.GET("/add_course", service.AdminAddCourseGetHandler)
	r.GET("/add_college", service.AdminAddCollegeGetHandler)
	r.GET("/add_major", service.AdminAddMajorGetHandler)
	r.GET("/add_class", service.AdminAddClassGetHandler)
	r.GET("/admin_person_info", service.AdminInfoGetHandler)
	r.GET("/admin_teacher_manager", service.TeacherManagerHandler)
	r.GET("/admin_student_manager", service.AdminStudentManagerHandler)
	r.GET("/admin_notice_manager", service.AdminNoticeManagerGetHandler)
	r.GET("admin_delete_teacher", service.AdminDeleteTeacherHandler)
	r.GET("/admin_edit_teacher", service.AdminEditTeacherHandler)

	r.POST("/update_admin_info", service.UpdateAdminPersonInfoHandler)
	r.POST("/add_teacher", service.AdminAddTeacherPostHandler)
	r.POST("/add_student", service.AdminAddStudentPostHandler)
	r.POST("/add_college", service.AdminAddCollegePostHandler)
	r.POST("/add_major", service.AdminAddMajorPostHandler)
	r.POST("/add_class", service.AdminAddClassPostHandler)
	r.POST("/add_course", service.AdminAddCoursePostHandler)
	r.POST("/admin_teacher_manager", service.AdminTeacherManagerHandler)
	r.POST("/admin_notice_manager", service.AdminNoticeManagerPostHandler)
	r.POST("/admin_update_teacher", service.AdminUpdateTeacherPersonInfoHandler)

	// teacher
	r.GET("/teacher_index", service.TeacherIndexHandler)
	r.GET("/teacher_person_info", service.TeacherInfoGetHandler)
	r.POST("/update_teacher_info", service.UpdateTeacherPersonInfoHandler)

	// student
	r.GET("/student_index", service.StudentIndexHandler)
	r.GET("/student_person_info", service.StudentInfoGetHandler)
	r.POST("/student_person_info", service.UpdateStudentPersonInfoHandler)

	// api
	// r.GET("/query", service.QueryGetHandler)
	r.GET("/get_all_college_name", service.GetAllCollegeNameHandler)

	r.GET("/update_password", service.UpdatePasswordGetHandler)
	r.POST("/update_password", service.UpdatePasswordPostHandler)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	r.Run(":8080")
}
