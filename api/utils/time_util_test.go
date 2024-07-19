package utils

import (
	"fmt"
	"testing"
)

func TestConvertToUTC(t *testing.T) {

	timeStr := "2024-03-02 02:20:10 Etc/GMT"
	layout := "2006-01-02 15:04:05"

	utcTime, err := ConvertToUTC(timeStr, layout)
	if err != nil {
		fmt.Println("转换时出错:", err)
		return
	}

	fmt.Println("UTC 时间:", utcTime)
}

func TestIsAfterCurrentDate(t *testing.T) {

}
