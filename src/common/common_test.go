package common

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestSendTextToWechat(t *testing.T) {
	rand.Seed(time.Now().Unix())
	text := time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006")
	desp := strconv.Itoa(rand.Intn(100))

	err := SendTextToWechat(text, desp)
	if err != nil {
		t.Error(err)
	}

	t.Log("success")
}
