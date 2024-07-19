package utils

import (
	"fmt"
	"testing"
)

func TestVerifyMobile(t *testing.T) {

	k := VerifyMobile("65", "83322483")

	e := VerifyEmail("sudolewiss..ad@gmail.comss")
	fmt.Println(e)
	//n := IsNum("1111111111")
	t.Log(k)
	//t.Log(e)
	//t.Log(n)
}
