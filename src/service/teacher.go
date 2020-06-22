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
	"errors"
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
		log.Error(err)
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
		"introduce":     notice.GetNotice(),
		"school_title":  notice.GetTitle(),
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
		c.JSON(http.StatusBadRequest, gin.H{
			"err_code": 1,
			"err_msg":  "登录状态错误！",
		})
		return
	}

	coursesInfo, _ := api.GetTeacherCourseByTeacherUid(t.Info.GetTeacherUid())

	result := []struct {
		CourseUid  uint64
		CourseName string
		Credit     float32
		Hour       float32
		Status     string
		Type       string
	}{}

	tempM := make(map[uint64]struct {
		CourseUid  uint64
		CourseName string
		Credit     float32
		Hour       float32
		Status     string
		Type       string
	})

	for _, v := range coursesInfo {
		tempM[v.GetCourseUid()] = struct {
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
		}
	}

	// 处理重复数据
	for _, v := range tempM {
		result = append(result, v)
	}

	c.JSON(http.StatusOK, result)
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
		ClassUid   uint64
		CourseName string
		CourseUid  uint64
	}{}
	for _, v := range result {
		data = append(data, struct {
			ClassName  string
			ClassUid   uint64
			CourseName string
			CourseUid  uint64
		}{
			v.GetName(),
			v.GetClassUid(),
			api.GetCourseNameByCourseUid(course_uid),
			course_uid,
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
	var result []struct {
		CollegeUid uint64
		CourseUid  uint64
		Name       string
		Credit     float32
		Hour       float32
		Status     string
		Type       string
	}

	for _, v := range course {
		result = append(result, struct {
			CollegeUid uint64
			CourseUid  uint64
			Name       string
			Credit     float32
			Hour       float32
			Status     string
			Type       string
		}{
			v.GetCollegeUid(),
			v.GetCourseUid(),
			v.GetName(),
			v.GetCredit(),
			v.GetHour(),
			v.GetStatus().String(),
			v.GetType().String(),
		})
	}
	rsp, _ := json.Marshal(result)
	c.JSON(http.StatusOK, string(rsp))
}

func TeacherGetClassHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err_code": 1,
			"err_msg":  "登录状态错误！",
		})
		return
	}

	courseUIDStr := c.Query("course_uid")
	courseUID, err := strconv.ParseUint(courseUIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err_code": 1,
			"err_msg":  "请求参数错误！",
		})
		return
	}

	teacherUID := t.Info.GetTeacherUid()
	data, _ := api.GetTeacherClassByTeacherUidAndCourseUid(teacherUID, courseUID)
	var result []struct {
		ClassName string
		ClassUid  uint64
	}
	for _, v := range data {
		result = append(result, struct {
			ClassName string
			ClassUid  uint64
		}{
			v.GetName(),
			v.GetClassUid(),
		})
	}
	c.JSON(http.StatusOK, result)
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

	// 获取学生列表
	stu_data, _ := api.GetStudentListByClassUid(class_uid)
	student_rsp, _ := json.Marshal(stu_data)
	course_name, _ := api.GetNamebyUid(course_uid, "course", "course_uid")

	// 获取成绩的占比情况
	course_percent, _ := api.GetCoursePercent(course_uid)

	// 获取学生的成绩列表
	var student_score []DataCenter.ScoreInfo
	for _, v := range stu_data {
		score, err := api.GetScoreByStudentUidAndCourseUid(v.GetStudentUid(), course_uid)
		if err != nil {
			student_score = append(student_score, DataCenter.ScoreInfo{})
			continue
		}

		student_score = append(student_score, score)
	}
	student_score_rsp, _ := json.Marshal(student_score)
	course_data := struct {
		CourseName   string
		CourseUid    uint64
		UsualPercent uint32
		MidPercent   uint32
		EndPercent   uint32
		Type         DataCenter.CourseScorePercentInfo_PERCENT_TYPE
	}{
		course_name,
		course_uid,
		course_percent.GetUsualPercent(),
		course_percent.GetMidPercent(),
		course_percent.GetEndPercent(),
		course_percent.GetType(),
	}

	course_rsp, _ := json.Marshal(course_data)

	// 组织返回值
	c.HTML(http.StatusOK, "teacher_input_score_third.html", gin.H{
		"loginer_name":   t.Info.GetName(),
		"student_data":   string(student_rsp),
		"course_data":    string(course_rsp),
		"students_score": string(student_score_rsp),
	})
}

func TeacherInputScoreFromQueryHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	courseUIDStr := c.Query("course_uid")
	classUIDStr := c.Query("class_uid")
	courseUID, _ := strconv.ParseUint(courseUIDStr, 10, 64)
	classUID, _ := strconv.ParseUint(classUIDStr, 10, 64)

	// 获取学生列表
	stuData, _ := api.GetStudentListByClassUid(classUID)
	studentRsp, _ := json.Marshal(stuData)
	courseName, _ := api.GetNamebyUid(courseUID, "course", "course_uid")

	// 获取成绩的占比情况
	coursePercent, _ := api.GetCoursePercent(courseUID)

	// 获取学生的成绩列表
	studentScore, _ := api.GetStudentScoreByClassUidAndCourseUid(classUID, courseUID)
	studentScoreRsp, _ := json.Marshal(studentScore)
	courseData := struct {
		CourseName   string
		CourseUid    uint64
		UsualPercent uint32
		MidPercent   uint32
		EndPercent   uint32
		Type         DataCenter.CourseScorePercentInfo_PERCENT_TYPE
	}{
		courseName,
		courseUID,
		coursePercent.GetUsualPercent(),
		coursePercent.GetMidPercent(),
		coursePercent.GetEndPercent(),
		coursePercent.GetType(),
	}

	courseRsp, _ := json.Marshal(courseData)

	// 组织返回值
	c.HTML(http.StatusOK, "teacher_input_score_third.html", gin.H{
		"loginer_name":   t.Info.GetName(),
		"student_data":   string(studentRsp),
		"course_data":    string(courseRsp),
		"students_score": string(studentScoreRsp),
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
		Status         string
		TeamYear       int32
		TeamTh         int32
		ScoreType      DataCenter.ScoreInfo_SCORE_TYPE
	}

	for _, v := range student_list {
		score, err := api.GetScoreByStudentUidAndCourseUid(v.GetStudentUid(), course_uid)
		if err != nil {
			result_data = append(result_data, struct {
				Name           string
				MidScore       float32
				EndScore       float32
				UsualScore     float32
				AcaDemicCredit float32
				Credit         float32
				Score          uint32
				Status         string
				TeamYear       int32
				TeamTh         int32
				ScoreType      DataCenter.ScoreInfo_SCORE_TYPE
			}{
				v.GetName(),
				0,
				0,
				0,
				0,
				0,
				0,
				score.GetType().String(),
				score.GetTeamYear(),
				score.GetTeamTh(),
				score.GetScoreType(),
			})
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
			Status         string
			TeamYear       int32
			TeamTh         int32
			ScoreType      DataCenter.ScoreInfo_SCORE_TYPE
		}{
			v.GetName(),
			score.GetMidtermScore(),
			score.GetEndtermScore(),
			score.GetUsualScore(),
			score.GetAcademicCredit(),
			score.GetCredit(),
			score.GetScore(),
			score.GetType().String(),
			score.GetTeamYear(),
			score.GetTeamTh(),
			score.GetScoreType(),
		})
	}

	rsp, _ := json.Marshal(result_data)
	log.Info(result_data)

	c.HTML(http.StatusOK, "teacher_query_score_result.html", gin.H{
		"loginer_name": t.Info.GetName(),
		"rsp_data":     string(rsp),
	})

}

func TeacherInputScoreHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	buf := make([]byte, 1024000)
	n, _ := c.Request.Body.Read(buf)

	body_data := string(buf[0:n])

	var err error
	var err_code int = 0
	var body_m map[string]interface{}
	_ = json.Unmarshal([]byte(body_data), &body_m)
	log.Info(body_m)
	cmd := body_m["Cmd"]
	if cmd == "save" {
		// 保存命令
		// 保存比例
		err = SavePercent(body_m, 0)
		if err != nil {
			log.Error(err)
			err_code = 1001
		}
		err = SaveStudentScore(body_m, 0)
		if err != nil {
			log.Error(err)
			err_code = 1001
		}
	} else if cmd == "submit" {
		// 提交
		err = SavePercent(body_m, 0)
		if err != nil {
			log.Error(err)
			err_code = 1002
		}
		err = SaveStudentScore(body_m, 1)
		if err != nil {
			log.Error(err)
			err_code = 1002
		}

	} else {
		// 错误
	}
	c.JSON(http.StatusOK, gin.H{
		"err_code": err_code,
	})

}

func GetRealCreditAndAcaCredit(courseUID uint64, score int, scoreType int) (float32, float32) {
	courseInfo, err := api.GetCourseByCourseUid(courseUID)
	if err != nil {
		log.Error(err)
		return 0, 0
	}
	log.Info(score)
	if scoreType == 0 {
		// 分数制
		if score < 60 && courseInfo.GetCredit() > 1 {
			return 1.0, 1.0
		}
		return float32((float32(score) / 100) * courseInfo.GetCredit()), float32(((float32(score) / 100) * courseInfo.GetCredit()) * 0.6)

	} else {
		// 分级制
		if score >= 60 {
			return courseInfo.GetCredit(), courseInfo.GetCredit() * 0.6
		}
	}

	return 0, 0
}

