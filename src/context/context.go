/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: context.go
 * Author: ahaoozhang
 * Date: 2020-03-04 22:09:28 (Wednesday)
 * Describe:
 ******************************************************************/
package context

import (
	"GradeManager/src/dao"
	DataCenter "GradeManager/src/proto"
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

type AdminContext struct {
	Info DataCenter.AdminInfo
	// TODO:
}

type StudentContext struct {
	Info DataCenter.StudentInfo
	// TODO:
}

type TeacherContext struct {
	Info DataCenter.TeacherInfo
	// TODO:
}

func (a *AdminContext) IsValid() error {
	// TODO:
	return nil
}

func (s *StudentContext) IsValid() error {
	// TODO:
	return nil
}

func (t *TeacherContext) IsValid() error {
	// TODO:
	return nil
}

func (a *AdminContext) Login(username string, password string) error {
	// 管理员登录
	m, err := dao.DataBase.Queryf("select * from `admin` where `user`='%s'", username)
	if err != nil || len(m) != 1 {
		return errors.New("admin username name not exist")
	}
	if string(m[0]["password"].([]uint8)) != password {
		return errors.New("admin password err")
	}
	crtt, _ := strconv.Atoi(string(m[0]["create_time"].([]uint8)))
	a.Info = DataCenter.AdminInfo{
		User:       username,
		Password:   password,
		Mail:       string(m[0]["mail"].([]uint8)),
		CreateTime: int32(crtt),
		//ExprTime: m[0]["expr_time"].(int32),
	}
	// success
	log.Info(username, " admin login success!")

	return nil
}

func (s *StudentContext) Login(username string, password string) error {
	// 学生登录
	m, err := dao.DataBase.Queryf("select * from `student` where `student_uid`='%s'", username)
	if err != nil || len(m) != 1 {
		return errors.New("student username name not exist")
	}
	if string(m[0]["password"].([]uint8)) != password {
		return errors.New("student password err")
	}
	var crtt, stuid, colid, majid, claid uint64
	crtt, _ = strconv.ParseUint(string(m[0]["create_time"].([]uint8)), 10, 64)
	stuid, _ = strconv.ParseUint(string(m[0]["student_uid"].([]uint8)), 10, 64)
	colid, _ = strconv.ParseUint(string(m[0]["college_uid"].([]uint8)), 10, 64)
	majid, _ = strconv.ParseUint(string(m[0]["major_uid"].([]uint8)), 10, 64)
	claid, _ = strconv.ParseUint(string(m[0]["class_uid"].([]uint8)), 10, 64)
	s.Info = DataCenter.StudentInfo{
		StudentUid: stuid,
		Password:   password,
		CollegeUid: colid,
		MajorUid:   majid,
		ClassUid:   claid,
		Name:       string(m[0]["name"].([]uint8)),
		Sex:        string(m[0]["sex"].([]uint8)),
		NRIC:       string(m[0]["NRIC"].([]uint8)),
		CreateTime: int32(crtt),
	}
	// success
	log.Info(username, " student login success!")
	return nil
}

func (t *TeacherContext) Login(username string, password string) error {
	m, err := dao.DataBase.Queryf("select * from `teacher` where `teacher_uid`='%s'", username)
	if err != nil || len(m) != 1 {
		return errors.New("teacher username name not exist")
	}
	if string(m[0]["password"].([]uint8)) != password {
		return errors.New("teacher password err")
	}
	// 设置登录信息
	var teaid, colid uint64
	var sta int
	sta, _ = strconv.Atoi(string(m[0]["status"].([]uint8)))
	teaid, _ = strconv.ParseUint(string(m[0]["teacher_uid"].([]uint8)), 10, 64)
	colid, _ = strconv.ParseUint(string(m[0]["college_uid"].([]uint8)), 10, 64)
	crtt, _ := strconv.Atoi(string(m[0]["create_time"].([]uint8)))
	t.Info = DataCenter.TeacherInfo{
		TeacherUid: teaid,
		Password:   password,
		CollegeUid: colid,
		Status:     DataCenter.TeacherInfo_STATUS(sta),
		Sex:        string(m[0]["sex"].([]uint8)),
		NRIC:       string(m[0]["NRIC"].([]uint8)),
		CreateTime: uint32(crtt),
	}
	// success
	log.Info(username, " teacher login success!")
	return nil
}

// 重定向到主页
func (a *AdminContext) RedirectIndex(c *gin.Context) error {
	// c.HTML(http.StatusOK, "admin_index.html", gin.H{
	// 	"err_code": "success",
	// })
	c.Redirect(http.StatusMovedPermanently, "/admin_index")
	return nil
}

// 重定向到主页
func (s *StudentContext) RedirectIndex(c *gin.Context) error {
	// c.HTML(http.StatusOK, "student_index.html", gin.H{
	// 	"err_code": "success",
	// })
	c.Redirect(http.StatusMovedPermanently, "student_index")
	return nil
}

// 重定向到主页
func (t *TeacherContext) RedirectIndex(c *gin.Context) error {
	// c.HTML(http.StatusOK, "teacher_index.html", gin.H{
	// 	"err_code": "success",
	// })
	c.Redirect(http.StatusMovedPermanently, "teacher_index")
	return nil
}

func (a *AdminContext) SetCookies(c *gin.Context) error {
	data, err := proto.Marshal(&a.Info)
	if err != nil {
		log.Error(err)
		return err
	}
	val := base64.StdEncoding.EncodeToString(data)
	c.SetCookie("userinfo", val, 1000, "/", "localhost", false, true)
	return nil
}

func (s *StudentContext) SetCookies(c *gin.Context) error {
	data, err := proto.Marshal(&s.Info)
	if err != nil {
		log.Error(err)
		return err
	}
	val := base64.StdEncoding.EncodeToString(data)
	c.SetCookie("username", val, 1000, "/", "localhost", false, true)
	return nil
}

func (t *TeacherContext) SetCookies(c *gin.Context) error {
	data, err := proto.Marshal(&t.Info)
	if err != nil {
		log.Error(err)
		return err
	}
	val := base64.StdEncoding.EncodeToString(data)
	c.SetCookie("userinfo", val, 1000, "/", "localhost", false, true)
	return nil
}
