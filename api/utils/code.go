package utils

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pkg/errors"
)

func CreateClient(accessKeyId, accessKeySecret, endpoint *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = endpoint
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func Send(sendSmsRequest *dysmsapi20170525.SendSmsRequest, ak, sk, endpoint string) (_err error) {
	client, _err := CreateClient(tea.String(ak), tea.String(sk), tea.String(endpoint))
	if _err != nil {
		return _err
	}
	runtime := &util.RuntimeOptions{}
	return func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		res, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}
		if "OK" != *res.Body.Message || "OK" != *res.Body.Code {
			return errors.New(*res.Body.Message)
		}
		return nil
	}()
}
