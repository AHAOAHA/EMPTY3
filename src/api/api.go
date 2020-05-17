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

	log "github.com/sirupsen/logrus"
)

var TeacherCache *map[uint64]DataCenter.TeacherInfo
var StudentCache *map[uint64]DataCenter.StudentInfo
var CollegeCache *map[uint64]DataCenter.CollegeInfo
var MajorCache *map[uint64]DataCenter.MajorInfo

func init() {
	if TeacherCache == nil {
		TeacherCache = new(map[uint64]DataCenter.TeacherInfo)
	}
	if StudentCache == nil {
		StudentCache = new(map[uint64]DataCenter.StudentInfo)
	}
	if CollegeCache == nil {
		CollegeCache = new(map[uint64]DataCenter.CollegeInfo)
	}
	if MajorCache == nil {
		MajorCache = new(map[uint64]DataCenter.MajorInfo)
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
			CreateTime: string(v["create_time"].([]uint8)),
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
		var student_uid, class_uid, college_uid, major_uid uint64
		var sta int
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
			CreateTime: string(v["create_time"].([]uint8)),
		}
	}

	// update cache
	StudentCache = &m
	return m, nil
}

// Query student info by teacher_uid, name, college_uid, college_name.
func GetTeacherListByTeacherUid(teacher_uid uint64) (DataCenter.TeacherInfo, error) {
	var rsp DataCenter.TeacherInfo
	var ok bool
	m := make(map[uint64]DataCenter.TeacherInfo)
	if TeacherCache != nil {
		rsp, ok = (*TeacherCache)[teacher_uid]
		if ok {
			return rsp, nil
		}
	}

	// TeacherCache nil or teacher_uid not exist, query from database.
	db_m, err := dao.DataBase.Queryf("select * from `teacher` where `teacher_uid`='%d'", teacher_uid)
	if err != nil || len(db_m) != 1 {
		return rsp, errors.New("query teacher Info err")
	}
	var college_uid uint64
	var sta int

	sta, _ = strconv.Atoi(string(db_m[0]["status"].([]uint8)))
	college_uid, _ = strconv.ParseUint(string(db_m[0]["college_uid"].([]uint8)), 10, 64)
	m[teacher_uid] = DataCenter.TeacherInfo{
		TeacherUid: teacher_uid,
		CollegeUid: college_uid,
		Name:       string(db_m[0]["name"].([]uint8)),
		Password:   string(db_m[0]["password"].([]uint8)),
		Sex:        string(db_m[0]["sex"].([]uint8)),
		NRIC:       string(db_m[0]["NRIC"].([]uint8)),
		Status:     DataCenter.TeacherInfo_STATUS(sta),
		CreateTime: string(db_m[0]["create_time"].([]uint8)),
	}

	rsp = m[teacher_uid]
	return rsp, nil
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
		CreateTime: string(db_m[0]["create_time"].([]uint8)),
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
			CreateTime: string(v["create_time"].([]uint8)),
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
			CreateTime: string(v["create_time"].([]uint8)),
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

func GetCollegeUidByName(college_name string) (uint64, error) {
	cdbm, err := dao.DataBase.Queryf("SELECT 'college_uid' from `college` where `name`='%s'", college_name)
	if err != nil || len(cdbm) != 1 {
		return 0, err
	}

	college_uid, _ := strconv.ParseUint(string(cdbm[0]["college_uid"].([]uint8)), 10, 64)

	return college_uid, nil
}

func GetAllMajerName() ([]string, error) {
	cdbm, err := dao.DataBase.Queryf("select `name` from `major`")
	if err != nil || len(cdbm) == 0 {
		return nil, errors.New("query major err")
	}

	var major_name []string
	for _, val := range cdbm {
		major_name = append(major_name, string(val["name"].([]uint8)))
	}

	return major_name, nil
}

func GetMajorUidByName(major_name string) (uint64, error) {
	cdbm, err := dao.DataBase.Queryf("SELECT `major_uid` from `major` where `name`='%s'", major_name)
	if err != nil || len(cdbm) != 1 {
		return 0, err
	}

	major_uid, _ := strconv.ParseUint(string(cdbm[0]["major_uid"].([]uint8)), 10, 64)

	return major_uid, nil
}

func GetAllClassName() ([]string, error) {
	cdbm, err := dao.DataBase.Queryf("select `name` from `class`")
	if err != nil || len(cdbm) == 0 {
		return nil, errors.New("query class err")
	}

	var major_name []string
	for _, val := range cdbm {
		major_name = append(major_name, string(val["name"].([]uint8)))
	}

	return major_name, nil
}

func GetStudentByStudentUid(student_uid uint64) (DataCenter.StudentInfo, error) {
	val, ok := (*StudentCache)[student_uid]
	if ok {
		return val, nil
	}

	var s DataCenter.StudentInfo
	sm, err := dao.DataBase.Queryf("select * from `student` where `student_uid`='%d'", student_uid)
	if err != nil || len(sm) != 1 {
		return s, err
	}

	v := sm[0]
	// teacher_uid college_uid password name sex NRIC status create_time
	var college_uid, major_uid, class_uid uint64
	var sta int
	sta, _ = strconv.Atoi(string(v["status"].([]uint8)))
	college_uid, _ = strconv.ParseUint(string(v["college_uid"].([]uint8)), 10, 64)
	class_uid, _ = strconv.ParseUint(string(v["class_uid"].([]uint8)), 10, 64)
	major_uid, _ = strconv.ParseUint(string(v["major_uid"].([]uint8)), 10, 64)
	s = DataCenter.StudentInfo{
		StudentUid: student_uid,
		CollegeUid: college_uid,
		MajorUid:   major_uid,
		ClassUid:   class_uid,
		Name:       string(v["name"].([]uint8)),
		Password:   string(v["password"].([]uint8)),
		Sex:        string(v["sex"].([]uint8)),
		NRIC:       string(v["NRIC"].([]uint8)),
		Status:     DataCenter.StudentInfo_STATUS(sta),
		CreateTime: string(v["create_time"].([]uint8)),
	}
	return s, nil
}

func GetStudentByNRIC(NRIC string) (map[uint64]DataCenter.StudentInfo, error) {

	sm, err := dao.DataBase.Queryf("select * from `student` where `NRIC`='%s'", NRIC)
	if err != nil {
		return nil, err
	}

	m := make(map[uint64]DataCenter.StudentInfo)

	for _, v := range sm {
		// student_uid class_uid college_uid major_uid password name sex NRIC status create_time
		var class_uid, student_uid, college_uid, major_uid uint64
		var sta int
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
			CreateTime: string(v["create_time"].([]uint8)),
		}
	}
	return m, nil
}