func SaveStudentScore(body map[string]interface{}, save_type int) error {
	score_type := body["ScoreType"]
	score_data := body["Data"].([]interface{})

	teamYearStr := body["TeamYear"].(string)
	teamThStr := body["TeamTh"].(string)

	var courseUIDStr string
	var classUID uint64

	if score_type == "0" {
		// 分数制
		log.Info(score_data)
		for _, v := range score_data {
			log.Info(v)
			student_data := v.(map[string]interface{})
			usual_score, _ := strconv.ParseFloat(student_data["UsualScore"].(string), 32)
			mid_score, _ := strconv.ParseFloat(student_data["MidScore"].(string), 32)
			end_score, _ := strconv.ParseFloat(student_data["EndScore"].(string), 32)
			scoreF, _ := strconv.ParseFloat(student_data["Score"].(string), 32)
			score := int(scoreF)

			// 计算学分绩点
			courseUID, _ := strconv.ParseUint(student_data["CourseUid"].(string), 10, 64)
			credit, acaCredit := GetRealCreditAndAcaCredit(courseUID, score, 0)

			// 保存备份信息
			courseUIDStr = student_data["CourseUid"].(string)
			studentUIDStr := student_data["StudentUid"].(string)
			studentUID, _ := strconv.ParseUint(studentUIDStr, 10, 64)
			stuInfo, _ := api.GetStudentByStudentUid(studentUID)
			classUID = stuInfo.GetClassUid()

			// 逐条保存
			// 查看数据是否存在
			var is_exist bool = true
			m, err := dao.DataBase.Queryf("select * from `score` where `student_uid`='%s' and `course_uid`='%s'", student_data["StudentUid"].(string), student_data["CourseUid"].(string))
			if err != nil || len(m) != 1 {
				is_exist = false
			}

			if is_exist == true {
				if string(m[0]["type"].([]uint8)) == "0" {
					// 可以修改
					err := dao.DataBase.Execf("update `score` set `usual_score`='%f', `midterm_score`='%f', `endterm_score`='%f', `score`='%d', `type`='%d', `score_type`='0', `credit`='%f', `academic_credit`='%f', `team_year`='%s', `team_th`='%s' where `student_uid`='%s' and `course_uid`='%s'", usual_score, mid_score, end_score, score, save_type, credit, acaCredit, teamYearStr, teamThStr, student_data["StudentUid"].(string), student_data["CourseUid"].(string))
					if err != nil {
						return err
					}
				} else {
					return errors.New("type is submit")
				}
			} else {
				err := dao.DataBase.Execf("insert into `score`(`student_uid`, `course_uid`, `usual_score`, `midterm_score`, `endterm_score`,`score`, `type`, `score_type`, `credit`, `academic_credit`, `team_year`, `team_th`) values ('%s', '%s', '%f', '%f', '%f', '%d', '%d', '0', '%f', '%f', `%s`, `%s`)", student_data["StudentUid"].(string), student_data["CourseUid"].(string), usual_score, mid_score, end_score, score, save_type, credit, acaCredit, teamYearStr, teamThStr)
				if err != nil {
					return err
				}
			}
		}
	} else if score_type == "1" {
		// 分级制
		for _, v := range score_data {
			log.Warn(v)
			student_data := v.(map[string]interface{})
			usual_score, _ := strconv.ParseFloat(student_data["UsualScore"].(string), 32)
			mid_score, _ := strconv.ParseFloat(student_data["MidScore"].(string), 32)
			end_score, _ := strconv.ParseFloat(student_data["EndScore"].(string), 32)
			percent := body["Percent"].(map[string]interface{})
			usual_percent, _ := strconv.Atoi(percent["UsualPercent"].(string))
			mid_percent, _ := strconv.Atoi(percent["MidPercent"].(string))
			end_percent, _ := strconv.Atoi(percent["EndPercent"].(string))
			score := int(usual_score*float64(usual_percent)/100 + mid_score*float64(mid_percent)/100 + end_score*float64(end_percent)/100)
			log.Info("score:", score)
			// 逐条保存

			// 计算学分绩点
			courseUID, _ := strconv.ParseUint(student_data["CourseUid"].(string), 10, 64)
			credit, acaCredit := GetRealCreditAndAcaCredit(courseUID, score, 1)

			var is_exist bool = true
			m, err := dao.DataBase.Queryf("select * from `score` where `student_uid`='%s' and `course_uid`='%s'", student_data["StudentUid"].(string), student_data["CourseUid"].(string))
			if err != nil || len(m) != 1 {
				is_exist = false
			}

			if is_exist == true {
				if string(m[0]["type"].([]uint8)) == "0" {
					// 可以修改
					err := dao.DataBase.Execf("update `score` set `usual_score`='%f', `midterm_score`='%f', `endterm_score`='%f', `score`='%d', `type`='%d', `score_type`='1', `credit`='%f', `academic_credit`='%f', `team_year`='%s', `team_th`='%s' where `student_uid`='%s' and `course_uid`='%s'", usual_score, mid_score, end_score, score, save_type, credit, acaCredit, teamYearStr, teamThStr, student_data["StudentUid"].(string), student_data["CourseUid"].(string))
					if err != nil {
						return err
					}
				} else {
					return errors.New("type is Sumbit")
				}
			} else {
				err := dao.DataBase.Execf("insert into `score`(`student_uid`, `course_uid`, `usual_score`, `midterm_score`, `endterm_score`,`score`, `type`, `score_type`, `credit`, `academic_credit`, `team_year`, `team_th`) values ('%s', '%s', '%f', '%f', '%f', '%d', '%d', '1', '%f', '%f', '%s', '%s')", student_data["StudentUid"].(string), student_data["CourseUid"].(string), usual_score, mid_score, end_score, score, save_type, credit, acaCredit, teamYearStr, teamThStr)
				if err != nil {
					return err
				}
			}
		}
	} else {
		// err
	}

	// 提交成绩时修改课程状态
	if save_type == 1 {
		// 更新课程状态
		err := dao.DataBase.Execf("update `student_course` set `status`='%d' where `class_uid`='%d' and `course_uid`='%s'", DataCenter.StudentCourseInfo_DONE, classUID, courseUIDStr)
		if err != nil {
			log.Error(err.Error())
		}
	}
	return nil
}

