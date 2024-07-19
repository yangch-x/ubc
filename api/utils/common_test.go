package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay/apple"
	"io/ioutil"
	"strings"
	"testing"
	"text/template"
)

const (
	appKey = "8431ae7a52694675498d683b"
	secret = "04c64519929960ca38d592c3"
)

var (
	client *apple.Client
	err    error

	iss = "57246542-96fe-1a63-e053-0824d011072a"
	bid = "com.example.testbundleid2021"
	kid = "2X9R4HXF34"
)

func TestTransformFeedback(t *testing.T) {

	j := `
{
"Overall Score": "10",
"Expression": "1",
"Language": "1",
"Personality": "1",
"Reaction": 1,
"Expression advantages": "Not applicable as the interviewee did not provide a response.",
"Expression suggestions": "It's important to provide a response to all questions during an interview. Even if you need time to think, communicate this to the interviewer.",
"Language advantages": "Not applicable as the interviewee did not provide a response.",
"Language suggestions": "Remember to use proper grammar and vocabulary when responding to interview questions.",
"Personality advantages": "Not applicable as the interviewee did not provide a response.",
"Personality suggestions": "Showcase your personality during the interview by sharing experiences and insights.",
"Reaction advantages": "Not applicable as the interviewee did not provide a response.",
"Reaction suggestions": "Work on your ability to respond to unexpected questions or challenges in a timely manner."
}
`
	jj := `{"Overall Score": "0/100", "Expression": "0/10", "Language": "0/10", "Personality": "0/10", "Reaction": "0/10", "Expression advantages": "None", "Expression suggestions": "The interviewee needs to respond to the questions asked. Complete silence does not allow for any assessment of expression.", "Language advantages": "None", "Language suggestions": "The interviewee needs to respond to the questions to allow for an assessment of language skills.", "Personality advantages": "None", "Personality suggestions": "The interviewee needs to engage in the conversation to allow for an assessment of personality.", "Reaction advantages": "None", "Reaction suggestions": "The interviewee needs to respond to the questions to allow for an assessment of reaction skills."}`
	jjj := `{
  "Overall Score": 10,
  "Expression": 1,
  "Language": 1,
  "Personality": 4,
  "Reaction": 4,
  "Expression advantages": "N/A",
  "Expression suggestions": "It's important to respond to all questions during an interview. Silence can be interpreted as a lack of preparation or interest.",
  "Language advantages": "N/A",
  "Language suggestions": "There was no response to assess. It's crucial to communicate effectively during an interview.",
  "Personality advantages": "N/A",
  "Personality suggestions": "Engaging with the interviewer is a key part of showing your personality. Make sure to respond to all questions.",
  "Reaction advantages": "N/A",
  "Reaction suggestions": "Your reaction to the question was non-existent. It's important to at least acknowledge the question, even if you need time to think about your answer."
}`
	m := map[string]interface{}{}

	err := json.Unmarshal([]byte(j), &m)
	if err != nil {
		t.Error(err)
		return
	}

	mm := map[string]interface{}{}

	err = json.Unmarshal([]byte(jj), &mm)
	if err != nil {
		t.Error(err)
		return
	}

	mmm := map[string]interface{}{}

	err = json.Unmarshal([]byte(jjj), &mmm)
	if err != nil {
		t.Error(err)
		return
	}

	feedback := TransformFeedbacks2(m)
	t.Log(feedback)
	feedback = TransformFeedbacks2(mm)
	t.Log(feedback)
	feedback = TransformFeedbacks2(mmm)
	t.Log(feedback)
}

func TestPushMessage(t *testing.T) {
	_mains()
}

type DNSConfig struct {
	TTL          int
	NS           string
	AdminEmail   string
	SerialNumber int
	RefreshTime  string
	RetryTime    string
	ExpireTime   string
	MinimumTTL   string
	NSRecords    []Record
	ARecords     []Record
	NS1IP        string
}

type Record struct {
	Name  string
	Value string
}

func _mains() {
	// 读取文件内容
	fileContent, err := ioutil.ReadFile("D:\\goproject\\coredns\\db.tgstor.com")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 将文件内容转换为字符串
	fileStr := string(fileContent)

	// 解析文件内容并填充到结构体中
	dnsConfig := parseDNSConfig(fileStr)

	// 定义模板字符串
	tmplStr := `
$TTL {{.TTL}};
@               IN          SOA     ns1.{{.NS}}. {{.AdminEmail}} (
                            {{.SerialNumber}}   
                            {{.RefreshTime}}    
                            {{.RetryTime}}     
                            {{.ExpireTime}}     
                            {{.MinimumTTL}} )   

{{- range .NSRecords}}
{{.Name}}             IN          NS      {{.Value}}
{{- end}}

{{- range .ARecords}}
{{.Name}}             IN          A       {{.Value}}
{{- end}}

ns1             	  IN          A       {{.NS1IP}}
`

	// 解析模板
	tmpl := template.Must(template.New("dnsTemplate").Parse(tmplStr))

	// 使用数据填充模板
	var tplBuffer bytes.Buffer
	err = tmpl.Execute(&tplBuffer, dnsConfig)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	// 打印生成的 DNS 配置
	fmt.Println(tplBuffer.String())
}

func parseDNSConfig(fileContent string) DNSConfig {
	dnsConfig := DNSConfig{}

	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		switch fields[0] {
		case "$TTL":
			fmt.Sscanf(fields[1], "%d;", &dnsConfig.TTL)
		}
	}

	return dnsConfig
}
