/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: api.go
 * Author: ahaoozhang
 * Date: 2020-03-15 19:27:57 (Sunday)
 * Describe:
 ******************************************************************/
package api

import (
	"GradeManager/src/dao"
	DataCenter "GradeManager/src/proto"
	"errors"
	"strconv"
)

var TeacherCache *map[uint64]DataCenter.TeacherInfo
var StudentCache *map[uint64]DataCenter.StudentInfo

func init() {
	if TeacherCache == nil {
		TeacherCache = new(map[uint64]DataCenter.TeacherInfo)
	}
	if StudentCache == nil {
		StudentCache = new(map[uint64]DataCenter.StudentInfo)
	}
}

// interface: proto Datacenter.TeacherInfo, every call update TeacherCache.
func GetAllTeacherList() (map[uint64]DataCenter.TeacherInfo, error) {
	sm, err := dao.DataBase.Queryf("select * from `teacher`")
	if err != nil {
		return nil, err
	}
	m := make(map[uint64]DataCenter.TeacherInfo)

	for _, v := range sm {
		// teacher_uid college_uid password name sex NRIC status create_time
		var teacher_uid, college_uid uint64
		var sta int
		sta, _ = strconv.Atoi(string(v["status"].([]uint8)))
		crtt, _ := strconv.Atoi(string(v["create_time"].([]uint8)))
		teacher_uid, _ = strconv.ParseUint(string(v["teacher_uid"].([]uint8)), 10, 64)
		college_uid, _ = strconv.ParseUint(string(v["college_uid"].([]uint8)), 10, 64)
		m[teacher_uid] = DataCenter.TeacherInfo{
			TeacherUid: teacher_uid,
			CollegeUid: college_uid,
			Name:       string(v["name"].([]uint8)),
			Password:   string(v["password"].([]uint8)),
			Sex:        string(v["sex"].([]uint8)),
			NRIC:       string(v["NRIC"].([]uint8)),
			Status:     DataCenter.TeacherInfo_STATUS(sta),
			CreateTime: uint32(crtt),
		}
	}

	// update cache
	TeacherCache = &m
	return m, nil
}

// interface: proto Datacenter.StudentInfo, every call update StudentCache.
func GetAllStudentList() (map[uint64]DataCenter.StudentInfo, error) {
	sm, err := dao.DataBase.Queryf("select * from `student`")
	if err != nil {
		return nil, err
	}
	m := make(map[uint64]DataCenter.StudentInfo)

	for _, v := range sm {
		// student_uid class_uid college_uid major_uid password name sex NRIC status create_time
		var student_uid, class_uid, college_uid, major_uid, crtt uint64
		var sta int
		crtt, _ = strconv.ParseUint(string(v["create_time"].([]uint8)), 10, 64)
		student_uid, _ = strconv.ParseUint(string(v["student_uid"].([]uint8)), 10, 64)
		college_uid, _ = strconv.ParseUint(string(v["college_uid"].([]uint8)), 10, 64)
		major_uid, _ = strconv.ParseUint(string(v["major_uid"].([]uint8)), 10, 64)
		class_uid, _ = strconv.ParseUint(string(v["class_uid"].([]uint8)), 10, 64)
		m[student_uid] = DataCenter.StudentInfo{
			StudentUid: student_uid,
			CollegeUid: college_uid,
			MajorUid:   major_uid,
			ClassUid:   class_uid,
			Name:       string(v["name"].([]uint8)),
			Password:   string(v["password"].([]uint8)),
			Sex:        string(v["sex"].([]uint8)),
			NRIC:       string(v["NRIC"].([]uint8)),
			Status:     DataCenter.StudentInfo_STATUS(sta),
			CreateTime: int32(crtt),
		}
	}

	// update cache
	StudentCache = &m
	return m, nil
}

// Query student info by teacher_uid, name, college_uid, college_name.
func GetTeacherListByTeacherUid(teacher_uid uint64) (map[uint64]DataCenter.TeacherInfo, error) {
	m := make(map[uint64]DataCenter.TeacherInfo)
	if TeacherCache != nil {
		teacher_info, ok := (*TeacherCache)[teacher_uid]
		if ok {
			m[teacher_uid] = teacher_info
			return m, nil
		}
	}

	// TeacherCache nil or teacher_uid not exist, query from database.
	db_m, err := dao.DataBase.Queryf("select * from `teacher` where `teacher_uid`='%d'", teacher_uid)
	if err != nil || len(db_m) != 1 {
		return nil, errors.New("query teacher Info err")
	}
	var college_uid uint64
	var sta int
	sta, _ = strconv.Atoi(string(db_m[0]["status"].([]uint8)))
	crtt, _ := strconv.Atoi(string(db_m[0]["create_time"].([]uint8)))
	college_uid, _ = strconv.ParseUint(string(db_m[0]["college_uid"].([]uint8)), 10, 64)
	m[teacher_uid] = DataCenter.TeacherInfo{
		TeacherUid: teacher_uid,
		CollegeUid: college_uid,
		Name:       string(db_m[0]["name"].([]uint8)),
		Password:   string(db_m[0]["password"].([]uint8)),
		Sex:        string(db_m[0]["sex"].([]uint8)),
		NRIC:       string(db_m[0]["NRIC"].([]uint8)),
		Status:     DataCenter.TeacherInfo_STATUS(sta),
		CreateTime: uint32(crtt),
	}

	// update TeacherCache
	if TeacherCache != nil {
		(*TeacherCache)[teacher_uid] = m[teacher_uid]
	}
	return m, nil
}

