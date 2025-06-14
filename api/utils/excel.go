package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Export 将数据导出为 Excel 文件流
// T 为任意类型
// data 参数为 T 类型的指针数组
// name 为导出的文件名
func Export[T any](data []*T, name string, sheetName string) ([]byte, error) {
	if !strings.HasSuffix(name, ".xlsx") {
		name += ".xlsx"
	}

	f := excelize.NewFile()
	defer f.Close()

	// 确保创建的工作表为第一个
	if err := f.SetSheetName("Sheet1", sheetName); err != nil {
		return nil, fmt.Errorf("failed to rename sheet: %w", err)
	}

	// 检查数据是否为空
	if len(data) == 0 {
		return nil, fmt.Errorf("data slice is empty")
	}

	val := reflect.ValueOf(data[0]).Elem()
	typeOfT := val.Type()

	// 创建表头样式（设置背景色为灰色，加粗，居中对齐）
	headerStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:  "solid",
			Color: []string{"CCCCCC"}, // 背景色为灰色
		},
		Font: &excelize.Font{
			Bold: true, // 加粗
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center", // 水平居中
			Vertical:   "center", // 垂直居中
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create header style: %w", err)
	}

	// 添加标题行
	for i := 0; i < typeOfT.NumField(); i++ {
		field := typeOfT.Field(i)

		// 判断是否有tag: `excel:"-"`
		if tag := field.Tag.Get("excel"); tag == "-" {
			continue // 跳过不需要写入的字段
		}

		cell, err := excelize.CoordinatesToCellName(i+1, 1)
		if err != nil {
			return nil, err
		}
		if err := f.SetCellValue(sheetName, cell, field.Name); err != nil {
			return nil, fmt.Errorf("failed to set header value: %w", err)
		}
		if err := f.SetCellStyle(sheetName, cell, cell, headerStyle); err != nil {
			return nil, fmt.Errorf("failed to set header style: %w", err)
		}
	}

	// 设置行高
	if err := f.SetRowHeight(sheetName, 1, 20); err != nil {
		return nil, fmt.Errorf("failed to set row height: %w", err)
	}

	// 填充数据
	for rowIndex, item := range data {
		val := reflect.ValueOf(item).Elem()
		typ := val.Type()
		for colIndex := 0; colIndex < val.NumField(); colIndex++ {
			field := typ.Field(colIndex)
			// 跳过带有 `excel:"-"` 标签的字段
			if tag := field.Tag.Get("excel"); tag == "-" {
				continue
			}
			cell, err := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			if err != nil {
				return nil, err
			}
			if err := f.SetCellValue(sheetName, cell, val.Field(colIndex).Interface()); err != nil {
				return nil, fmt.Errorf("failed to set cell value: %w", err)
			}
		}
	}

	// 写入到 buffer
	buf := new(bytes.Buffer)
	if err := f.Write(buf); err != nil {
		return nil, fmt.Errorf("failed to write file to buffer: %w", err)
	}

	return buf.Bytes(), nil
}
