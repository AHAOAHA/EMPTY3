/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: admin.go
 * Author: ahaoozhang
 * Date: 2020-03-09 22:34:51 (Monday)
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

func AdminIndexHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
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
		notice_tmp, err := api.GetNotice()
		if err == nil {
			log.Info("notice is nil")
		}

		notice = notice_tmp
		wg.Done()
	}()
	wg.Wait()
	c.HTML(http.StatusOK, "admin_index.html", gin.H{
		"title":         "login",
		"student_count": count_student,
		"teacher_count": count_teacher,
		"college_count": count_college,
		"major_count":   count_major,
		"loginer_name":  a.Info.GetUser(),
		"introduce":     notice.GetNotice(),
		"school_title":  notice.GetTitle(),
	})
}

func AdminAddTeacherGetHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "form_add_teacher.html", gin.H{
		"err_code": "ok",
	})
}

func AdminAddStudentGetHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}
	c.HTML(http.StatusOK, "form_add_student.html", gin.H{
		"err_code": "ok",
	})
}

// 添加教师
func AdminAddTeacherPostHandler(c *gin.Context) {
	var a context.AdminContext
	cookie, err := c.Request.Cookie("user_cookie")
	if err != nil {
		// cookie不存在，跳转401
		c.HTML(http.StatusBadRequest, "401.html", nil)
		log.Error(err)
		return
	}
	err = a.Detcry(cookie.Value)
	if err != nil {
		// cookie内容验证失败
		log.Error(err)
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	for k, v := range c.Request.PostForm {
		if len(v) == 0 {
			// 表单内容为空，拒绝注册
			log.Warn("Form value is nil, ", k)
			return
		}
	}

	// 获取学院id
	m, err := dao.DataBase.Queryf("select `college_uid` from `college` where name='%s'", c.Request.PostForm.Get("college_name"))
	if err != nil || len(m) == 0 {
		log.Warn("college not exist err")
		c.JSON(http.StatusOK, gin.H{
			"err_msg": "college name err",
		})
		return
	}

	college_uid := string(m[0]["college_uid"].([]uint8))
	form_data := c.Request.PostForm
	err = dao.DataBase.Execf("insert into `teacher` (`teacher_uid`, `college_uid`, `password`, `name`, `sex`, `NRIC`) values('%s', '%s', '%s', '%s', '%s', '%s' )",
		form_data.Get("teacher_uid"), college_uid, form_data.Get("password"), form_data.Get("username"), form_data.Get("sex"), form_data.Get("NRIC"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"err_msg": err,
		})
		return
	}

	c.HTML(http.StatusOK, "success.html", gin.H{
		"second": 3,
		"url":    "admin_index",
	})
}

// 添加学生
func AdminAddStudentPostHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	for k, v := range c.Request.PostForm {
		if len(v) == 0 {
			// 表单内容为空，拒绝注册
			log.Warn("Form value is nil, ", k)
			return
		}
	}

	// 获取学院id 班级id 专业id
	var ok bool = true
	var wg sync.WaitGroup
	var college_uid, major_uid, class_uid string
	wg.Add(3)
	go func() {
		m, err := dao.DataBase.Queryf("select `college_uid` from `college` where name='%s'", c.Request.PostForm.Get("college_name"))
		if err != nil || len(m) == 0 {
			ok = false
		}
		college_uid = string(m[0]["college_uid"].([]uint8))
		wg.Done()
	}()
	go func() {
		m, err := dao.DataBase.Queryf("select `class_uid` from `class` where name='%s'", c.Request.PostForm.Get("class_name"))
		if err != nil || len(m) == 0 {
			ok = false
		}
		class_uid = string(m[0]["class_uid"].([]uint8))
		wg.Done()
	}()
	go func() {
		m, err := dao.DataBase.Queryf("select `major_uid` from `major` where name='%s'", c.Request.PostForm.Get("major_name"))
		if err != nil || len(m) == 0 {
			ok = false
		}
		major_uid = string(m[0]["major_uid"].([]uint8))
		wg.Done()
	}()
	wg.Wait()

	if !ok {
		// 输入内容错误，拒绝创建
		c.JSON(http.StatusOK, gin.H{
			"err_msg": "format err",
		})
		return
	}
	formdata := c.Request.PostForm

	err := dao.DataBase.Execf("insert into `student`(`student_uid`, `class_uid`, `college_uid`, `major_uid`, `password`, `name`, `sex`, `NRIC`) values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')",
		formdata.Get("student_uid"), class_uid, college_uid, major_uid, formdata.Get("password"), formdata.Get("name"), formdata.Get("sex"), formdata.Get("NRIC"))
	if err != nil {
		// 插入失败
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"err_msg": "insert student err",
		})
		return
	}

	c.HTML(http.StatusOK, "success.html", gin.H{
		"second": 3,
		"url":    "admin_index",
	})
}

func AdminAddCollegeGetHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "form_add_college.html", gin.H{
		"err_msg": "ok",
	})
}

func AdminAddCollegePostHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}
	c.Request.ParseForm()
	formdata := c.Request.PostForm
	for k, v := range formdata {
		log.Debug(k, ":", v)
		if len(v) == 0 {
			log.Error("form format err", k)
			c.JSON(http.StatusOK, gin.H{
				"err_msg": "form format err",
			})
			return
		}
	}

	err := dao.DataBase.Execf("insert into `college`(`college_uid`, `name`) values('%s', '%s')", formdata.Get("college_uid"), formdata.Get("name"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"err_msg": err,
		})
		return
	}

	c.HTML(http.StatusOK, "success.html", gin.H{
		"second": 3,
		"url":    "admin_index",
	})

}

func AdminAddMajorGetHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "form_add_major.html", gin.H{
		"err_msg": "ok",
	})
}

func AdminAddMajorPostHandler(c *gin.Context) {
	var a context.AdminContext
	cookie, err := c.Request.Cookie("user_cookie")
	if err != nil {
		// cookie不存在，跳转401
		c.HTML(http.StatusBadRequest, "401.html", nil)
		log.Error(err)
		return
	}
	err = a.Detcry(cookie.Value)
	if err != nil {
		// cookie内容验证失败
		log.Error(err)
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	formdata := c.Request.PostForm
	if err := postFormIsValid(c); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_mag": err,
		})
		return
	}
	m, err := dao.DataBase.Queryf("select `college_uid` from `college` where `name`='%s'", formdata.Get("college_name"))
	if err != nil || len(m) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_msg": err,
		})
		return
	}
	college_uid := string(m[0]["college_uid"].([]uint8))
	err = dao.DataBase.Execf("insert into `major`(`major_uid`, `college_uid`, `name`)values('%s', '%s', '%s')", formdata.Get("major_uid"), college_uid, formdata.Get("name"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_msg": err,
		})
		return
	}
	c.HTML(http.StatusOK, "success.html", gin.H{
		"second": 3,
		"url":    "admin_index",
	})
}

func AdminAddClassGetHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "form_add_class.html", gin.H{
		"err_msg": "ok",
	})
}

func AdminAddClassPostHandler(c *gin.Context) {
	var a context.AdminContext
	cookie, err := c.Request.Cookie("user_cookie")
	if err != nil {
		// cookie不存在，跳转401
		c.HTML(http.StatusBadRequest, "401.html", nil)
		log.Error(err)
		return
	}
	err = a.Detcry(cookie.Value)
	if err != nil {
		// cookie内容验证失败
		log.Error(err)
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	formdata := c.Request.PostForm
	if err := postFormIsValid(c); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_mag": err,
		})
		return
	}
	var ok bool = true
	var college_uid, major_uid string
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		m, err := dao.DataBase.Queryf("select `college_uid` from `college` where `name`='%s'", formdata.Get("college_name"))
		if err != nil || len(m) == 0 {
			ok = false
		} else {
			college_uid = string(m[0]["college_uid"].([]uint8))
		}
		wg.Done()
	}()
	go func() {
		m, err := dao.DataBase.Queryf("select `major_uid` from `major` where `name`='%s'", formdata.Get("major_name"))
		if err != nil || len(m) == 0 {
			ok = false
		} else {
			major_uid = string(m[0]["major_uid"].([]uint8))
		}
		wg.Done()
	}()
	wg.Wait()

	if !ok {
		log.Error("there err")
	}

	err = dao.DataBase.Execf("insert into `class`(`class_uid`, `college_uid`, `major_uid`, `name`) values ('%s', '%s', '%s', '%s')",
		formdata.Get("class_uid"), college_uid, major_uid, formdata.Get("name"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_msg": err,
		})
	}

	c.HTML(http.StatusOK, "success.html", gin.H{
		"second": 3,
		"url":    "admin_index",
	})
}

func AdminAddCourseGetHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "form_add_course.html", gin.H{
		"err_msg": "ok",
	})
}

func AdminAddCoursePostHandler(c *gin.Context) {
	var a context.AdminContext
	cookie, err := c.Request.Cookie("user_cookie")
	if err != nil {
		// cookie不存在，跳转401
		c.HTML(http.StatusBadRequest, "401.html", nil)
		log.Error(err)
		return
	}
	err = a.Detcry(cookie.Value)
	if err != nil {
		// cookie内容验证失败
		log.Error(err)
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	formdata := c.Request.PostForm
	if err = postFormIsValid(c); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_mag": err,
		})
		return
	}

	m, err := dao.DataBase.Queryf("select `college_uid` from `college` where `name`='%s'", formdata.Get("college_name"))
	if err != nil || len(m) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_msg": err,
		})
	}
	college_uid := string(m[0]["college_uid"].([]uint8))
	err = dao.DataBase.Execf("insert into `course`(`course_uid`, `college_uid`, `name`, `credit`, `hour`, `type`) values ('%s', '%s', '%s', '%s', '%s', '%s')",
		formdata.Get("course_uid"), college_uid, formdata.Get("name"), formdata.Get("credit"), formdata.Get("hour"), formdata.Get("type"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_msg": err,
		})
	}
	c.HTML(http.StatusOK, "success.html", gin.H{
		"second": 3,
		"url":    "admin_index",
	})
}

func postFormIsValid(c *gin.Context) error {
	formdata := c.Request.PostForm
	for k, v := range formdata {
		log.Debug(k, ":", v)
		if len(v) == 0 {
			log.Error("form format err", k)
			c.JSON(http.StatusOK, gin.H{
				"err_msg": "form format err",
			})
			return errors.New("fornat err")
		}
	}
	return nil
}

// Query teacher page.
func TeacherManagerHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	college_name, _ := api.GetALlCollegeName()
	var college_option string
	for _, v := range college_name {
		college_option += "<option>" + v + "</option>"
	}

	c.HTML(http.StatusOK, "admin_teacher_manager.html", gin.H{
		"loginer_name":   a.Info.GetUser(),
		"college_name":   college_name,
		"college_option": college_option,
	})
}

// Query teacher result page.
func AdminTeacherManagerHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	// 根据需要获取教师的信息列表
	c.Request.ParseForm()

	log.Info(c.Request.PostForm)

	college_name := c.Request.PostForm.Get("college_name")
	name := c.Request.PostForm.Get("name")
	NRIC := c.Request.PostForm.Get("NRIC")
	teacher_uid_str := c.Request.PostForm.Get("teacher_uid")

	m := make(map[uint64]DataCenter.TeacherInfo)
	var err error

	if teacher_uid_str != "" {
		// 直接通过teacher_uid查询
		teacher_uid, _ := strconv.ParseUint(teacher_uid_str, 10, 64)
		data, err := api.GetTeacherListByTeacherUid(teacher_uid)
		m[teacher_uid] = data
		if err != nil {
			// 出现错误
			c.HTML(http.StatusInternalServerError, "401.html", nil)
			return
		}

	} else if NRIC != "" {
		m, err = api.GetTeacherListByNRIC(NRIC)
		if err != nil {
			// 出现错误
			c.HTML(http.StatusInternalServerError, "401.html", nil)
			return
		}
	} else if name != "" {
		m, err = api.GetTeacherListByTeacherName(name)
		if err != nil {
			// 出现错误
			c.HTML(http.StatusInternalServerError, "401.html", nil)
			return
		}

	} else if college_name != "" {
		if college_name == "不限" {
			m, _ = api.GetAllTeacherList()
		} else {
			m, _ = api.GetTeacherListByCollegeName(college_name)
		}
	} else {
		// TODO:
	}

	// 渲染html
	if len(m) == 0 {
		c.JSON(http.StatusOK, nil)
		return
	}
	val, _ := json.Marshal(m)
	c.JSON(http.StatusOK, string(val))

}

func AdminStudentManagerHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	college_name, _ := api.GetALlCollegeName()

	c.HTML(http.StatusOK, "admin_student_manager.html", gin.H{
		"loginer_name": a.Info.GetUser(),
		"college_name": college_name,
	})
}

func AdminNoticeManagerGetHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "admin_notice_manager.html", gin.H{
		"loginer_name": a.Info.GetUser(),
	})
}

func AdminNoticeManagerPostHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()

	title := c.Request.PostForm.Get("title")
	notice := c.Request.PostForm.Get("notice_text")

	err := dao.DataBase.Execf("insert into `notice`(`title`, `data`) values ('%s', '%s')", title, notice)
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusInternalServerError, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "success.html", gin.H{
		"second": 3,
		"url":    "admin_index",
	})
}

func AdminDeleteTeacherHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	teacher_uid := c.Query("teacher_uid")

	err := dao.DataBase.Execf("delete from `teacher` where `teacher_uid`='%s'", teacher_uid)
	if err != nil {
		log.Error(err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, "admin_index")
}

func AdminEditTeacherHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	teacher_uid_str := c.Query("teacher_uid")
	teacher_uid, _ := strconv.ParseUint(teacher_uid_str, 10, 64)
	m, _ := api.GetTeacherListByTeacherUid(teacher_uid)
	teacher_info := m

	c.HTML(http.StatusOK, "admin_edit_teacher.html", gin.H{
		"loginer_name": a.Info.GetUser(),
		"name":         teacher_info.GetName(),
		"sex":          teacher_info.GetSex(),
		"NRIC":         teacher_info.GetNRIC(),
		"status":       teacher_info.GetStatus(),
		"teacher_uid":  teacher_info.GetTeacherUid(),
		"college_name": teacher_info.GetCollegeUid(),
		"create_time":  teacher_info.GetCreateTime(),
	})
}

func AdminUpdateTeacherPersonInfoHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	teacher_uid := c.Query("teacher_uid")

	c.Request.ParseForm()
	name := c.Request.PostForm.Get("name")
	sex := c.Request.PostForm.Get("sex")
	NRIC := c.Request.PostForm.Get("NRIC")
	status := c.Request.PostForm.Get("status")
	teacher_uid_update := c.Request.PostForm.Get("teacher_uid")
	college_name := c.Request.PostForm.Get("college_name")
	college_uid, err := api.GetCollegeUidByName(college_name)
	if err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	err = dao.DataBase.Execf("update `teacher` set `name`='%s', `sex`='%s', `NRIC`='%s', `status`='%s', `college_uid`='%s', `teacher_uid`='%s' where `teacher_uid`='%s'",
		name, sex, NRIC, status, college_uid, teacher_uid_update, teacher_uid)
	if err != nil {
		c.HTML(http.StatusBadGateway, "502.html", nil)
		return
	}

	c.HTML(http.StatusOK, "success.html", gin.H{
		"url":    "admin_index",
		"second": "3",
	})
}

func AdminStudentManagerPostHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "502.html", nil)
		return
	}

	c.Request.ParseForm()
	log.Info(c.Request.PostForm)
	student_name := c.Request.PostForm.Get("student_name")
	major_name := c.Request.PostForm.Get("major_name")
	class_name := c.Request.PostForm.Get("class_name")
	college_name := c.Request.PostForm.Get("college_name")
	NRIC := c.Request.PostForm.Get("NRIC")
	student_uid_str := c.Request.PostForm.Get("student_uid")

	m := make(map[uint64]DataCenter.StudentInfo)
	var err error

	if len(student_uid_str) != 0 {
		student_uid, _ := strconv.ParseUint(student_uid_str, 10, 64)
		val, err := api.GetStudentByStudentUid(student_uid)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "502.html", nil)
			return
		}
		m[student_uid] = val
	} else if len(NRIC) != 0 {
		m, err = api.GetStudentByNRIC(NRIC)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "502.html", nil)
			return
		}
	} else if len(student_name) != 0 {
		m, err = api.GetStudentByName(student_name)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "502.html", nil)
			return
		}
	} else if len(class_name) != 0 {
		class_uid, err := api.GetClassUidByName(class_name)
		if err != nil || class_uid == 0 {
			c.HTML(http.StatusInternalServerError, "502.html", nil)
			return
		}
		// 通过class_uid 获取所有学生
		m, err = api.GetStudentListByClassUid(class_uid)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "502.html", nil)
			return
		}
	} else if len(major_name) != 0 {
		major_uid, err := api.GetMajorUidByName(major_name)
		log.Info(major_uid)
		if err != nil || major_uid == 0 {
			c.HTML(http.StatusInternalServerError, "502.html", nil)
			return
		}
		m, err = api.GetStudentListByMajorUid(major_uid)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "502.html", nil)
			return
		}
	} else if len(college_name) != 0 {
		if college_name == "不限" {
			m, err = api.GetAllStudentList()
			if err != nil {
				c.HTML(http.StatusInternalServerError, "502.html", nil)
				return
			}
		} else {
			college_uid, err := api.GetCollegeUidByName(college_name)
			if err != nil {
				c.HTML(http.StatusInternalServerError, "502.html", nil)
				return
			}

			m, err = api.GetStudentListByCollegeUid(college_uid)
			if err != nil {
				c.HTML(http.StatusInternalServerError, "502.html", nil)
				return
			}
		}
	} else {
		m, err = api.GetAllStudentList()
		if err != nil {
			c.HTML(http.StatusInternalServerError, "502.html", nil)
			return
		}
	}
	// map to json
	json_val, _ := json.Marshal(m)
	c.JSON(http.StatusOK, string(json_val))
}

func AdminDeleteStudentHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "502.html", nil)
		return
	}

	student_uid_str := c.Query("student_uid")

	err := dao.DataBase.Execf("delete from `student` where `student_uid`='%s'", student_uid_str)
	if err != nil {
		log.Error(err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "admin_index")
}

func AdminEditStudentHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	student_uid_str := c.Query("student_uid")
	student_uid, _ := strconv.ParseUint(student_uid_str, 10, 64)
	m, _ := api.GetStudentByStudentUid(student_uid)
	student_info := m

	class_name, _ := api.GetNamebyUid(student_info.GetClassUid(), "class", "class_uid")
	major_name, _ := api.GetNamebyUid(student_info.GetMajorUid(), "major", "major_uid")
	college_name, _ := api.GetNamebyUid(student_info.GetCollegeUid(), "college", "college_uid")

	c.HTML(http.StatusOK, "admin_edit_student.html", gin.H{
		"loginer_name": a.Info.GetUser(),
		"name":         student_info.GetName(),
		"sex":          student_info.GetSex(),
		"NRIC":         student_info.GetNRIC(),
		"status":       student_info.GetStatus(),
		"student_uid":  student_info.GetStudentUid(),
		"major_name":   major_name,
		"class_name":   class_name,
		"college_name": college_name,
		"create_time":  student_info.GetCreateTime(),
	})
}

func AdminUpdateStudentPersonInfoHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	student_uid := c.Query("student_uid")

	c.Request.ParseForm()
	name := c.Request.PostForm.Get("name")
	sex := c.Request.PostForm.Get("sex")
	NRIC := c.Request.PostForm.Get("NRIC")
	status := c.Request.PostForm.Get("status")
	student_uid_update := c.Request.PostForm.Get("student_uid")
	college_name := c.Request.PostForm.Get("college_name")
	major_name := c.Request.PostForm.Get("major_name")
	class_name := c.Request.PostForm.Get("class_name")
	college_uid, err := api.GetCollegeUidByName(college_name)
	if err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}
	major_uid, err := api.GetMajorUidByName(major_name)
	if err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}
	class_uid, err := api.GetClassUidByName(class_name)
	if err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}
	err = dao.DataBase.Execf("update `student` set `name`='%s', `sex`='%s', `NRIC`='%s', `status`='%s', `college_uid`='%d', `student_uid`='%s',`major_uid`='%d',`class_uid`='%d'  where `student_uid`='%s'",
		name, sex, NRIC, status, college_uid, student_uid_update, major_uid, class_uid, student_uid)
	if err != nil {
		c.HTML(http.StatusBadGateway, "502.html", nil)
		return
	}

	c.HTML(http.StatusOK, "success.html", gin.H{
		"url":    "admin_index",
		"second": "3",
	})
}

func AdminCourseGetHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "admin_course_manager.html", gin.H{
		"loginer_name": a.Info.GetUser(),
	})
}