func GetClassUidByName(class_name string) (uint64, error) {
	m, err := dao.DataBase.Queryf("select `class_uid` from `class` where `name`='%s'", class_name)
	if err != nil || len(m) != 1 {
		return 0, err
	}

	class_uid_str := string(m[0]["class_uid"].([]uint8))
	class_uid, _ := strconv.ParseUint(class_uid_str, 10, 64)
	return class_uid, nil
}
func GetCourseUidByName(course_name string) (uint64, error) {
	m, err := dao.DataBase.Queryf("select `course_uid` from `course` where `name`='%s'", course_name)
	if err != nil || len(m) != 1 {
		return 0, err
	}

	course_uid_str := string(m[0]["course_uid"].([]uint8))
	course_uid, _ := strconv.ParseUint(course_uid_str, 10, 64)
	return course_uid, nil
}

func GetStudentListByClassUid(class_uid uint64) (map[uint64]DataCenter.StudentInfo, error) {
	sm, err := dao.DataBase.Queryf("select * from `student` where `class_uid`='%d'", class_uid)
	if err != nil {
		return nil, err
	}
	m := make(map[uint64]DataCenter.StudentInfo)

	for _, v := range sm {
		// student_uid class_uid college_uid major_uid password name sex NRIC status create_time
		var student_uid, college_uid, major_uid uint64
		var sta int
		student_uid, _ = strconv.ParseUint(string(v["student_uid"].([]uint8)), 10, 64)
		college_uid, _ = strconv.ParseUint(string(v["college_uid"].([]uint8)), 10, 64)
		major_uid, _ = strconv.ParseUint(string(v["major_uid"].([]uint8)), 10, 64)
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
			CreateTime: string(v["create_time"].([]uint8)),
		}
	}

	return m, nil
}

