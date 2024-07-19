package utils

import (
	"fmt"
	"testing"
)

func TestSendCodeEmail(t *testing.T) {
	smtpServer := "smtp.feishu.cn"
	smtpPort := 465
	authUsername := "noreply@wapen.app"
	authPassword := "94o1baq75TFix1BK"
	from := "noreply@wapen.app"
	code := GenerateVerificationCode()
	body := "Your verification code is: " + code
	subject := "Verification Code"
	err := SendEmail(smtpServer, authUsername, authPassword, from, "414685046@qq.com", body, subject, "text/plain", smtpPort)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(code)
}

func TestSendResetPwdEmail(t *testing.T) {

	smtpServer := "smtp.feishu.cn"
	smtpPort := 465
	authUsername := "noreply@wapen.app"
	authPassword := "94o1baq75TFix1BK"
	from := "noreply@wapen.app"
	uid := GetUUID()
	err := SendEmail(smtpServer, authUsername, authPassword, from, "414685046@qq.com", BuildEmailTemplate(fmt.Sprintf("https://pen.wapen.app/reset-password?uid=%s&email=%s", uid, "123@qq.com")), "重置密码", "text/html", smtpPort)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(uid)
}

func TestGenerateInviteCode(t *testing.T) {
	code := GenerateInviteCode("wapen-41468504@qq.com")
	t.Log(code)
}
