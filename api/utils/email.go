package utils

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"math/rand"
	"time"
)

// SendEmail 发送邮箱
func SendEmail(smtpServer, authUsername, authPassword, from, to, body, subject, contentType string, smtpPort int) (err error) {

	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody(contentType, body)

	// 设置SMTP服务器和端口
	dialer := gomail.NewDialer(smtpServer, smtpPort, authUsername, authPassword)

	// 启用SSL/TLS
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// 发送邮件
	if err = dialer.DialAndSend(message); err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	return
}

// GenerateVerificationCode 生成6位随机数字验证码
func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code)
}

func BuildEmailTemplate(link string) string {
	// 定义模板内容
	templateText := `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Email Template</title>
</head>
<body>
	<h1>重置密码</h1>
	<p>点击以下链接重置您的密码：</p>
	<a href="{{ .Link }}">{{ .Link }}</a>
</body>
</html>
	`
	// 准备数据
	data := struct {
		Link string
	}{
		Link: link,
	}
	// 创建模板
	tmpl := template.Must(template.New("email").Parse(templateText))
	// 渲染模板并捕获输出
	var output bytes.Buffer
	err := tmpl.Execute(&output, data)
	if err != nil {
		return ""
	}
	return output.String()

}

func GenerateInviteCode(email string) string {
	// 创建哈希对象
	hash := sha256.New()
	// 将邮箱地址转换为字节数组并写入哈希对象
	hash.Write([]byte(email))
	// 计算哈希值并转换为十六进制格式
	hashed := hash.Sum(nil)
	// 取前6位作为邀请码
	inviteCode := hex.EncodeToString(hashed)[:6]
	return inviteCode
}
