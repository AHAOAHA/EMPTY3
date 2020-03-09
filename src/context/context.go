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
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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
	m, err := dao.DataBase.Queryf("select `password` from `admin` where `user`='%s'", username)
	if err != nil || len(m) != 1 {
		return errors.New("admin username name not exist")
	}
	if string(m[0]["password"].([]uint8)) != password {
		return errors.New("admin password err")
	}
	// success
	log.Info(username, " admin login success!")

	return nil
}

func (s *StudentContext) Login(username string, password string) error {
	// 学生登录
	m, err := dao.DataBase.Queryf("select `password` from `student` where `student_uid`='%s'", username)
	if err != nil || len(m) != 1 {
		return errors.New("student username name not exist")
	}
	if string(m[0]["password"].([]uint8)) != password {
		return errors.New("student password err")
	}
	// success
	log.Info(username, " student login success!")
	return nil
}

func (t *TeacherContext) Login(username string, password string) error {
	m, err := dao.DataBase.Queryf("select `password` from `teacher` where `teacher_uid`='%s'", username)
	if err != nil || len(m) != 1 {
		return errors.New("teacher username name not exist")
	}
	if string(m[0]["password"].([]uint8)) != password {
		return errors.New("teacher password err")
	}
	// success
	log.Info(username, " teacher login success!")
	return nil
}

// 重定向到主页
func (a *AdminContext) RedirectIndex(c *gin.Context) error {
	c.Redirect(http.StatusMovedPermanently, "/admin_index")
	return nil
}

// 重定向到主页
func (a *StudentContext) RedirectIndex(c *gin.Context) error {
	c.Redirect(http.StatusMovedPermanently, "/student_index")
	return nil
}

// 重定向到主页
func (a *TeacherContext) RedirectIndex(c *gin.Context) error {
	c.Redirect(http.StatusMovedPermanently, "/student_index")
	return nil
}
