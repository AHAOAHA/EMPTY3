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
	"GradeManager/src/common"
	_ "GradeManager/src/config"
	"GradeManager/src/logFileWriter"
	"GradeManager/src/service"
	"net/http"
	"time"

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

	fileWriter, _ := logFileWriter.CeateNewLogFileWriter()
	log.SetOutput(fileWriter)

	log.SetLevel(log.InfoLevel)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Static("/static", "./views/static")
	r.LoadHTMLGlob("/home/ahaoozhang/dev_code/GradeManager/views/templates/*")

	// rounte
	r.Any("/", func(g *gin.Context) {
		g.Redirect(http.StatusFound, "/login")
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
	r.GET("/admin_delete_student", service.AdminDeleteStudentHandler)
	r.GET("/admin_edit_student", service.AdminEditStudentHandler)
	r.GET("/admin_course_manager", service.AdminCourseGetHandler)
	r.GET("/admin_add_course", service.AdminAddCourseHandler)
	r.GET("/admin_add_course_second", service.AdminAddCourseSecondHandler)
	r.GET("/admin_score_manager", service.AdminScoreManagerGetHandler)

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
	r.POST("/admin_student_manager", service.AdminStudentManagerPostHandler)
	r.POST("/admin_update_student", service.AdminUpdateStudentPersonInfoHandler)
	r.POST("/admin_query_class_course", service.AdminQueryClassCourseHandler)
	r.POST("/admin_add_course", service.AdminAddCourseHandler)
	r.POST("/admin_get_student_score", service.AdminGetStudentScoreByStudentUidHandler)
	r.POST("/admin_input_score", service.AdminChangeStudentScoreHandler)

	// teacher
	r.GET("/teacher_index", service.TeacherIndexHandler)
	r.GET("/teacher_person_info", service.TeacherInfoGetHandler)
	r.GET("/teacher_course_query", service.TeacherCourseQueryGetHandler)
	r.GET("/get_teacher_courses", service.TeacherGetTeacherCoursesHandler)
	r.GET("/get_teacher_course_class", service.TeacherGetCourseClassHandler)
	r.GET("/teacher_input_score", service.TeacherInputScoreFirstHandler)
	r.GET("/teacher_query_score", service.TeacherQueryScoreGetHandler)
	r.GET("/get_teacher_course", service.TeacherGetCourseHandler)
	r.GET("/get_teacher_class", service.TeacherGetClassHandler)
	r.GET("/teacher_course_reach", service.TeacherCourseReachFirstHandler)
	r.GET("/teacher_input_score_from_query", service.TeacherInputScoreFromQueryHandler)
	r.POST("/update_teacher_info", service.UpdateTeacherPersonInfoHandler)
	// r.POST("/teacher_input_score_first", service.TeacherInputScoreFirstPostHandler)
	r.POST("/teacher_input_score_third", service.TeacherInputScoreThirdHandler)
	r.POST("/teacher_query_score_first", service.TeacherQueryScoreFirstHandler)
	r.POST("/teacher_input_score", service.TeacherInputScoreHandler)
	r.POST("/teacher_course_reach_second", service.TeacherCourseReachSecondHandler)

	// student
	r.GET("/student_index", service.StudentIndexHandler)
	r.GET("/student_person_info", service.StudentInfoGetHandler)
	r.GET("/student_score_query", service.StudentScoreQueryHandler)
	r.GET("/student_query_course", service.StudentQueryCourseHandler)
	r.POST("/student_person_info", service.UpdateStudentPersonInfoHandler)

	// api
	// r.GET("/query", service.QueryGetHandler)
	r.GET("/get_all_college_name", service.GetAllCollegeNameHandler)
	r.GET("/get_all_major_name", service.GetAllMajorNameHandler)
	r.GET("/get_all_class_name", service.GetAllClassNameHandler)
	r.GET("/get_all_course_name", service.GetAllCourseInfoHandler)
	r.GET("/get_all_teacher_name", service.GetAllTeacherInfoHandler)
	r.GET("location-to-addr", service.LocationToAddrHandler)
	r.GET("ip-to-addr", service.IpToAddrHandler)

	r.GET("/update_password", service.UpdatePasswordGetHandler)
	r.GET("/get_teacher_info", service.GetTeacherInfoHandler)
	r.GET("/get_course_info", service.GetCourseInfoByCollegeUidHandler)
	r.POST("/update_password", service.UpdatePasswordPostHandler)

	//debug
	r.GET("/debug", service.DebugHandler)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	common.SendTextToWechat("服务重启", time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
	r.Run(":8080")
}
