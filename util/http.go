package util

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 发送带头的http get请求,可升级
func SendGETReqWithHeader(url,token string) (err error,b []byte) {
	client := &http.Client{}
	//提交请求
	request, err := http.NewRequest("GET", url, nil)
	//增加header选项
	request.Header.Add("token", token)
	if err != nil {
		return err,nil
	}
	//处理返回结果
	response, err := client.Do(request)
	defer response.Body.Close()
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	return err,responseBodyBytes
}
// 发送简单GET请求
// url:请求地址
// 返回值:失败/成功,nil/成功信息
func Get(url string) (bool,string) {
	if !strings.HasPrefix(url,"http://"){
		url = "http://" + url
	}
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return false,""
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return false,""
		}
	}
	response := result.String()
	return true,response
}
