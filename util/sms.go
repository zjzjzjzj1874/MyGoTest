package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SendSms(mobile, message string) (bool, error) {
	data := url.Values{}
	data.Set("mobile", mobile)
	data.Set("message", message+"【朋友房】")


	req, err := http.NewRequest("POST", "http://sms-api.luosimao.com/v1/send.json", strings.NewReader(data.Encode()))
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("api", "key-7eec90e0e2afd5c247302845ce0cce1c")
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	type res struct {
		error int
		msg   string
	}
	r := res{}
	err = json.Unmarshal(result, &r)
	if err != nil || r.error != 0 {
		return false, err
	}
	return true, nil
}