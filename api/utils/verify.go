package utils

import (
	"regexp"

	"github.com/nyaruka/phonenumbers"
	"google.golang.org/protobuf/proto"
)

// VerifyEmail email verify
func VerifyEmail(email string) bool {
	pattern := `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$` //匹配电子邮箱
	//pattern := `^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyMobileZh mobile verify
func VerifyMobileZh(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func VerifyMobile(countryCode, mobileNum string) bool {

	number := &phonenumbers.PhoneNumber{
		CountryCode:    proto.Int32(int32(ParseInt(countryCode))),
		NationalNumber: proto.Uint64(uint64(ParseInt64(mobileNum))),
	}
	return phonenumbers.IsValidNumber(number)
}
