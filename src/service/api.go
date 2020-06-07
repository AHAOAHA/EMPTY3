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
	"GradeManager/src/config"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

func IpToAddrHandler(c *gin.Context) {
	ip := c.Query("ip")
	param := config.Config.IPToAddr.Path + "?ip=" + ip + "&key=" + config.Config.IPToAddr.Key
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(param + config.Config.IPToAddr.SK))
	sig := hex.EncodeToString(md5Ctx.Sum(nil))
	param = param + "&sig=" + sig
	log.Info(param)
	GenUrl := config.Config.IPToAddr.Url + param
	u, _ := url.Parse(GenUrl)
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode
	rsp, err := http.Get(u.String())
	if err != nil {
		log.Error(err)
		return
	}
	body, _ := ioutil.ReadAll(rsp.Body)

	var result struct {
		Status  int32  `json:"status"`
		Message string `json:"message"`
		Result  struct {
			IP       string `json:"ip"`
			Location struct {
				Lng float32 `json:"lng"`
				Lat float32 `json:"lat"`
			} `json:"location"`
			AdInfo struct {
				Nation   string `json:"nation"`
				Province string `json:"province"`
				City     string `json:"city"`
				AdCode   int32  `json:"adcode"`
			} `json:"ad_info"`
		} `json:"result"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Error(err)
	}

	c.JSON(http.StatusOK, result)
}

func LocationToAddrHandler(c *gin.Context) {
	location := c.Query("location")
	param := config.Config.LocationToAddr.Path + "?key=" + config.Config.LocationToAddr.Key + "&location=" + location
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(param + config.Config.LocationToAddr.SK))
	sig := hex.EncodeToString(md5Ctx.Sum(nil))
	param = param + "&sig=" + sig
	log.Info(param)
	GenUrl := config.Config.LocationToAddr.Url + param
	u, _ := url.Parse(GenUrl)
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode
	rsp, err := http.Get(u.String())
	if err != nil {
		log.Error(err)
		return
	}

	body, _ := ioutil.ReadAll(rsp.Body)
	var result struct {
		Status    int32  `json:"status"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
		Result    struct {
			Location struct {
				Lat float32 `json:"lat"`
				Lng float32 `json:"lng"`
			} `json:"location"`
			AddressComponent struct {
				Nation   string `json:"nation"`
				AdLevel1 string `json:"ad_level_1"`
				AdLevel2 string `json:"ad_level_2"`
				AdLevel3 string `json:"ad_level_3"`
				Street   string `json:"street"`
				Locality string `json:"locality"`
			} `json:"address_component"`
			AdInfo struct {
				NationCode string `json:"nation_Code"`
				CityCode   string `json:"city_Code"`
				Location   struct {
					Lat float32 `json:"lat"`
					Lng float32 `json:"lng"`
				} `json:"location"`
			} `json:"ad_info"`
			Address string `json:"address"`
		} `json:"result"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Error(err)
	}

	c.JSON(http.StatusOK, result)
}
