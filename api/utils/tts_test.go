package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"testing"
	"time"
)

func TestGetPollyToken(t *testing.T) {
	// 使用你的 AWS 访问密钥和密钥 ID
	accessKey := "AKIAX4YOFML7NMNCFCWA"
	secretKey := "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"
	region := "us-east-2"

	// 创建 AWS Session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		fmt.Println("Failed to create AWS session", err)
	}
	s := polly.New(sess)
	ssmlText := fmt.Sprintf("<speak><prosody rate='%s'>%s</prosody></speak>", "x-fast",
		"You are a professional interviewer. Based on the previous interview initial questions, you should present a different interview initial question, please note that the problem requires a different initial problem. so that the interviewee can have different interview experience each time. Directly give me the question without any other description.")

	params := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		Text:         aws.String(ssmlText),
		TextType:     aws.String("ssml"),
		LanguageCode: aws.String("en-US"),
		VoiceId:      aws.String("Emma"),
	}
	req, _ := s.SynthesizeSpeechRequest(params)
	url, err := req.Presign(10 * time.Minute)
	if err != nil {
		fmt.Println("Failed to sign request", err)
	}

	t.Log(url)
}
