package xerr

import (
	"fmt"
	"google.golang.org/grpc/status"
	"strconv"
	"sync/atomic"

	"github.com/pkg/errors"
)

/*
*
常用通用固定错误
*/
var (
	_messages atomic.Value            // NOTE: stored map[string]map[int]string
	_codes    = map[uint32]struct{}{} // register codes.
)

type Code uint32

// Register register ecode message map.
func Register(cm map[uint32]string) {
	_messages.Store(cm)
}

func New(e uint32) Code {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return add(e)
}

func add(e uint32) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return Int(e)
}

func Add(e uint32) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return Int(e)
}

func Int(i uint32) Code { return Code(i) }

func (e Code) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

func (e Code) Equal(err error) bool { return EqualError(e, err) }

func (e Code) Code() uint32 { return uint32(e) }

func (e Code) Message() string {
	if cm, ok := _messages.Load().(map[uint32]string); ok {
		if msg, ok := cm[e.Code()]; ok {
			return msg
		}
	}
	return e.Error()
}

func IsCodeErr(e uint32) bool {
	if _, ok := _codes[e]; ok {
		return true
	}
	return false
}

// Equal equal a and b by code int.
func Equal(a, b Code) bool {
	return a.Code() == b.Code()
}

// EqualError equal error
func EqualError(code Code, err error) bool {

	causeErr := errors.Cause(err)                      // err类型
	if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
		grpcCode := uint32(gstatus.Code())
		return Equal(code, Code(grpcCode))
	}
	return causeErr.Error() == code.Error()
}
