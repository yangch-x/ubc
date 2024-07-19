package utils

import (
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"testing"
)

func TestSend(t *testing.T) {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		TemplateCode:  tea.String("SMS_190782115"),
		PhoneNumbers:  tea.String("13635466628"),
		SignName:      tea.String("Washine蛙赛面试"),
		TemplateParam: tea.String(`{"code": "123"}`),
	}
	err := Send(sendSmsRequest, "", "", "")

	if err != nil {
		panic(err)
	}
}
