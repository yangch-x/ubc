package utils

import (
	"math/rand"
	"time"
)

func IsListContains(data []string, match string) bool {
	if len(data) == 0 {
		return false
	}
	for id := range data {
		if data[id] == match {
			return true
		}
	}
	return false
}

func IsNumberInSlice(slice []int, number int) bool {
	for _, value := range slice {
		if value == number {
			return true
		}
	}
	return false
}

func ListRemoveItem(data []string, match string) []string {
	for i := 0; i < len(data); i++ {
		if data[i] == match {
			data = append(data[:i], data[i+1:]...)
			i--
		}
	}
	return data
}

func ListRemoveSubList(data []string, sub []string) []string {
	for i := 0; i < len(sub); i++ {
		data = ListRemoveItem(data, sub[i])
	}
	return data
}

// ListContainsSubList 原始数组是否包含整个子数组
func ListContainsSubList(data []string, sub []string) bool {
	for i := 0; i < len(sub); i++ {
		if !IsListContains(data, sub[i]) {
			return false
		}
	}
	return true
}

func ListIntersection(data1 []string, data2 []string) []string {
	var data []string

	for i := 0; i < len(data1); i++ {
		for j := 0; j < len(data2); j++ {
			if data1[i] == data2[j] {
				data = append(data, data1[i])
			}
		}
	}
	return data
}

func SelectRandomElements(num int, elements []string) []string {
	rand.Seed(time.Now().UnixNano())

	if num <= 0 || num > len(elements) {
		return nil
	}

	copyOfElements := make([]string, len(elements))
	copy(copyOfElements, elements)

	var selected []string

	for i := 0; i < num; i++ {
		// 生成一个随机索引
		randomIndex := rand.Intn(len(copyOfElements))

		selected = append(selected, copyOfElements[randomIndex])

		copyOfElements = append(copyOfElements[:randomIndex], copyOfElements[randomIndex+1:]...)
	}

	return selected
}
