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
	"GradeManager/src/context"
	"GradeManager/src/dao"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func StudentIndexHandler(c *gin.Context) {
	var s context.StudentContext
	// check cookie
	if err := s.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	// 获取当前学生总人数，教师总人数，专业总数，学院总数
	var count_student, count_teacher, count_college, count_major string
	var wg sync.WaitGroup
	wg.Add(4)
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
	wg.Wait()

	// login status ok
	c.HTML(http.StatusOK, "student_index.html", gin.H{
		"loginer_name":  s.Info.GetName(),
		"student_count": count_student,
		"teacher_count": count_teacher,
		"college_count": count_college,
		"major_count":   count_major,
		"introduce":     "欢迎您",
		"school_title":  "成绩管理系统学生主页",
	})
}
