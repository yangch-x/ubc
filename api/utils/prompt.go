package utils

import (
	"fmt"
	"strings"
	"text/template"
)

type PromptTmpl struct {
	RewriteTarget      string
	UserIdentity       string
	ToneOption         string
	UserInputText      string
	EssayIdea          string
	UserBackgroundInfo string
	StudentBasicInfo   string
	UserEmail          string
}

func BuildPromptStr(content string, p PromptTmpl) (string, error) {
	// 解析模板
	t := template.Must(template.New("gptTmpl").Parse(content))

	var sb strings.Builder
	err := t.Execute(&sb, p)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return "", err
	}
	return sb.String(), nil
}