func GetStudentListByMajorUid(major_uid uint64) (map[uint64]DataCenter.StudentInfo, error) {
	sm, err := dao.DataBase.Queryf("select * from `student` where `major_uid`='%d'", major_uid)
	if err != nil {
		return nil, err
	}
	m := make(map[uint64]DataCenter.StudentInfo)

	for _, v := range sm {
		// student_uid class_uid college_uid major_uid password name sex NRIC status create_time
		var student_uid, class_uid, college_uid uint64
		var sta int
		student_uid, _ = strconv.ParseUint(string(v["student_uid"].([]uint8)), 10, 64)
		class_uid, _ = strconv.ParseUint(string(v["class_uid"].([]uint8)), 10, 64)
		college_uid, _ = strconv.ParseUint(string(v["college_uid"].([]uint8)), 10, 64)
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
			CreateTime: string(v["create_time"].([]uint8)),
		}
	}

	return m, nil
}

func GetStudentListByCollegeUid(college_uid uint64) (map[uint64]DataCenter.StudentInfo, error) {
	sm, err := dao.DataBase.Queryf("select * from `student` where `college_uid`='%d'", college_uid)
	if err != nil {
		return nil, err
	}
	m := make(map[uint64]DataCenter.StudentInfo)

	for _, v := range sm {
		// student_uid class_uid college_uid major_uid password name sex NRIC status create_time
		var student_uid, class_uid, major_uid uint64
		var sta int

		student_uid, _ = strconv.ParseUint(string(v["student_uid"].([]uint8)), 10, 64)
		class_uid, _ = strconv.ParseUint(string(v["class_uid"].([]uint8)), 10, 64)
		major_uid, _ = strconv.ParseUint(string(v["major_uid"].([]uint8)), 10, 64)
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
			CreateTime: string(v["create_time"].([]uint8)),
		}
	}

	return m, nil
}

func GetStudentByName(student_name string) (map[uint64]DataCenter.StudentInfo, error) {
	sm, err := dao.DataBase.Queryf("select * from `student` where `name`='%s'", student_name)
	if err != nil {
		return nil, err
	}
	m := make(map[uint64]DataCenter.StudentInfo)

	for _, v := range sm {
		var student_uid, class_uid, college_uid, major_uid uint64
		var sta int
		student_uid, _ = strconv.ParseUint(string(v["student_uid"].([]uint8)), 10, 64)
		college_uid, _ = strconv.ParseUint(string(v["college_uid"].([]uint8)), 10, 64)
		major_uid, _ = strconv.ParseUint(string(v["major_uid"].([]uint8)), 10, 64)
		class_uid, _ = strconv.ParseUint(string(v["class_uid"].([]uint8)), 10, 64)
		m[student_uid] = DataCenter.StudentInfo{
			StudentUid: student_uid,
			CollegeUid: college_uid,
			MajorUid:   major_uid,
			ClassUid:   class_uid,
			Name:       student_name,
			Password:   string(v["password"].([]uint8)),
			Sex:        string(v["sex"].([]uint8)),
			NRIC:       string(v["NRIC"].([]uint8)),
			Status:     DataCenter.StudentInfo_STATUS(sta),
			CreateTime: string(v["create_time"].([]uint8)),
		}
	}

	return m, nil
}

