/******************************************************************
 * Copyright(C) 2020-2020. All right reserved.
 *
 * Filename: commom.go
 * Author: ahaoozhang
 * Date: 2020-04-10 08:26:15 (Friday)
 * Describe:
 ******************************************************************/

package common

import (
	"GradeManager/src/config"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

func SendTextToWechat(text string, desp string) error {
	var Url string = config.Config.Alarm.Url
	reqUrl := Url + "?text=" + text + "&desp=" + desp
	u, _ := url.Parse(reqUrl)
	q := u.Query()
	u.RawQuery = q.Encode() //urlencode
	rsp, err := http.Get(u.String())
	if err != nil {
		return err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Error(string(body))
		return err
	}

	return nil
}
