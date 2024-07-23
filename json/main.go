package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io/ioutil"
)

const (
	apiKey         = "sk-proj-t8eNeNsIoceuNnPgC8AJT3BlbkFJB7D5zYJv00KhqBYZjH1I"
	inputFile      = "D:\\goproject\\UBC\\json\\cp.json"
	jsonOutputFile = "aaa.json"
	gptModel       = "gpt-4o"
)

type Order struct {
	PO            string `json:"po"`
	Color         string `json:"color"`
	StyleNumber   string `json:"style_number"`
	Sizes         []Size `json:"sizes"`
	TotalQuantity int    `json:"total_quantity"`
	CTNS          int    `json:"CTNS"`
}

type Size struct {
	Size     string `json:"size"`
	Quantity int    `json:"quantity"`
}

func main() {
	// 读取输入文件内容
	fileContent, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("无法读取输入文件:", err)
		return
	}

	client := openai.NewClient(apiKey)

	messages := []openai.ChatCompletionMessage{
		{Role: "system", Content: "Don't make assumptions about what values to plug into functions."},
		{Role: "user", Content: "Export order details to a file including various attributes like PO#, color, style#, size, quantity and more.\n" + string(fileContent)},
	}

	// 构建 ChatCompletionRequest
	req := openai.ChatCompletionRequest{
		Model:    gptModel,
		Messages: messages,
		Functions: []openai.FunctionDefinition{
			{
				Name:        "export_simple_order_details",
				Description: "Export simplified order details to a file including PO#, color, style number, sizes, and quantities.",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"orders": map[string]interface{}{
							"type":        "array",
							"description": "An array of simplified order objects",
							"items": map[string]interface{}{
								"type": "object",
								"properties": map[string]interface{}{
									"po":           map[string]interface{}{"type": "string", "description": "Purchase order number"},
									"color":        map[string]interface{}{"type": "string", "description": "Color of the item"},
									"style_number": map[string]interface{}{"type": "string", "description": "Style number of the item"},
									"sizes": map[string]interface{}{
										"type":        "array",
										"description": "List of sizes and their quantities",
										"items": map[string]interface{}{
											"type": "object",
											"properties": map[string]interface{}{
												"size":     map[string]interface{}{"type": "string", "description": "Size of the item"},
												"quantity": map[string]interface{}{"type": "number", "description": "Quantity of the size"},
											},
											"required": []string{"size", "quantity"},
										},
									},
									"total_quantity": map[string]interface{}{"type": "number", "description": "Total quantity of all sizes combined"},
									"CTNS":           map[string]interface{}{"type": "number", "description": "Number of cartons for the order"},
								},
								"required": []string{"po", "color", "style_number", "sizes", "total_quantity", "CTNS"},
							},
						},
					},
					"required": []string{"orders"},
				},
			},
		},
		FunctionCall: openai.FunctionCall{
			Name: "export_simple_order_details",
		},
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("无法生成 ChatCompletion 响应:", err)
		return
	}

	if len(resp.Choices) == 0 {
		fmt.Println("响应中没有选择")
		return
	}

	// 提取工具调用参数
	var toolCall struct {
		Arguments string `json:"arguments"`
	}
	err = json.Unmarshal([]byte(resp.Choices[0].Message.FunctionCall.Arguments), &toolCall)
	if err != nil {
		fmt.Println("无法解析工具调用参数:", err)
		return
	}

	// 解析工具调用参数
	var ordersData map[string]interface{}
	err = json.Unmarshal([]byte(toolCall.Arguments), &ordersData)
	if err != nil {
		fmt.Println("无法解析工具调用参数:", err)
		return
	}

	// 将数据写入 JSON 文件
	ordersDataBytes, err := json.MarshalIndent(ordersData, "", "    ")
	if err != nil {
		fmt.Println("无法序列化订单数据:", err)
		return
	}

	err = ioutil.WriteFile(jsonOutputFile, ordersDataBytes, 0644)
	if err != nil {
		fmt.Println("无法写入 JSON 文件:", err)
		return
	}

	fmt.Println("数据已经成功保存到", jsonOutputFile, "文件中。")
}