func GetNamebyUid(uid uint64, table string, field string) (string, error) {
	m, err := dao.DataBase.Queryf("select `name` from `%s` where `%s`='%d'", table, field, uid)
	if err != nil || len(m) != 1 {
		return "", err
	}

	return string(m[0]["name"].([]uint8)), nil
}

func GetCollegeInfoByCollegeUid(collegeUID uint64) (DataCenter.CollegeInfo, error) {
	var result DataCenter.CollegeInfo

	cm, err := dao.DataBase.Queryf("select * from `college` where `college_uid`='%d'", collegeUID)
	if err != nil || len(cm) != 1 {
		return result, err
	}

	result = DataCenter.CollegeInfo{
		CollegeUid: collegeUID,
		Name:       string(cm[0]["name"].([]uint8)),
	}

	return result, nil
}

func GetTeacherCourseByTeacherUid(teacher_uid uint64) ([]DataCenter.CourseInfo, error) {
	m, err := dao.DataBase.Queryf("select DISTINCT `course_uid` from `student_course` where `teacher_uid`='%d'", teacher_uid)
	if err != nil {
		return nil, err
	}

	// db query result type string.
	var courses []string
	for _, v := range m {
		courses = append(courses, string(v["course_uid"].([]uint8)))
	}

	// Query course info by course uid.
	var coursesInfo []DataCenter.CourseInfo
	for _, v := range courses {
		sm, err := dao.DataBase.Queryf("select * from `course` where `course_uid`='%s'", v)
		if err != nil || len(sm) != 1 {
			continue
		}

		course_uid, _ := strconv.ParseUint(v, 10, 64)
		college_uid, _ := strconv.ParseUint(string(sm[0]["college_uid"].([]uint8)), 10, 64)
		course_name := string(sm[0]["name"].([]uint8))
		credit, _ := strconv.ParseFloat(string(sm[0]["credit"].([]uint8)), 32)
		hour, _ := strconv.ParseFloat(string(sm[0]["hour"].([]uint8)), 32)
		course_type, _ := strconv.Atoi(string(sm[0]["type"].([]uint8)))
		status, _ := strconv.Atoi(string(sm[0]["status"].([]uint8)))
		coursesInfo = append(coursesInfo, DataCenter.CourseInfo{
			CourseUid:  course_uid,
			CollegeUid: college_uid,
			Name:       course_name,
			Credit:     float32(credit),
			Hour:       float32(hour),
			Type:       DataCenter.CourseInfo_TYPE(course_type),
			Status:     DataCenter.CourseInfo_STATUS(status),
			CreateTime: string(sm[0]["create_time"].([]uint8)),
		})
		// log.Infof("%v", coursesInfo)
	}

	return coursesInfo, nil
}

func GetTeacherCourseClass(teacher_uid uint64, course_uid uint64) ([]DataCenter.ClassInfo, error) {
	m, err := dao.DataBase.Queryf("select `class_uid` from `student_course` where `teacher_uid`='%d' and `course_uid`='%d'", teacher_uid, course_uid)
	if err != nil {
		return nil, err
	}

	var classes []string
	for _, v := range m {
		classes = append(classes, string(v["class_uid"].([]uint8)))
	}

	var classInfo []DataCenter.ClassInfo
	for _, v := range classes {
		sm, err := dao.DataBase.Queryf("select * from `class` where `class_uid`='%s'", v)
		if err != nil || len(sm) != 1 {
			continue
		}

		college_uid, _ := strconv.ParseUint(string(sm[0]["college_uid"].([]uint8)), 10, 64)
		major_uid, _ := strconv.ParseUint(string(sm[0]["major_uid"].([]uint8)), 10, 64)

		class_name := string(sm[0]["name"].([]uint8))
		classInfo = append(classInfo, DataCenter.ClassInfo{
			CollegeUid: college_uid,
			Name:       class_name,
			MajorUid:   major_uid,
			CreateTime: string(sm[0]["create_time"].([]uint8)),
		})
	}

	return classInfo, nil
}

