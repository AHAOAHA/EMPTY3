/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: api.go
 * Author: ahaoozhang
 * Date: 2020-03-26 11:12:43 (Thursday)
 * Describe:
 ******************************************************************/
package service

import (
	"GradeManager/src/api"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCollegeNameHandler(c *gin.Context) {
	data, _ := api.GetALlCollegeName()
	var m []gin.H
	m = append(m, gin.H{
		"name": "不限",
	})
	for _, v := range data {
		m = append(m, gin.H{
			"name": v,
		})
	}

	val, _ := json.Marshal(m)
	c.JSON(http.StatusOK, string(val))
}

func GetAllMajorNameHandler(c *gin.Context) {
	data, _ := api.GetAllMajerName()
	var m []gin.H
	for _, v := range data {
		m = append(m, gin.H{
			"name": v,
		})
	}

	val, _ := json.Marshal(m)
	c.JSON(http.StatusOK, string(val))
}

func GetAllClassNameHandler(c *gin.Context) {
	data, _ := api.GetAllClassName()
	var m []gin.H
	for _, v := range data {
		m = append(m, gin.H{
			"name": v,
		})
	}

	val, _ := json.Marshal(m)
	c.JSON(http.StatusOK, string(val))
}

func GetAllClassHandler(c *gin.Context) {
	data, _ := api.GetAllClassInfo()
	var result []struct {
		ClassUid  uint64
		ClassName string
	}
	for _, v := range data {
		result = append(result, struct {
			ClassUid  uint64
			ClassName string
		}{
			v.GetClassUid(),
			v.GetName(),
		})
	}

	c.JSON(http.StatusOK, result)
}

func GetAllCourseInfoHandler(c *gin.Context) {
	data, _ := api.GetAllCourseInfo()
	var result []struct {
		CourseUid  uint64
		CourseName string
	}
	for _, v := range data {
		result = append(result, struct {
			CourseUid  uint64
			CourseName string
		}{
			v.GetCourseUid(),
			v.GetName(),
		})
	}
	c.JSON(http.StatusOK, result)
}

func DebugHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "debug.html", nil)
}
