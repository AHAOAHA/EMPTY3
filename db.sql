CREATE DATABASE GrageManager;
USE `GradeManager`;
-- college表创建
CREATE TABLE IF NOT EXISTS `college` (
  `college_uid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `name` varchar(64) NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`college_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--major表创建
CREATE TABLE IF NOT EXISTS `major` (
	`major_uid` BIGINT UNSIGNED NOT NULL,
	`college_uid` BIGINT UNSIGNED,
	`name` VARCHAR(64) NOT NULL,
	`create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`major_uid`),
	CONSTRAINT sc_college_uid FOREIGN KEY (`college_uid`) REFERENCES college(`college_uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8;

-- class表创建
CREATE TABLE IF NOT EXISTS `class` (
	`class_uid` BIGINT UNSIGNED NOT NULL,
	`college_uid` BIGINT UNSIGNED,
	`major_uid` BIGINT UNSIGNED,
	`name` VARCHAR(64) NOT NULL,
	`create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`class_uid`),
	CONSTRAINT c_sc_college_uid FOREIGN KEY (`college_uid`) REFERENCES college(`college_uid`),
	CONSTRAINT c_sc_major_uid FOREIGN KEY (`major_uid`) REFERENCES major(`major_uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8;

-- student表创建
CREATE TABLE IF NOT EXISTS `student` (
	`student_uid` BIGINT UNSIGNED NOT NULL,
	`class_uid` BIGINT UNSIGNED,
	`college_uid` BIGINT UNSIGNED,
	`major_uid` BIGINT UNSIGNED,
	`password` VARCHAR(48) NOT NULL,
	`name` VARCHAR(64) NOT NULL DEFAULT '',
	`sex` VARCHAR(8) NOT NULL DEFAULT '',
	`NRIC` VARCHAR(48) NOT NULL,
	`status` INT UNSIGNED NOT NULL DEFAULT '0',
	`create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`student_uid`),
	CONSTRAINT s_sc_college_uid FOREIGN KEY (`college_uid`) REFERENCES college(`college_uid`),
	CONSTRAINT s_sc_major_uid FOREIGN KEY (`major_uid`) REFERENCES major(`major_uid`),
	CONSTRAINT s_sc_class_uid FOREIGN KEY (`class_uid`) REFERENCES class(`class_uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8;

-- score 表创建
CREATE TABLE IF NOT EXISTS `score` (
	`score_uid` BIGINT UNSIGNED NOT NULL,
	`student_uid` BIGINT UNSIGNED,
	`course_uid` BIGINT UNSIGNED,
	`midterm_score` FLOAT NOT NULL DEFAULT '0',
	`endterm_score` FLOAT NOT NULL DEFAULT '0',
	`usual_score` FLOAT NOT NULL DEFAULT '0',
	`academic_credit` FLOAT NOT NULL DEFAULT '0',
	`credit` FLOAT NOT NULL DEFAULT '0',
	`status` INT UNSIGNED NOT NULL DEFAULT '0',
	`create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`score_uid`),
	CONSTRAINT score_sc_student_uid FOREIGN KEY (`student_uid`) REFERENCES student(`student_uid`),
	CONSTRAINT score_sc_course_uid FOREIGN KEY (`course_uid`) REFERENCES course(`course_uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `student_score` (
	`student_score_uid` BIGINT UNSIGNED NOT NULL,
	`student_uid` BIGINT UNSIGNED,
	`course_uid` BIGINT UNSIGNED,
	`create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`student_score_uid`),
	CONSTRAINT ss_sc_student_uid FOREIGN KEY (`student_uid`) REFERENCES student(`student_uid`),
	CONSTRAINT ss_sc_course_uid FOREIGN KEY (`course_uid`) REFERENCES course(`course_uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8;