func GetCourseNameByCourseUid(course_uid uint64) string {
	m, err := dao.DataBase.Queryf("select `name` from `course` where `course_uid`='%d'", course_uid)
	if err != nil || len(m) != 1 {
		return ""
	}

	return string(m[0]["name"].([]uint8))
}
func GetTeacherClassByTeacherUid(teacher_uid uint64) ([]DataCenter.ClassInfo, error) {
	m, err := dao.DataBase.Queryf("select `class_uid` from `student_course` where `teacher_uid`='%d'", teacher_uid)
	if err != nil {
		return nil, err
	}

	var classes []string
	for _, v := range m {
		classes = append(classes, string(v["class_uid"].([]uint8)))
	}

	var classInfo []DataCenter.ClassInfo
	for _, v := range classes {
		sm, err := dao.DataBase.Queryf("select * from `class` where `class_uid`='%s'", v)
		if err != nil || len(sm) != 1 {
			continue
		}

		college_uid, _ := strconv.ParseUint(string(sm[0]["college_uid"].([]uint8)), 10, 64)
		major_uid, _ := strconv.ParseUint(string(sm[0]["major_uid"].([]uint8)), 10, 64)
		class_uid, _ := strconv.ParseUint(string(sm[0]["class_uid"].([]uint8)), 10, 64)
		class_name := string(sm[0]["name"].([]uint8))
		classInfo = append(classInfo, DataCenter.ClassInfo{
			CollegeUid: college_uid,
			ClassUid:   class_uid,
			Name:       class_name,
			MajorUid:   major_uid,
			CreateTime: string(sm[0]["create_time"].([]uint8)),
		})
	}

	return classInfo, nil
}

func IsCourseBelongClass(course_uid uint64, class_uid uint64) bool {
	m, err := dao.DataBase.Queryf("select `student_uid` from `student_course` where `course_uid`='%d' and `class_uid`='%d'", course_uid, class_uid)
	if err != nil || len(m) == 0 {
		return false
	}
	return true
}

func GetCourseScorePercentByCourseUid(course_uid uint64) (DataCenter.CourseScorePercentInfo, error) {
	var ret DataCenter.CourseScorePercentInfo
	m, err := dao.DataBase.Queryf("select * from `course_score_percent` where `course_uid`='%d'", course_uid)
	if err != nil || len(m) != 1 {
		return ret, errors.New("query course score percent err")
	}

	course_score_percent_uid, _ := strconv.ParseUint(string(m[0]["course_score_percent_uid"].([]uint8)), 10, 64)
	usual_percent, _ := strconv.ParseUint(string(m[0]["usual_percent"].([]uint8)), 10, 32)
	mid_percent, _ := strconv.ParseUint(string(m[0]["mid_percent"].([]uint8)), 10, 32)
	end_percent, _ := strconv.ParseUint(string(m[0]["end_percent"].([]uint8)), 10, 32)
	ret = DataCenter.CourseScorePercentInfo{
		CourseScorePercentUid: course_score_percent_uid,
		CourseUid:             course_uid,
		UsualPercent:          uint32(usual_percent),
		MidPercent:            uint32(mid_percent),
		EndPercent:            uint32(end_percent),
		CreateTime:            string(m[0]["create_time"].([]uint8)),
	}

	return ret, nil
}

func IsCourseHavePercent(course_uid uint64) bool {
	m, err := dao.DataBase.Queryf("select * from `course_score_percent` where `course_uid`='%d'", course_uid)
	if err != nil || len(m) != 1 {
		return false
	}

	return true
}

func InsertCoursePercent(course_uid uint64, usual_percent uint32, mid_percent uint32, end_percent uint32, percent_type uint32) error {
	err := dao.DataBase.Execf("insert into `course_score_percent` (`course_uid`, `usual_percent`, `mid_percent`, `end_percent`, `type`) values ('%d', '%d', '%d', '%d', '%d')", course_uid, usual_percent, mid_percent, end_percent, percent_type)
	return err
}

