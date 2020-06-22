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
	"GradeManager/src/api"
	"GradeManager/src/context"
	"GradeManager/src/dao"
	DataCenter "GradeManager/src/proto"
	"encoding/json"
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
	c.HTML(http.StatusOK, "student_index.html", gin.H{
		"loginer_name":  s.Info.GetName(),
		"student_count": count_student,
		"teacher_count": count_teacher,
		"college_count": count_college,
		"major_count":   count_major,
		"introduce":     notice.GetNotice(),
		"school_title":  notice.GetTitle(),
	})
}

// 学生成绩查询
func StudentScoreQueryHandler(c *gin.Context) {
	var s context.StudentContext
	// check cookie
	if err := s.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	result, _ := api.GetStudentSubmitScoreByStudentUid(s.Info.GetStudentUid())
	var rsp_data []struct {
		StudentUid     uint64
		StudentName    string
		CourseName     string
		MidScore       float32
		EndScore       float32
		UsualScore     float32
		Score          uint32
		AcademicCredit float32
		Credit         float32
		ScoreType      DataCenter.ScoreInfo_SCORE_TYPE
		TeamYear       int32
		TeamTh         int32
	}

	for _, v := range result {
		student_name, _ := api.GetNamebyUid(v.GetStudentUid(), "student", "student_uid")
		course_name, _ := api.GetNamebyUid(v.GetCourseUid(), "course", "course_uid")
		rsp_data = append(rsp_data, struct {
			StudentUid     uint64
			StudentName    string
			CourseName     string
			MidScore       float32
			EndScore       float32
			UsualScore     float32
			Score          uint32
			AcademicCredit float32
			Credit         float32
			ScoreType      DataCenter.ScoreInfo_SCORE_TYPE
			TeamYear       int32
			TeamTh         int32
		}{
			v.GetStudentUid(),
			student_name,
			course_name,
			v.GetMidtermScore(),
			v.GetEndtermScore(),
			v.GetUsualScore(),
			v.GetScore(),
			v.GetAcademicCredit(),
			v.GetCredit(),
			v.GetScoreType(),
			v.GetTeamYear(),
			v.GetTeamTh(),
		})
	}
	rsp, _ := json.Marshal(rsp_data)
	c.HTML(http.StatusOK, "student_score_query.html", gin.H{
		"loginer_name": s.Info.GetName(),
		"score_data":   string(rsp),
	})
}

func StudentQueryCourseHandler(c *gin.Context) {
	var s context.StudentContext
	// check cookie
	if err := s.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	studentUID := s.Info.GetStudentUid()
	studentInfo, _ := api.GetStudentByStudentUid(studentUID)
	courses, _ := api.GetCourseByClassUid(studentInfo.GetClassUid())
	var result []struct {
		CourseName  string
		CollegeName string
		Hour        float32
		Credit      float32
		CourseType  string
		Status      string
	}

	for _, v := range courses {
		collegeName, _ := api.GetNamebyUid(v.GetCollegeUid(), "college", "college_uid")
		result = append(result, struct {
			CourseName  string
			CollegeName string
			Hour        float32
			Credit      float32
			CourseType  string
			Status      string
		}{
			v.GetName(),
			collegeName,
			v.GetHour(),
			v.GetCredit(),
			v.GetType().String(),
			v.GetStatus().String(),
		})
	}

	rsp, _ := json.Marshal(result)
	c.HTML(http.StatusOK, "student_query_course.html", gin.H{
		"loginer_name": s.Info.GetName(),
		"course_data":  string(rsp),
	})
}