func SavePercent(body map[string]interface{}, save_type int) error {
	var is_exist bool = true
	percent := body["Percent"].(map[string]interface{})
	usual_percent := percent["UsualPercent"].(string)
	mid_percent := percent["MidPercent"].(string)
	end_percent := percent["EndPercent"].(string)
	course_uid_str := percent["CourseUid"].(string)
	m, err := dao.DataBase.Queryf("select * from `course_score_percent` where `course_uid`=%s", course_uid_str)
	if err != nil || len(m) != 1 {
		is_exist = false
	}

	if is_exist == true {
		if string(m[0]["type"].([]uint8)) == "1" {
			return errors.New("type is submit")
		} else {
			err := dao.DataBase.Execf("update `course_score_percent` set `usual_percent`='%s', `mid_percent`='%s', `end_percent`='%s', `type`='%d' where `course_uid`='%s'", usual_percent, mid_percent, end_percent, save_type, course_uid_str)
			if err != nil {
				return err
			}
		}
	} else {
		err := dao.DataBase.Execf("insert into `course_score_percent` (`course_uid`, `usual_percent`, `mid_percent`, `end_percent`, `type`) values('%s', '%s', '%s', '%s', '%d');", course_uid_str, usual_percent, mid_percent, end_percent, save_type)
		if err != nil {
			return err
		}
	}
	return nil
}

func TeacherCourseReachFirstHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "teacher_course_reach_first.html", gin.H{
		"loginer_name": t.Info.GetName(),
	})
}

func TeacherCourseReachSecondHandler(c *gin.Context) {
	var t context.TeacherContext
	// check cookie
	if err := t.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	courseUIDStr := c.Request.PostForm.Get("course_uid")
	courseUID, _ := strconv.ParseUint(courseUIDStr, 10, 64)
	classUIDStr := c.Request.PostForm.Get("class_uid")
	classUID, _ := strconv.ParseUint(classUIDStr, 10, 64)

	data, _ := api.GetStudentScoreByClassUidAndCourseUid(classUID, courseUID)

	var result []struct {
		CourseUid   uint64
		CourseName  string
		StudentUid  uint64
		StudentName string
		ClassUId    uint64
		ClassName   string
		UsualScore  float32
		MidScore    float32
		EndScore    float32
		Score       uint32
	}

	for _, v := range data {
		if v.GetType() == DataCenter.ScoreInfo_SAVE {
			//continue
		}
		courseName, _ := api.GetNamebyUid(v.GetCourseUid(), "course", "course_uid")
		studentName, _ := api.GetNamebyUid(v.GetStudentUid(), "student", "student_uid")
		className, _ := api.GetNamebyUid(classUID, "class", "class_uid")
		result = append(result, struct {
			CourseUid   uint64
			CourseName  string
			StudentUid  uint64
			StudentName string
			ClassUId    uint64
			ClassName   string
			UsualScore  float32
			MidScore    float32
			EndScore    float32
			Score       uint32
		}{
			v.GetCourseUid(),
			courseName,
			v.GetStudentUid(),
			studentName,
			classUID,
			className,
			v.GetUsualScore(),
			v.GetMidtermScore(),
			v.GetEndtermScore(),
			v.GetScore(),
		})
	}

	percentData, _ := api.GetCourseScorePercentByCourseUid(courseUID)
	percentResult, _ := json.Marshal(percentData)

	jsonResult, _ := json.Marshal(result)
	c.HTML(http.StatusOK, "teacher_course_reach_third.html", gin.H{
		"loginer_name":  t.Info.GetName(),
		"student_score": string(jsonResult),
		"percent":       string(percentResult),
	})
}