func GetScoreByStudentUidAndCourseUid(student_uid uint64, course_uid uint64) (DataCenter.ScoreInfo, error) {
	var ret DataCenter.ScoreInfo
	m, err := dao.DataBase.Queryf("select * from `score` where `student_uid`='%d' and `course_uid`='%d'", student_uid, course_uid)
	if err != nil || len(m) == 0 {
		return ret, errors.New("query score err")
	}

	val := m[0]
	score_uid, _ := strconv.ParseUint(string(val["score_uid"].([]uint8)), 10, 64)
	usual_score, _ := strconv.ParseFloat(string(val["usual_score"].([]uint8)), 32)
	mid_score, _ := strconv.ParseFloat(string(val["midterm_score"].([]uint8)), 32)
	end_score, _ := strconv.ParseFloat(string(val["endterm_score"].([]uint8)), 32)
	academic_credit, _ := strconv.ParseFloat(string(val["academic_credit"].([]uint8)), 32)
	credit, _ := strconv.ParseFloat(string(val["credit"].([]uint8)), 32)
	status, _ := strconv.Atoi(string(val["status"].([]uint8)))

	ret = DataCenter.ScoreInfo{
		ScoreUid:       score_uid,
		StudentUid:     student_uid,
		CourseUid:      course_uid,
		MidtermScore:   float32(mid_score),
		EndtermScore:   float32(end_score),
		UsualScore:     float32(usual_score),
		AcademicCredit: float32(academic_credit),
		Credit:         float32(credit),
		Status:         DataCenter.ScoreInfo_STATUS(status),
		CreateTime:     string(val["create_time"].([]uint8)),
	}

	return ret, nil
}

func GetCoursePercent(course_uid uint64) (DataCenter.CourseScorePercentInfo, error) {
	var ret DataCenter.CourseScorePercentInfo
	m, err := dao.DataBase.Queryf("select * from `course_score_percent` where `course_uid`='%d'", course_uid)
	if err != nil || len(m) != 1 {
		return ret, err
	}
	usual_percent, _ := strconv.Atoi(string(m[0]["usual_percent"].([]uint8)))
	mid_percent, _ := strconv.Atoi(string(m[0]["mid_percent"].([]uint8)))
	end_percent, _ := strconv.Atoi(string(m[0]["end_percent"].([]uint8)))
	percent_type, _ := strconv.Atoi(string(m[0]["type"].([]uint8)))

	ret = DataCenter.CourseScorePercentInfo{
		CourseUid:    course_uid,
		UsualPercent: uint32(usual_percent),
		MidPercent:   uint32(mid_percent),
		EndPercent:   uint32(end_percent),
		Type:         DataCenter.CourseScorePercentInfo_PERCENT_TYPE(percent_type),
	}

	return ret, nil
}

