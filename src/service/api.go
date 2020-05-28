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
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCollegeNameHandler(c *gin.Context) {
	data, _ := api.GetALlCollegeInfo()
	var result []struct {
		CollegeName string
		CollegeUid  uint64
	}
	for _, v := range data {
		result = append(result, struct {
			CollegeName string
			CollegeUid  uint64
		}{
			v.GetName(),
			v.GetCollegeUid(),
		})
	}
	c.JSON(http.StatusOK, result)
}

func GetAllMajorNameHandler(c *gin.Context) {
	var filterSwitch bool = true
	filterCollegeUIDStr := c.Query("college_uid")
	filterCollegeUID, err := strconv.ParseUint(filterCollegeUIDStr, 10, 64)
	if err != nil {
		filterSwitch = false
	}

	data, _ := api.GetAllMajerInfo()
	var result []struct {
		MajorName  string
		MajorUid   uint64
		CollegeUid uint64
	}
	for _, v := range data {
		if filterSwitch {
			if v.GetCollegeUid() != filterCollegeUID {
				continue
			}
		}
		result = append(result, struct {
			MajorName  string
			MajorUid   uint64
			CollegeUid uint64
		}{
			v.GetName(),
			v.GetMajorUid(),
			v.GetCollegeUid(),
		})
	}

	c.JSON(http.StatusOK, result)
}

func GetAllClassNameHandler(c *gin.Context) {
	var filterSwitch bool = true
	filterMajorUIDStr := c.Query("major_uid")
	filterMajorUID, err := strconv.ParseUint(filterMajorUIDStr, 10, 64)
	if err != nil {
		filterSwitch = false
	}

	data, _ := api.GetAllClassInfo()
	var result []struct {
		ClassName  string
		ClassUid   uint64
		MajorUid   uint64
		CollegeUid uint64
	}
	for _, v := range data {
		if filterSwitch {
			if v.GetMajorUid() != filterMajorUID {
				continue
			}
		}
		result = append(result, struct {
			ClassName  string
			ClassUid   uint64
			MajorUid   uint64
			CollegeUid uint64
		}{
			v.GetName(),
			v.GetClassUid(),
			v.GetMajorUid(),
			v.GetCollegeUid(),
		})
	}

	c.JSON(http.StatusOK, result)
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

func GetAllTeacherInfoHandler(c *gin.Context) {
	data, _ := api.GetAllTeacherList()
	var result []struct {
		TeacherUid  uint64
		TeacherName string
	}

	for _, v := range data {
		result = append(result, struct {
			TeacherUid  uint64
			TeacherName string
		}{
			v.GetTeacherUid(),
			v.GetName(),
		})
	}

	c.JSON(http.StatusOK, result)
}

func GetTeacherInfoHandler(c *gin.Context) {
	collegeUIDStr := c.Query("college_uid")
	collegeUID, _ := strconv.ParseUint(collegeUIDStr, 10, 64)

	data, _ := api.GetTeacherInfoByCollegeUid(collegeUID)
	var result []struct {
		TeacherUid  uint64
		TeacherName string
		CollegeUid  uint64
	}

	for _, v := range data {
		result = append(result, struct {
			TeacherUid  uint64
			TeacherName string
			CollegeUid  uint64
		}{
			v.GetTeacherUid(),
			v.GetName(),
			v.GetCollegeUid(),
		})
	}

	c.JSON(http.StatusOK, result)
}

func GetCourseInfoByCollegeUidHandler(c *gin.Context) {
	collegeUIDStr := c.Query("college_uid")
	collegeUID, _ := strconv.ParseUint(collegeUIDStr, 10, 64)

	data, _ := api.GetCourseByCollegeUid(collegeUID)
	var result []struct {
		CourseUid  uint64
		CourseName string
		CollegeUid uint64
	}
	for _, v := range data {
		result = append(result, struct {
			CourseUid  uint64
			CourseName string
			CollegeUid uint64
		}{
			v.GetCourseUid(),
			v.GetName(),
			v.GetCollegeUid(),
		})
	}

	c.JSON(http.StatusOK, result)
}
