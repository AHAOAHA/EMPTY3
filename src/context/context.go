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
	if a.Info.GetLType() != DataCenter.LoginType_ADMIN {
		return errors.New("Login type err")
	}
	return nil
}

func (s *StudentContext) IsValid() error {
	if s.Info.GetLType() != DataCenter.LoginType_STUDENT {
		return errors.New("Login type err")
	}
	return nil
}

func (t *TeacherContext) IsValid() error {
	if t.Info.GetLType() != DataCenter.LoginType_TEACHER {
		return errors.New("Login type err")
	}
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
	exprt, _ := strconv.Atoi(string(m[0]["expr_time"].([]uint8)))
	a.Info = DataCenter.AdminInfo{
		User:       username,
		Password:   password,
		Mail:       string(m[0]["mail"].([]uint8)),
		CreateTime: int32(crtt),
		ExprTime:   int32(exprt),
		LType:      DataCenter.LoginType_ADMIN,
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
		LType:      DataCenter.LoginType_STUDENT,
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
		Name:       string(m[0]["name"].([]uint8)),
		LType:      DataCenter.LoginType_TEACHER,
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
	c.Redirect(http.StatusMovedPermanently, "/teacher_index")
	return nil
}

func (a *AdminContext) Entcry() string {
	data, err := proto.Marshal(&a.Info)
	if err != nil {
		log.Error(err)
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(data)
}

func (a *AdminContext) Detcry(cookie string) error {
	data, err := base64.RawURLEncoding.DecodeString(cookie)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(data, &a.Info)
	return err
}

func (s *StudentContext) Entcry() string {
	data, err := proto.Marshal(&s.Info)
	if err != nil {
		log.Error(err)
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(data)
}

func (s *StudentContext) Detcry(cookie string) error {
	data, err := base64.RawURLEncoding.DecodeString(cookie)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(data, &s.Info)
	return err
}

func (t *TeacherContext) Entcry() string {
	data, err := proto.Marshal(&t.Info)
	if err != nil {
		log.Error(err)
		return ""
	}
	return base64.RawStdEncoding.EncodeToString(data)
}

func (t *TeacherContext) Detcry(cookie string) error {
	data, err := base64.RawURLEncoding.DecodeString(cookie)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(data, &t.Info)
	return err
}

func (a *AdminContext) SetCookies(c *gin.Context) {
	c.SetCookie("user_cookie", a.Entcry(), 1000, "/", "", false, true)
}

func (t *TeacherContext) SetCookies(c *gin.Context) {
	c.SetCookie("user_cookie", t.Entcry(), 1000, "/", "", false, true)
}

func (s *StudentContext) SetCookies(c *gin.Context) {
	c.SetCookie("user_cookie", s.Entcry(), 1000, "/", "", false, true)
}

func (a *AdminContext) CheckCookies(c *gin.Context, key string) error {
	cookie, err := c.Request.Cookie(key)
	if err != nil {
		return err
	}
	err = a.Detcry(cookie.Value)
	if err != nil {
		return err
	}

	if err = a.IsValid(); err != nil {
		return err
	}

	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	return nil
}

func (t *TeacherContext) CheckCookies(c *gin.Context, key string) error {
	cookie, err := c.Request.Cookie(key)
	if err != nil {
		return err
	}
	err = t.Detcry(cookie.Value)
	if err != nil {
		return err
	}

	log.Warnf("%d", t.Info.GetLType())

	if err = t.IsValid(); err != nil {
		return err
	}

	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	return nil
}

func (s *StudentContext) CheckCookies(c *gin.Context, key string) error {
	cookie, err := c.Request.Cookie(key)
	if err != nil {
		return err
	}
	err = s.Detcry(cookie.Value)
	if err != nil {
		return err
	}
	if err = s.IsValid(); err != nil {
		return err
	}
	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	return nil
}

func (a *AdminContext) GetPassword() string {
	if a == nil {
		return ""
	}
	return a.Info.GetPassword()
}

func (t *TeacherContext) GetPassword() string {
	if t == nil {
		return ""
	}
	return t.Info.GetPassword()
}

func (s *StudentContext) GetPassword() string {
	if s == nil {
		return ""
	}
	return s.Info.GetPassword()
}

func (a *AdminContext) UpdatePassword(new_password string) error {
	if err := a.IsValid(); err != nil {
		return err
	}

	return dao.DataBase.Execf("update `admin` set `password`='%s' where `user`='%s'", new_password, a.Info.GetUser())
}

func (t *TeacherContext) UpdatePassword(new_password string) error {
	if err := t.IsValid(); err != nil {
		return err
	}
	return dao.DataBase.Execf("update `teacher` set `password`='%s' where `teacher_uid`='%s'", new_password, t.Info.GetTeacherUid())
}

func (s *StudentContext) UpdatePassword(new_password string) error {
	if err := s.IsValid(); err != nil {
		return err
	}
	return dao.DataBase.Execf("update `student` set `password`='%s' where `student_uid`='%s'", new_password, s.Info.GetStudentUid())
}
