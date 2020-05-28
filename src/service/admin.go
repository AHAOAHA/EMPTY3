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
		"loginer_name": a.Info.GetUser(),
	})
}

// 添加教师
func AdminAddTeacherPostHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()

	form_data := c.Request.PostForm
	err := dao.DataBase.Execf("insert into `teacher` (`teacher_uid`, `college_uid`, `password`, `name`, `sex`, `NRIC`) values('%s', '%s', '%s', '%s', '%s', '%s' )",
		form_data.Get("teacher_uid"), form_data.Get("college_uid"), form_data.Get("password"), form_data.Get("username"), form_data.Get("sex"), form_data.Get("NRIC"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_msg":  err.Error(),
			"err_code": 1,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "添加成功！",
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
	// 获取学院id 班级id 专业id
	formdata := c.Request.PostForm

	err := dao.DataBase.Execf("insert into `student`(`student_uid`, `class_uid`, `college_uid`, `major_uid`, `password`, `name`, `sex`, `NRIC`) values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')",
		formdata.Get("student_uid"), formdata.Get("class_uid"), formdata.Get("college_uid"), formdata.Get("major_uid"), formdata.Get("password"), formdata.Get("name"), formdata.Get("sex"), formdata.Get("NRIC"))
	if err != nil {
		// 插入失败
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"err_msg":  "insert student err",
			"err_code": 2,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "插入成功！",
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

	err := dao.DataBase.Execf("insert into `college`(`college_uid`, `name`) values('%s', '%s')", formdata.Get("college_uid"), formdata.Get("name"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"err_msg":  err.Error(),
			"err_code": 2,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "添加成功！",
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
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	formdata := c.Request.PostForm

	err := dao.DataBase.Execf("insert into `major`(`major_uid`, `college_uid`, `name`)values('%s', '%s', '%s')", formdata.Get("major_uid"), formdata.Get("college_uid"), formdata.Get("name"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_msg":  err.Error(),
			"err_code": 3,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "添加成功！",
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
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	formdata := c.Request.PostForm

	err := dao.DataBase.Execf("insert into `class`(`class_uid`, `college_uid`, `major_uid`, `name`) values ('%s', '%s', '%s', '%s')",
		formdata.Get("class_uid"), formdata.Get("college_uid"), formdata.Get("major_uid"), formdata.Get("name"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_msg":  err.Error(),
			"err_code": 1,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "添加成功！",
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
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	formdata := c.Request.PostForm

	err := dao.DataBase.Execf("insert into `course`(`course_uid`, `college_uid`, `name`, `credit`, `hour`, `type`) values ('%s', '%s', '%s', '%s', '%s', '%s')",
		formdata.Get("course_uid"), formdata.Get("college_uid"), formdata.Get("name"), formdata.Get("credit"), formdata.Get("hour"), formdata.Get("type"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err_msg":  err.Error(),
			"err_code": 3,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "添加成功！",
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

	c.HTML(http.StatusOK, "admin_teacher_manager.html", gin.H{
		"loginer_name": a.Info.GetUser(),
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

	collegeUIDStr := c.Request.PostForm.Get("college_uid")
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
			c.JSON(http.StatusOK, gin.H{
				"err_code": 1,
				"err_msg":  err.Error(),
			})
			return
		}

	} else if NRIC != "" {
		m, err = api.GetTeacherListByNRIC(NRIC)
		if err != nil {
			// 出现错误
			c.JSON(http.StatusOK, gin.H{
				"err_code": 2,
				"err_msg":  err.Error(),
			})
			return
		}
	} else if name != "" {
		m, err = api.GetTeacherListByTeacherName(name)
		if err != nil {
			// 出现错误
			c.JSON(http.StatusOK, gin.H{
				"err_code": 3,
				"err_msg":  err.Error(),
			})
			return
		}

	} else if collegeUIDStr != "" {
		if collegeUIDStr == "不限" {
			m, _ = api.GetAllTeacherList()
		} else {
			collegeUID, _ := strconv.ParseUint(collegeUIDStr, 10, 64)
			m, _ = api.GetTeacherListByCollegeUid(collegeUID)
		}
	} else {
		log.Error("teacher manager query teacher list fail!")
		return
	}

	// 渲染html
	if len(m) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 4,
			"err_msg":  "查询结果为空",
		})
		return
	}

	var result []struct {
		TeacherName string
		TeacherSex  string
		TeacherUid  uint64
		CollegeName string
	}

	for _, v := range m {
		college, _ := api.GetCollegeInfoByCollegeUid(v.GetCollegeUid())
		result = append(result, struct {
			TeacherName string
			TeacherSex  string
			TeacherUid  uint64
			CollegeName string
		}{
			v.GetName(),
			v.GetSex(),
			v.GetTeacherUid(),
			college.GetName(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"result":   result,
		"err_msg":  "查询成功！",
	})
}

func AdminStudentManagerHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	college_name, _ := api.GetALlCollegeInfo()

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
	college, _ := api.GetCollegeInfoByCollegeUid(teacher_info.GetCollegeUid())

	c.HTML(http.StatusOK, "admin_edit_teacher.html", gin.H{
		"loginer_name": a.Info.GetUser(),
		"name":         teacher_info.GetName(),
		"sex":          teacher_info.GetSex(),
		"NRIC":         teacher_info.GetNRIC(),
		"status":       teacher_info.GetStatus(),
		"teacher_uid":  teacher_info.GetTeacherUid(),
		"college_name": college.GetName(),
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
	majorUIDStr := c.Request.PostForm.Get("major_uid")
	classUIDStr := c.Request.PostForm.Get("class_uid")
	collegeUIDStr := c.Request.PostForm.Get("college_uid")
	NRIC := c.Request.PostForm.Get("NRIC")
	student_uid_str := c.Request.PostForm.Get("student_uid")

	m := make(map[uint64]DataCenter.StudentInfo)
	var err error

	if len(student_uid_str) != 0 {
		student_uid, _ := strconv.ParseUint(student_uid_str, 10, 64)
		val, err := api.GetStudentByStudentUid(student_uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err_code": 1,
				"err_msg":  err.Error(),
			})
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
	} else if len(classUIDStr) != 0 {
		classUID, _ := strconv.ParseUint(classUIDStr, 10, 64)
		// 通过class_uid 获取所有学生
		m, err = api.GetStudentListByClassUid(classUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err_code": 2,
				"err_msg":  err.Error(),
			})
			return
		}
	} else if len(majorUIDStr) != 0 {
		majorUID, _ := strconv.ParseUint(majorUIDStr, 10, 64)
		m, err = api.GetStudentListByMajorUid(majorUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err_code": 3,
				"err_msg":  err.Error(),
			})
			return
		}
	} else if len(collegeUIDStr) != 0 {
		if collegeUIDStr == "不限" {
			m, err = api.GetAllStudentList()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"err_code": 4,
					"err_msg":  err.Error(),
				})
				return
			}
		} else {
			collegeUID, _ := strconv.ParseUint(collegeUIDStr, 10, 64)
			m, err = api.GetStudentListByCollegeUid(collegeUID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"err_code": 5,
					"err_msg":  err.Error(),
				})
				return
			}
		}
	} else {
		log.Error("student manager query student list fail!")
		return
	}
	var result []struct {
		StudentName string
		StudentSex  string
		StudentUid  uint64
		ClassName   string
		MajorName   string
		CollegeName string
	}

	for _, v := range m {
		collegeName, _ := api.GetNamebyUid(v.GetCollegeUid(), "college", "college_uid")
		majorName, _ := api.GetNamebyUid(v.GetMajorUid(), "major", "major_uid")
		className, _ := api.GetNamebyUid(v.GetClassUid(), "class", "class_uid")
		result = append(result, struct {
			StudentName string
			StudentSex  string
			StudentUid  uint64
			ClassName   string
			MajorName   string
			CollegeName string
		}{
			v.GetName(),
			v.GetSex(),
			v.GetStudentUid(),
			className,
			majorName,
			collegeName,
		})
	}
	c.JSON(http.StatusOK, result)
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

func AdminQueryClassCourseHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()
	class_uid_str := c.Request.PostForm.Get("class_uid")
	class_uid, _ := strconv.ParseUint(class_uid_str, 10, 64)
	courses, _ := api.GetCourseByClassUid(class_uid)
	var result []struct {
		CourseName  string
		CollegeName string
		Credit      float32
		Hour        float32
		TeacherName string
		Type        string
		Status      string
	}

	for _, v := range courses {
		college_name, _ := api.GetNamebyUid(v.GetCollegeUid(), "college", "college_uid")
		teacherInfo, _ := api.GetTeacherInfoByClassUidAndCourseUid(class_uid, v.GetCourseUid())
		result = append(result, struct {
			CourseName  string
			CollegeName string
			Credit      float32
			Hour        float32
			TeacherName string
			Type        string
			Status      string
		}{
			v.GetName(),
			college_name,
			v.GetCredit(),
			v.GetHour(),
			teacherInfo.GetName(),
			v.GetType().String(),
			v.GetStatus().String(),
		})
	}

	//json_result, _ := json.Marshal(result)
	c.JSON(http.StatusOK, result)

}

func AdminAddCourseHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	classUIDStr := c.Query("class_uid")
	classUID, _ := strconv.ParseUint(classUIDStr, 10, 64)

	c.HTML(http.StatusOK, "admin_add_course.html", gin.H{
		"loginer_name": a.GetLoginerName(),
		"class_uid":    classUID,
	})
}

func AdminAddCourseSecondHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	classUIDStr := c.Query("class_uid")
	classUID, _ := strconv.ParseUint(classUIDStr, 10, 64)
	courseUIDStr := c.Query("course_uid")
	courseUID, _ := strconv.ParseUint(courseUIDStr, 10, 64)
	teacherUIDStr := c.Query("teacher_uid")
	teacherUID, _ := strconv.ParseUint(teacherUIDStr, 10, 64)

	// 查看班级是否包含改课程
	m, err := dao.DataBase.Queryf("select * from `student_course` where `class_uid`='%d' and `course_uid`='%d'", classUID, courseUID)
	if err != nil || len(m) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"err_code": 1001,
			"err_msg":  "课程已存在",
		})
		return
	}

	err = dao.DataBase.Execf("insert into `student_course`(`class_uid`, `course_uid`, `teacher_uid`) values ('%d', '%d', '%d')", classUID, courseUID, teacherUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"err_code": 1,
			"err_msg":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"err_code": 0,
		"err_msg":  "添加成功！",
	})

}

func AdminScoreManagerGetHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.HTML(http.StatusOK, "admin_student_score_manager_first.html", gin.H{
		"loginer_name": a.Info.GetUser(),
	})
}

func AdminGetStudentScoreByStudentUidHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	c.Request.ParseForm()

	studentUIDStr := c.Request.PostForm.Get("student_uid")
	studentUID, _ := strconv.ParseUint(studentUIDStr, 10, 64)
	data, _ := api.GetStudentSubmitScoreByStudentUid(studentUID)
	var result []struct {
		StudentName string
		StudentUid  uint64
		CourseName  string
		CourseUid   uint64
		UsualScore  float32
		MidScore    float32
		EndScore    float32
		Score       uint32
		Credit      float32
		ACredit     float32
	}

	for _, v := range data {
		studentName, _ := api.GetNamebyUid(v.GetStudentUid(), "student", "student_uid")
		courseName, _ := api.GetNamebyUid(v.GetCourseUid(), "course", "course_uid")
		result = append(result, struct {
			StudentName string
			StudentUid  uint64
			CourseName  string
			CourseUid   uint64
			UsualScore  float32
			MidScore    float32
			EndScore    float32
			Score       uint32
			Credit      float32
			ACredit     float32
		}{
			studentName,
			v.GetStudentUid(),
			courseName,
			v.GetCourseUid(),
			v.GetUsualScore(),
			v.GetMidtermScore(),
			v.GetEndtermScore(),
			v.GetScore(),
			v.GetCredit(),
			v.GetAcademicCredit(),
		})
	}

	rsp, _ := json.Marshal(result)
	c.HTML(http.StatusOK, "admin_student_score_manager.html", gin.H{
		"loginer_name":  a.Info.GetUser(),
		"student_score": string(rsp),
	})
}

func AdminChangeStudentScoreHandler(c *gin.Context) {
	var a context.AdminContext
	if err := a.CheckCookies(c, "user_cookie"); err != nil {
		c.HTML(http.StatusBadRequest, "401.html", nil)
		return
	}

	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)

	body_data := string(buf[0:n])

	var body map[string]interface{}
	_ = json.Unmarshal([]byte(body_data), &body)
	log.Info(body)
	score_data := body["Data"].([]interface{})
	for _, v := range score_data {
		log.Warn(v)
		student_data := v.(map[string]interface{})
		usual_score, _ := strconv.Atoi(student_data["UsualScore"].(string))
		mid_score, _ := strconv.Atoi(student_data["MidScore"].(string))
		end_score, _ := strconv.Atoi(student_data["EndScore"].(string))
		score, _ := strconv.Atoi(student_data["Score"].(string))

		err := dao.DataBase.Execf("update `score` set `usual_score`='%d', `midterm_score`='%d', `endterm_score`='%d', `score`='%d' where `student_uid`='%s' and `course_uid`='%s'", usual_score, mid_score, end_score, score, student_data["StudentUid"].(string), student_data["CourseUid"].(string))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err_code": 1,
				"err_msg":  err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"err_code": 0,
			"err_msg":  "修改成功！",
		})
	}
}
