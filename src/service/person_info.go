/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: person_info.go
 * Author: ahaoozhang
 * Date: 2020-03-16 16:40:07 (Monday)
 * Describe:
 ******************************************************************/
package service

import (
	"GradeManager/src/api"
	"GradeManager/src/context"
	"GradeManager/src/dao"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AdminInfoGetHandler(c *gin.Context) {
	// checkout cookie
	var admin context.AdminContext
	if err := admin.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}
	log.Info(admin.Info.GetMail())
	c.HTML(http.StatusOK, "admin_person_info.html", gin.H{
		"name":        admin.Info.GetUser(),
		"mail":        admin.Info.GetMail(),
		"create_time": time.Unix(int64(admin.Info.GetCreateTime()), 0).Format("2006-01-02 03:04:05 PM"),
		"expr_time":   time.Unix(int64(admin.Info.GetExprTime()), 0).Format("2006-01-02 03:04:05 PM"),
		"login_ip":    c.ClientIP(),
	})
}

func UpdateAdminPersonInfoHandler(c *gin.Context) {
	// check cookie
	var admin context.AdminContext
	if err := admin.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}
	c.Request.ParseForm()
	name := c.Request.PostForm.Get("name")
	mail := c.Request.PostForm.Get("mail")

	err := dao.DataBase.Execf("update `admin` set `user`='%s', `mail`='%s' where `user`='%s'", name, mail, admin.Info.GetUser())
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusBadGateway, "502.html", nil)
		return
	}
	c.SetCookie("user_cookie", "out", 10, "/", "", false, true)
	c.Redirect(http.StatusFound, "/login")
}

func TeacherInfoGetHandler(c *gin.Context) {
	var t context.TeacherContext
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "teacher_person_info.html", gin.H{
		"teacher_uid":  t.Info.GetTeacherUid(),
		"name":         t.Info.GetName(),
		"sex":          t.Info.GetSex(),
		"NRIC":         t.Info.GetNRIC(),
		"college_name": t.Info.GetCollegeUid(),
		"status":       t.Info.GetStatus(),
		"create_time":  t.Info.GetCreateTime(),
		"loginer_name": t.Info.GetName(),
		"login_ip":     c.ClientIP(),
	})
}

func UpdateTeacherPersonInfoHandler(c *gin.Context) {
	var t context.TeacherContext
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	name := c.Request.PostForm.Get("name")
	sex := c.Request.PostForm.Get("sex")

	err := dao.DataBase.Execf("update `teacher` set `name`='%s', sex='%s' where `teacher_uid`='%s'", name, sex, t.Info.GetTeacherUid())
	if err != nil {
		c.HTML(http.StatusBadGateway, "502.html", nil)
		return
	}

	c.SetCookie("user_cookie", "out", 10, "/", "", false, true)
	c.Redirect(http.StatusFound, "/login")
}

func StudentInfoGetHandler(c *gin.Context) {
	var s context.StudentContext
	if err := s.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	collegeName, _ := api.GetNamebyUid(s.Info.GetCollegeUid(), "college", "college_uid")
	majorName, _ := api.GetNamebyUid(s.Info.GetMajorUid(), "major", "major_uid")
	className, _ := api.GetNamebyUid(s.Info.GetClassUid(), "class", "class_uid")

	c.HTML(http.StatusOK, "student_person_info.html", gin.H{
		"student_uid":  s.Info.GetStudentUid(),
		"name":         s.Info.GetName(),
		"sex":          s.Info.GetSex(),
		"NRIC":         s.Info.GetNRIC(),
		"college_name": collegeName,
		"major_name":   majorName,
		"class_name":   className,
		"status":       s.Info.GetStatus(),
		"create_time":  s.Info.GetCreateTime(),
		"loginer_name": s.Info.GetName(),
		"login_ip":     c.ClientIP(),
	})
}

func UpdateStudentPersonInfoHandler(c *gin.Context) {
	var s context.StudentContext
	if err := s.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	name := c.Request.PostForm.Get("name")
	sex := c.Request.PostForm.Get("sex")

	err := dao.DataBase.Execf("update `student` set `name`='%s', sex='%s' where `student_uid`='%d'", name, sex, s.Info.GetStudentUid())
	if err != nil {
		c.HTML(http.StatusBadGateway, "502.html", nil)
		return
	}

	c.SetCookie("user_cookie", "out", 10, "/", "", false, true)
	c.Redirect(http.StatusFound, "/login")
}