// 获取全班同学的某门课成绩
func GetStudentScoreByClassUidAndCourseUid(class_uid uint64, course_uid uint64) ([]DataCenter.ScoreInfo, error) {
	m, err := dao.DataBase.Queryf("select `student_uid` from `student` where `class_uid`='%d'", class_uid)
	if err != nil || len(m) == 0 {
		return nil, errors.New("Db Query err")
	}
	var students []uint64

	for _, v := range m {
		student_uid, _ := strconv.ParseUint(string(v["student_uid"].([]uint8)), 10, 64)
		students = append(students, student_uid)
	}

	var result []DataCenter.ScoreInfo
	for _, student := range students {
		score, err := dao.DataBase.Queryf("select * from `score` where `student_uid`='%d' and `course_uid`='%d'", student, course_uid)
		if err != nil || len(score) != 1 {
			log.Error("score query err", err)
			continue
		}

		mid_s, _ := strconv.ParseFloat(string(score[0]["midterm_score"].([]uint8)), 32)
		usu_s, _ := strconv.ParseFloat(string(score[0]["usual_score"].([]uint8)), 32)
		end_s, _ := strconv.ParseFloat(string(score[0]["endterm_score"].([]uint8)), 32)
		s, _ := strconv.Atoi(string(score[0]["score"].([]uint8)))
		type_s, _ := strconv.Atoi(string(score[0]["score_type"].([]uint8)))
		status, _ := strconv.Atoi(string(score[0]["status"].([]uint8)))
		student_score := DataCenter.ScoreInfo{
			StudentUid:   student,
			CourseUid:    course_uid,
			MidtermScore: float32(mid_s),
			UsualScore:   float32(usu_s),
			EndtermScore: float32(end_s),
			Score:        uint32(s),
			ScoreType:    DataCenter.ScoreInfo_SCORE_TYPE(type_s),
			Status:       DataCenter.ScoreInfo_STATUS(status),
		}
		result = append(result, student_score)
	}
	return result, nil
}

func GetStudentSubmitScoreByStudentUid(student_uid uint64) ([]DataCenter.ScoreInfo, error) {
	var result []DataCenter.ScoreInfo
	m, err := dao.DataBase.Queryf("select * from `score` where `student_uid`='%d' and `type`='1'", student_uid)
	if err != nil {
		return nil, err
	}
	for _, v := range m {
		course_uid, _ := strconv.ParseUint(string(v["course_uid"].([]uint8)), 10, 64)
		ms, _ := strconv.ParseFloat(string(v["midterm_score"].([]uint8)), 32)
		us, _ := strconv.ParseFloat(string(v["usual_score"].([]uint8)), 32)
		es, _ := strconv.ParseFloat(string(v["endterm_score"].([]uint8)), 32)
		s, _ := strconv.Atoi(string(v["score"].([]uint8)))
		ac, _ := strconv.ParseFloat(string(v["academic_credit"].([]uint8)), 32)
		c, _ := strconv.ParseFloat(string(v["credit"].([]uint8)), 32)
		st, _ := strconv.Atoi(string(v["score_type"].([]uint8)))
		result = append(result, DataCenter.ScoreInfo{
			StudentUid:     student_uid,
			CourseUid:      course_uid,
			MidtermScore:   float32(ms),
			UsualScore:     float32(us),
			EndtermScore:   float32(es),
			Score:          uint32(s),
			AcademicCredit: float32(ac),
			Credit:         float32(c),
			ScoreType:      DataCenter.ScoreInfo_SCORE_TYPE(st),
		})
	}
	return result, nil
}

func GetCourseByClassUid(class_uid uint64) ([]DataCenter.CourseInfo, error) {
	var result []DataCenter.CourseInfo
	m, err := dao.DataBase.Queryf("select * from `student_course` where `class_uid`='%d'", class_uid)
	if err != nil {
		return result, err
	}
	var courses []uint64
	for _, v := range m {
		course_uid, _ := strconv.ParseUint(string(v["course_uid"].([]uint8)), 10, 64)
		courses = append(courses, course_uid)
	}

	for _, v := range courses {
		cm, err := dao.DataBase.Queryf("select * from `course` where `course_uid`='%d'", v)
		if err != nil {
			continue
		}
		college_uid, _ := strconv.ParseUint(string(cm[0]["college_uid"].([]uint8)), 10, 64)
		credit, _ := strconv.ParseFloat(string(cm[0]["credit"].([]uint8)), 32)
		hour, _ := strconv.ParseFloat(string(cm[0]["hour"].([]uint8)), 32)
		type_c, _ := strconv.Atoi(string(cm[0]["type"].([]uint8)))
		status, _ := strconv.Atoi(string(cm[0]["status"].([]uint8)))
		result = append(result, DataCenter.CourseInfo{
			CourseUid:  v,
			CollegeUid: college_uid,
			Name:       string(cm[0]["name"].([]uint8)),
			Credit:     float32(credit),
			Hour:       float32(hour),
			Type:       DataCenter.CourseInfo_TYPE(type_c),
			Status:     DataCenter.CourseInfo_STATUS(status),
		})
	}

	return result, nil
}

