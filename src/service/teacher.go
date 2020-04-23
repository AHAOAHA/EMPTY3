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

	log.Infof("%+v", coursesInfo)
	result := []struct {
		CourseUid  uint64
		CourseName string
		Credit     float32
		Hour       float32
		Status     string
		Type       string
	}{}

	for _, v := range coursesInfo {
		result = append(result, struct {
			CourseUid  uint64
			CourseName string
			Credit     float32
			Hour       float32
			Status     string
			Type       string
		}{
			v.GetCourseUid(),
			v.GetName(),
			v.GetCredit(),
			v.GetHour(),
			v.GetStatus().String(),
			v.GetType().String(),
		})
	}

	rsp, err := json.Marshal(result)
	if err != nil {
		log.Error(err)
	}
	log.Infof("%v", string(rsp))
	c.JSON(http.StatusOK, string(rsp))
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

func TeacherInputScoreFirstHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "teacher_input_score_first.html", gin.H{
		"loginer_name": t.Info.GetName(),
	})
}

func TeacherQueryScoreGetHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "teacher_query_score_first.html", gin.H{
		"loginer_name": t.Info.GetName(),
	})
}

func TeacherGetCourseHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	teacher_uid := t.Info.GetTeacherUid()
	course, _ := api.GetTeacherCourseByTeacherUid(teacher_uid)
	rsp, _ := json.Marshal(course)
	c.JSON(http.StatusOK, string(rsp))
}

func TeacherGetClassHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	teacher_uid := t.Info.GetTeacherUid()
	data, _ := api.GetTeacherClassByTeacherUid(teacher_uid)
	rsp, _ := json.Marshal(data)
	c.JSON(http.StatusOK, string(rsp))
}

func TeacherInputScoreFirstPostHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	// query course class is exist.
	course_uid_str := c.Request.PostForm.Get("course_uid")
	class_uid_str := c.Request.PostForm.Get("class_uid")
	course_uid, _ := strconv.ParseUint(course_uid_str, 10, 64)
	class_uid, _ := strconv.ParseUint(class_uid_str, 10, 64)
	course_name, _ := api.GetNamebyUid(course_uid, "course", "course_uid")
	class_name, _ := api.GetNamebyUid(class_uid, "class", "class_uid")
	ok := api.IsCourseBelongClass(course_uid, class_uid)
	var percent_ok bool = true
	// 获取成绩占比信息
	course_score_percent, err := api.GetCourseScorePercentByCourseUid(course_uid)
	if err != nil {
		percent_ok = false
	}

	result := struct {
		UsualPercent uint32
		MidPercent   uint32
		EndPercent   uint32
		PercentOK    bool
		OK           bool
		CourseName   string
		ClassName    string
	}{
		course_score_percent.GetUsualPercent(),
		course_score_percent.GetMidPercent(),
		course_score_percent.GetEndPercent(),
		percent_ok,
		ok,
		course_name,
		class_name,
	}

	rsp, _ := json.Marshal(result)
	c.HTML(http.StatusOK, "teacher_input_score_second.html", gin.H{
		"loginer_name": t.Info.GetName(),
		"data":         string(rsp),
	})

}

func TeacherInputScoreThirdHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()

	course_uid_str := c.Request.PostForm.Get("course_uid")
	class_uid_str := c.Request.PostForm.Get("class_uid")
	course_uid, _ := strconv.ParseUint(course_uid_str, 10, 64)
	class_uid, _ := strconv.ParseUint(class_uid_str, 10, 64)
	// course_name := c.Request.PostForm.Get("course_name")
	// class_name := c.Request.PostForm.Get("class_name")
	// usual_percent_str := c.Request.PostForm.Get("usual_percent")
	// mid_percent_str := c.Request.PostForm.Get("mid_percent")
	// end_percent_str := c.Request.PostForm.Get("end_percent")
	// is_update_percent_str := c.Request.PostForm.Get("is_update_percent")
	// usual_percent, _ := strconv.Atoi(usual_percent_str)
	// mid_percent, _ := strconv.Atoi(mid_percent_str)
	// end_percent, _ := strconv.Atoi(end_percent_str)
	// var is_update_percent bool = true
	// if is_update_percent_str == "0" {
	// 	is_update_percent = false
	// }

	// var percent_ok bool = true
	// if (usual_percent + mid_percent + end_percent) != 100 {
	// 	percent_ok = false
	// }

	// class_uid, _ := api.GetClassUidByName(class_name)
	// course_uid, _ := api.GetCourseUidByName(course_name)
	// query course have percent data.
	// if api.IsCourseHavePercent(course_uid) == false && is_update_percent == true && percent_ok == true {
	// 	// 将该门课程的占比写入数据库
	// 	err := api.InsertCoursePercent(course_uid, uint32(usual_percent), uint32(mid_percent), uint32(end_percent), 0)
	// 	if err != nil {
	// 		log.Error(err)
	// 		return
	// 	}
	// }

	// 获取学生列表
	stu_data, _ := api.GetStudentListByClassUid(class_uid)
	student_rsp, _ := json.Marshal(stu_data)
	course_name, _ := api.GetNamebyUid(course_uid, "course", "course_uid")

	course_data := struct {
		CourseName string
		CourseUid  uint64
	}{
		course_name,
		course_uid,
	}

	course_rsp, _ := json.Marshal(course_data)

	// 组织返回值
	c.HTML(http.StatusOK, "teacher_input_score_third.html", gin.H{
		"loginer_name": t.Info.GetName(),
		"student_data": string(student_rsp),
		"course_data":  string(course_rsp),
	})
}

func TeacherQueryScoreFirstHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}
	c.Request.ParseForm()
	course_uid_str := c.Request.PostForm.Get("course_uid")
	class_uid_str := c.Request.PostForm.Get("class_uid")

	class_uid, _ := strconv.ParseUint(class_uid_str, 10, 64)
	course_uid, _ := strconv.ParseUint(course_uid_str, 10, 64)

	student_list, _ := api.GetStudentListByClassUid(class_uid)
	var result_data []struct {
		Name           string
		MidScore       float32
		EndScore       float32
		UsualScore     float32
		AcaDemicCredit float32
		Credit         float32
		Score          uint32
		Status         DataCenter.ScoreInfo_STATUS
	}

	for _, v := range student_list {
		score, err := api.GetScoreByStudentUidAndCourseUid(v.GetStudentUid(), course_uid)
		if err != nil {
			continue
		}
		result_data = append(result_data, struct {
			Name           string
			MidScore       float32
			EndScore       float32
			UsualScore     float32
			AcaDemicCredit float32
			Credit         float32
			Score          uint32
			Status         DataCenter.ScoreInfo_STATUS
		}{
			v.GetName(),
			score.GetMidtermScore(),
			score.GetEndtermScore(),
			score.GetUsualScore(),
			score.GetAcademicCredit(),
			score.GetCredit(),
			score.GetScore(),
			score.GetStatus(),
		})
	}

	rsp, _ := json.Marshal(result_data)
	log.Info(result_data)

	c.HTML(http.StatusOK, "teacher_query_score_result.html", gin.H{
		"loginer_name": t.Info.GetName(),
		"rsp_data":     string(rsp),
	})

}
