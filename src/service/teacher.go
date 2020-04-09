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
	"GradeManager/src/api"
	"GradeManager/src/context"
	"GradeManager/src/dao"
	DataCenter "GradeManager/src/proto"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func TeacherIndexHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	// 获取当前学生总人数，教师总人数，专业总数，学院总数
	var count_student, count_teacher, count_college, count_major string
	var notice DataCenter.NoticeInfo
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		m, err := dao.DataBase.Queryf("select count(*) from `student`")
		if err == nil || len(m) != 0 {
			count_student = string(m[0]["count(*)"].([]uint8))
		}
		wg.Done()
	}()
	go func() {
		m, err := dao.DataBase.Queryf("select count(*) from `teacher`")
		if err == nil || len(m) != 0 {
			count_teacher = string(m[0]["count(*)"].([]uint8))
		}
		wg.Done()
	}()
	go func() {
		m, err := dao.DataBase.Queryf("select count(*) from `college`")
		if err == nil || len(m) != 0 {
			count_college = string(m[0]["count(*)"].([]uint8))
		}
		wg.Done()
	}()
	go func() {
		m, err := dao.DataBase.Queryf("select count(*) from `major`")
		if err == nil || len(m) != 0 {
			count_major = string(m[0]["count(*)"].([]uint8))
		}
		wg.Done()
	}()
	go func() {
		notice, _ = api.GetNotice()
		wg.Done()
	}()
	wg.Wait()

	// login status ok
	c.HTML(http.StatusOK, "teacher_index.html", gin.H{
		"loginer_name":  t.Info.GetName(),
		"student_count": count_student,
		"teacher_count": count_teacher,
		"college_count": count_college,
		"major_count":   count_major,
		"introduce":     notice.GetTitle(),
		"school_title":  notice.GetNotice(),
	})
}

func TeacherCourseQueryGetHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "teacher_course_query_result.html", gin.H{
		"loginer_name": t.Info.GetName(),
	})
}

func TeacherGetTeacherCoursesHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	coursesInfo, err := api.GetTeacherCourseByTeacherUid(t.Info.GetTeacherUid())
	if err != nil {
		// err
	}

	log.Info(coursesInfo)

	data, _ := json.Marshal(coursesInfo)
	c.JSON(http.StatusOK, string(data))
}

// json format, class name, course name, and operator info.
func TeacherGetCourseClassHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	course_uid_str := c.Query("course_uid")
	course_uid, _ := strconv.ParseUint(course_uid_str, 10, 64)
	result, _ := api.GetTeacherCourseClass(t.Info.GetTeacherUid(), course_uid)
	data := []struct {
		ClassName  string
		CourseName string
	}{}
	for _, v := range result {
		data = append(data, struct {
			ClassName  string
			CourseName string
		}{
			v.GetName(),
			api.GetCourseNameByCourseUid(course_uid),
		})
	}
	rsp, _ := json.Marshal(data)

	c.HTML(http.StatusOK, "teacher_course_class_result.html", gin.H{
		"loginer_name": t.Info.GetName(),
		"data":         string(rsp),
	})

}