func GetAllClassInfo() ([]DataCenter.ClassInfo, error) {
	var result []DataCenter.ClassInfo
	m, err := dao.DataBase.Queryf("select * from `class`")
	if err != nil {
		return result, err
	}

	for _, v := range m {
		classUID, _ := strconv.ParseUint(string(v["class_uid"].([]uint8)), 10, 64)
		collegeUID, _ := strconv.ParseUint(string(v["college_uid"].([]uint8)), 10, 64)
		majorUID, _ := strconv.ParseUint(string(v["major_uid"].([]uint8)), 10, 64)
		result = append(result, DataCenter.ClassInfo{
			ClassUid:   classUID,
			CollegeUid: collegeUID,
			MajorUid:   majorUID,
			Name:       string(v["name"].([]uint8)),
		})
	}
	return result, nil
}

func GetAllCourseInfo() ([]DataCenter.CourseInfo, error) {
	var result []DataCenter.CourseInfo
	m, err := dao.DataBase.Queryf("select * from `course`")
	if err != nil {
		return result, err
	}
	for _, v := range m {
		courseUID, _ := strconv.ParseUint(string(v["course_uid"].([]uint8)), 10, 64)
		collegeUID, _ := strconv.ParseUint(string(v["college_uid"].([]uint8)), 10, 64)
		credit, _ := strconv.ParseFloat(string(v["credit"].([]uint8)), 32)
		status, _ := strconv.Atoi(string(v["status"].([]uint8)))
		typeC, _ := strconv.Atoi(string(v["type"].([]uint8)))
		result = append(result, DataCenter.CourseInfo{
			CourseUid:  courseUID,
			CollegeUid: collegeUID,
			Name:       string(v["name"].([]uint8)),
			Credit:     float32(credit),
			Status:     DataCenter.CourseInfo_STATUS(status),
			Type:       DataCenter.CourseInfo_TYPE(typeC),
		})
	}

	return result, nil
}

func GetCourseByStudentUid(studentUID uint64) ([]DataCenter.CourseInfo, error) {
	var result []DataCenter.CourseInfo
	m, err := dao.DataBase.Queryf("select * from `student_course` where `student_uid`='%d'", studentUID)
	if err != nil {
		return result, err
	}
	var courseUIDs []uint64
	for _, v := range m {
		courseUID, _ := strconv.ParseUint(string(v["course_uid"].([]uint8)), 10, 64)
		courseUIDs = append(courseUIDs, courseUID)
	}

	for _, v := range courseUIDs {
		cm, err := dao.DataBase.Queryf("select * from `course` where `course_uid`='%d'", v)
		if err != nil {
			continue
		}
		collegeUID, _ := strconv.ParseUint(string(cm[0]["college_uid"].([]uint8)), 10, 64)
		credit, _ := strconv.ParseFloat(string(cm[0]["credit"].([]uint8)), 32)
		hour, _ := strconv.ParseFloat(string(cm[0]["hour"].([]uint8)), 32)
		courseName := string(cm[0]["name"].([]uint8))
		typeC, _ := strconv.Atoi(string(cm[0]["type"].([]uint8)))
		status, _ := strconv.Atoi(string(cm[0]["status"].([]uint8)))
		result = append(result, DataCenter.CourseInfo{
			CourseUid:  v,
			CollegeUid: collegeUID,
			Credit:     float32(credit),
			Hour:       float32(hour),
			Name:       courseName,
			Type:       DataCenter.CourseInfo_TYPE(typeC),
			Status:     DataCenter.CourseInfo_STATUS(status),
		})
	}

	return result, nil
}
