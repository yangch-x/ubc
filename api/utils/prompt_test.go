package utils

import (
	"testing"
)

func TestTemplate(t *testing.T) {

	// Rewrite this text with {{.rewriteTarget}} for a {{.userIdentity}}  in a {{if .toneOption}}in a {{.toneOption}}{{else}}{{end}} : {{.userInputText}}
	// Generate an essay outline for a <user_identity> in a <tone_option> tone based on this theme and content: <user_input_text>
	// Write an essay for a <user_identity> in a <tone_option> tone based on this outline: <user_input_text>
	// Rewrite this email for a <user_identity> in a <tone_option> tone with this background information: <user_input_text>
	// Rewrite this email for a <user_identity> in a <tone_option> tone considering this user background information: <user_background_info>. Here is the original email: <user_email>
	// Brainstorm an essay idea for a <user_identity> school application in a <tone_option> tone based on this basic information: <student_basic_info>
	// Write an essay for a <user_identity> school application in a <tone_option> tone based on this essay idea: <essay_idea>

	// 定义模板字符串
	//const tmpl = `Rewrite this text with {{.RewriteTarget}} for a {{.UserIdentity}} in a {{if .ToneOption}}{{.ToneOption}}{{else}}{{end}}: {{.UserInputText}}`
	//const tmpl = `Generate an essay outline for a {{.UserIdentity}} in a {{if .ToneOption}}{{.ToneOption}}{{else}}{{end}} tone based on this theme and content: {{.UserInputText}}`
	//const tmpl = `Write an essay for a {{.UserIdentity}} in a {{if .ToneOption}}{{.ToneOption}}{{else}}{{end}} tone based on this outline: {{.UserInputText}}`
	//const tmpl = `Rewrite this email for a {{.UserIdentity}} in a {{if .ToneOption}}{{.ToneOption}}{{else}}{{end}} tone with this background information: {{.UserInputText}}`
	//const tmpl = `Rewrite this email for a {{.UserIdentity}} in a {{if .ToneOption}}{{.ToneOption}}{{else}}{{end}} tone considering this user background information: {{.UserBackgroundInfo}}. Here is the original email: {{.UserEmail}}`
	//const tmpl = `Brainstorm an essay idea for a {{.UserIdentity}} school application in a {{if .ToneOption}}{{.ToneOption}}{{else}}{{end}} tone based on this basic information: {{.StudentBasicInfo}}`
	const tmpl = `Write an essay for a {{.UserIdentity}} school application in a {{if .ToneOption}}{{.ToneOption}}{{else}}{{end}} tone based on this essay idea: {{.EssayIdea}}`
	p := PromptTmpl{
		RewriteTarget: "rewriteTarget",
		UserIdentity:  "userIdentity",
		ToneOption:    "toneOption",
		EssayIdea:     "essayIdea",
	}
	str, err := BuildPromptStr(tmpl, p)
	if err != nil {
		t.Errorf("Error executing template:%s", err)
	}
	if str != "Write an essay for a userIdentity school application in a toneOption tone based on this essay idea: essayIdea" {
		t.Errorf("Error executing template:%s", err)
	}
}