func GetTeacherListByNRIC(NRIC string) (map[uint64]DataCenter.TeacherInfo, error) {
	m := make(map[uint64]DataCenter.TeacherInfo)
	db_m, err := dao.DataBase.Queryf("select * from `teacher` where `NRIC`='%s'", NRIC)
	if err != nil || len(db_m) != 1 {
		return nil, errors.New("query teacher Info err")
	}
	var college_uid, teacher_uid uint64
	var sta int
	sta, _ = strconv.Atoi(string(db_m[0]["status"].([]uint8)))
	crtt, _ := strconv.Atoi(string(db_m[0]["create_time"].([]uint8)))
	college_uid, _ = strconv.ParseUint(string(db_m[0]["college_uid"].([]uint8)), 10, 64)
	teacher_uid, _ = strconv.ParseUint(string(db_m[0]["teacher_uid"].([]uint8)), 10, 64)
	m[teacher_uid] = DataCenter.TeacherInfo{
		TeacherUid: teacher_uid,
		CollegeUid: college_uid,
		Name:       string(db_m[0]["name"].([]uint8)),
		Password:   string(db_m[0]["password"].([]uint8)),
		Sex:        string(db_m[0]["sex"].([]uint8)),
		NRIC:       string(db_m[0]["NRIC"].([]uint8)),
		Status:     DataCenter.TeacherInfo_STATUS(sta),
		CreateTime: uint32(crtt),
	}
	return m, nil
}

// Without Cache.
func GetTeacherListByTeacherName(teacher_name string) (map[uint64]DataCenter.TeacherInfo, error) {
	dbm, err := dao.DataBase.Queryf("select * from `teacher` where `name`='%s'", teacher_name)
	if err != nil || len(dbm) == 0 {
		return nil, errors.New("query teacher info by name err")
	}

	m := make(map[uint64]DataCenter.TeacherInfo)

	for _, v := range dbm {
		var teacher_uid, college_uid uint64
		var sta int
		sta, _ = strconv.Atoi(string(v["status"].([]uint8)))
		crtt, _ := strconv.Atoi(string(v["create_time"].([]uint8)))
		college_uid, _ = strconv.ParseUint(string(v["college_uid"].([]uint8)), 10, 64)
		teacher_uid, err = strconv.ParseUint(string(v["teacher_uid"].([]uint8)), 10, 64)
		m[teacher_uid] = DataCenter.TeacherInfo{
			TeacherUid: teacher_uid,
			CollegeUid: college_uid,
			Name:       string(v["name"].([]uint8)),
			Password:   string(v["password"].([]uint8)),
			Sex:        string(v["sex"].([]uint8)),
			NRIC:       string(v["NRIC"].([]uint8)),
			Status:     DataCenter.TeacherInfo_STATUS(sta),
			CreateTime: uint32(crtt),
		}
	}

	return m, nil
}

// Without cache.
func GetTeacherListByCollegeUid(college_uid uint64) (map[uint64]DataCenter.TeacherInfo, error) {
	dbm, err := dao.DataBase.Queryf("select * from `teacher` where `college_uid`='%d'", college_uid)
	if err != nil || len(dbm) == 0 {
		return nil, errors.New("query teacher info by college_uid err")
	}

	m := make(map[uint64]DataCenter.TeacherInfo)

	for _, v := range dbm {
		var teacher_uid, college_uid uint64
		var sta int
		sta, _ = strconv.Atoi(string(v["status"].([]uint8)))
		crtt, _ := strconv.Atoi(string(v["create_time"].([]uint8)))
		college_uid, _ = strconv.ParseUint(string(v["college_uid"].([]uint8)), 10, 64)
		teacher_uid, err = strconv.ParseUint(string(v["teacher_uid"].([]uint8)), 10, 64)
		m[teacher_uid] = DataCenter.TeacherInfo{
			TeacherUid: teacher_uid,
			CollegeUid: college_uid,
			Name:       string(v["name"].([]uint8)),
			Password:   string(v["password"].([]uint8)),
			Sex:        string(v["sex"].([]uint8)),
			NRIC:       string(v["NRIC"].([]uint8)),
			Status:     DataCenter.TeacherInfo_STATUS(sta),
			CreateTime: uint32(crtt),
		}
	}

	return m, nil
}

// Without cache.
func GetTeacherListByCollegeName(college_name string) (map[uint64]DataCenter.TeacherInfo, error) {
	cdbm, err := dao.DataBase.Queryf("select * from `college` where `name`='%s'", college_name)
	if err != nil || len(cdbm) != 1 {
		return nil, errors.New("query teacher info by college_uid err")
	}
	college_uid, _ := strconv.ParseUint(string(cdbm[0]["college_uid"].([]uint8)), 10, 64)

	return GetTeacherListByCollegeUid(college_uid)
}

func GetALlCollegeName() ([]string, error) {
	cdbm, err := dao.DataBase.Queryf("select `name` from `college`")
	if err != nil || len(cdbm) == 0 {
		return nil, errors.New("query teacher info by college_uid err")
	}

	var college_name []string
	for _, val := range cdbm {
		college_name = append(college_name, string(val["name"].([]uint8)))
	}

	return college_name, nil
}

func GetNotice() (DataCenter.NoticeInfo, error) {
	var n DataCenter.NoticeInfo
	cdbm, err := dao.DataBase.Queryf("SELECT * from `notice` where `notice_uid` = (SELECT max(`notice_uid`) FROM notice)")
	if err != nil || len(cdbm) != 1 {
		return n, errors.New("query notice err")
	}

	n = DataCenter.NoticeInfo{
		Title:  string(cdbm[0]["title"].([]uint8)),
		Notice: string(cdbm[0]["data"].([]uint8)),
	}
	return n, nil
}
