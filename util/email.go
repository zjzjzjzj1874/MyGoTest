// 简单发邮件方法
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package util

import (
	"net/smtp"
	"strings"
)

// 发件邮箱信息
const (
	User = "1396892127@qq.com"
	Password = "xpzsblesyfzyhahc"
	Host = "smtp.qq.com:25"
)

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

func Send(to,subject,content string)  {
	body := `
        <html>
        <body>
        <h3>
        ` + content + `
        </h3>
        </body>
        </html>
        `
	err := SendToMail(User, Password, Host, to, subject, body, "html")
	if err != nil {
		ErrorLog("Failed:send email to",to,"subject = ",subject,";content = ",content,"error = ",err)
	} else {
		InfoLog("Success:send email to",to,"subject = ",subject,";content = ",content)
	}
}
