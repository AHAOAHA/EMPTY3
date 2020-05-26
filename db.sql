CREATE DATABASE GrageManager;
USE `GradeManager`;
-- admin表创建
CREATE TABLE `admin` (
  `uid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user` varchar(32) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `mail` varchar(50) NOT NULL DEFAULT '',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `expr_time` int(11) DEFAULT NULL,
  `login_ip` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`uid`),
  UNIQUE KEY `shop_admin_user_password` (`user`,`password`),
  UNIQUE KEY `shop_admin_user_mail` (`user`,`mail`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- teacher表创建
CREATE TABLE `teacher` (
  `teacher_uid` bigint(20) unsigned NOT NULL,
  `college_uid` bigint(20) unsigned DEFAULT NULL,
  `password` varchar(48) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `sex` varchar(8) NOT NULL DEFAULT '',
  `NRIC` varchar(48) NOT NULL,
  `status` int(10) unsigned NOT NULL DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`teacher_uid`),
  KEY `t_sc_college_uid` (`college_uid`),
  CONSTRAINT `t_sc_college_uid` FOREIGN KEY (`college_uid`) REFERENCES `college` (`college_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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

-- course 表创建
CREATE TABLE `course` (
  `course_uid` bigint(20) unsigned NOT NULL,
  `college_uid` bigint(20) unsigned DEFAULT NULL,
  `name` varchar(64) NOT NULL,
  `credit` float NOT NULL,
  `hour` float NOT NULL,
  `type` float NOT NULL,
  `status` int(10) unsigned NOT NULL DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`course_uid`),
  KEY `co_sc_college_uid` (`college_uid`),
  CONSTRAINT `co_sc_college_uid` FOREIGN KEY (`college_uid`) REFERENCES `college` (`college_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- nitice 表创建
CREATE TABLE `notice` (
  `title` varchar(256) NOT NULL DEFAULT 'title',
  `data` text NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `notice_uid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`notice_uid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- score 表创建
CREATE TABLE `score` (
  `score_uid` bigint(20) unsigned NOT NULL,
  `student_uid` bigint(20) unsigned DEFAULT NULL,
  `course_uid` bigint(20) unsigned DEFAULT NULL,
  `midterm_score` float NOT NULL DEFAULT '0',
  `endterm_score` float NOT NULL DEFAULT '0',
  `usual_score` float NOT NULL DEFAULT '0',
  `academic_credit` float NOT NULL DEFAULT '0',
  `credit` float NOT NULL DEFAULT '0',
  `status` int(10) unsigned NOT NULL DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `score` int(10) unsigned NOT NULL DEFAULT '0',
  `type` int(10) unsigned NOT NULL DEFAULT '0',
  `score_type` int(10) unsigned NOT NULL DEFAULT '0',
  `end_percent` int(10) unsigned NOT NULL DEFAULT '0',
  `mid_percent` int(10) unsigned NOT NULL DEFAULT '0',
  `usual_percent` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`score_uid`),
  KEY `score_sc_student_uid` (`student_uid`),
  KEY `score_sc_course_uid` (`course_uid`),
  CONSTRAINT `score_sc_course_uid` FOREIGN KEY (`course_uid`) REFERENCES `course` (`course_uid`),
  CONSTRAINT `score_sc_student_uid` FOREIGN KEY (`student_uid`) REFERENCES `student` (`student_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `student_course` (
	`student_score_uid` BIGINT UNSIGNED NOT NULL,
	`student_uid` BIGINT UNSIGNED,
	`course_uid` BIGINT UNSIGNED,
	`create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`student_score_uid`),
	CONSTRAINT ss_sc_student_uid FOREIGN KEY (`student_uid`) REFERENCES student(`student_uid`),
	CONSTRAINT ss_sc_course_uid FOREIGN KEY (`course_uid`) REFERENCES course(`course_uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8;

CREATE  TABLE `course_score_percent` (
	`course_score_percent_uid` BIGINT UNSIGNED NOT NULL,
	`course_uid` BIGINT UNSIGNED NOT NULL, 
	`usual_percent` INT UNSIGNED NOT NULL,
	`mid_percent` INT UNSIGNED NOT NULL,
	`end_percent` INT UNSIGNED NOT NULL,
	`type` INT UNSIGNED NOT NULL DEFAULT '0',
	`create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`course_score_percent_uid`),
	CONSTRAINT csp_course_uid FOREIGN KEY (`course_uid`) REFERENCES course(`course_uid`)
)ENGINE=INNODB DEFAULT CHARSET=utf8;