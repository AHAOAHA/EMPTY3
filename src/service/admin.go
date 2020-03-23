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
	"GradeManager/src/context"
	"GradeManager/src/dao"
	"errors"
	"net/http"
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
	c.HTML(http.StatusOK, "admin_index.html", gin.H{
		"title":         "login",
		"student_count": count_student,
		"teacher_count": count_teacher,
		"college_count": count_college,
		"major_count":   count_major,
		"loginer_name":  a.Info.GetUser(),
		"introduce":     "西安科技大学简介",
		"school_title":  "西安科技大学",
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

	c.JSON(http.StatusOK, gin.H{
		"err_code": "ok",
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

	c.JSON(http.StatusOK, gin.H{
		"err_msg": "ok",
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

	c.JSON(http.StatusOK, gin.H{
		"err_msg": "ok",
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
	c.JSON(http.StatusOK, gin.H{
		"err_msg": "ok",
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

	c.JSON(http.StatusOK, gin.H{
		"err_msg": "ok",
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
	c.JSON(http.StatusOK, gin.H{
		"err_msg": "ok",
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

func TeacherManagerHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "admin_teacher_manager.html", gin.H{
		"loginer_name": a.Info.GetUser(),
	})
}

func AdminTeacherManagerHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	// 根据需要获取教师的信息列表
}
