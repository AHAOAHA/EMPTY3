/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: log.go
 * Author: ahaoozhang
 * Date: 2020-06-22 16:59:02 (Monday)
 * Describe:
 ******************************************************************/

package logFileWriter

import (
	"errors"
	"log"
	"os"
	"time"
)

type LogFileWriter struct {
	File       *os.File
	createTime time.Time
}

func CeateNewLogFileWriter() (*LogFileWriter, error) {
	file, err := os.OpenFile("./log/"+time.Now().Format("2006-01-02.PM.log"), os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		log.Fatal("file not opened")
		return nil, err
	}

	return &LogFileWriter{file, time.Now()}, nil
}

func isNextDay(oldTime time.Time) bool {
	oldTimeStr := oldTime.Format("2016-01-02")
	newTimeStr := time.Now().Format("2016-01-02")
	if oldTimeStr != newTimeStr {
		return true
	}
	return false
}

func (p *LogFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.File == nil {
		return 0, errors.New("file not opened")
	}
	n, e := p.File.Write(data)

	if isNextDay(p.createTime) {
		p.File.Close()
		tmp, e := CeateNewLogFileWriter()
		if e != nil {
			log.Fatal(e)
		}

		p.File = tmp.File
		p.createTime = tmp.createTime
	}
	return n, e
}
