package utils

import (
	"fmt"
	"testing"
)

func TestSelectRandomElements(t *testing.T) {
	strArray := []string{"苹果", "橙子", "香蕉", "西瓜", "葡萄", "梨子"}

	selected := SelectRandomElements(3, strArray)
	if selected != nil {
		fmt.Println("随机选取的三个元素:", selected)
	} else {
		fmt.Println("参数无效")
	}
}
