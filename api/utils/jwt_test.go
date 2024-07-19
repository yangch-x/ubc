package utils

import (
	"fmt"
	"testing"
)

func TestGetJwtToken(t *testing.T) {

	token, _, _, _ := GetJwtToken("lewisreg@126.com", "", "lewisreg@126.com", "a5931ad7-0528-42f3-8807-7fb2bda3a203", "l@TiD3HFUMliNP0fIf#ebYK0Hif0k23mAg#aw@myfJAqqagaNog5ej", 604800)

	t.Log(token)

}

func TestDemo(t *testing.T) {

	ss := []int{1, 2, 3, 4, 5, 6, 7}

	ss, _ = Build(ss, 10)

	fmt.Println(ss)

}

func Build(s []int, nums int64) ([]int, error) {
	num, _ := MaxNum(s)

	if num > 20 {

		s = s[1:]
		return Build(s, nums)
	}

	return s, nil
}

func MaxNum(s []int) (int, error) {

	var max int

	for _, num := range s {
		max += num
	}

	return max, nil

}
