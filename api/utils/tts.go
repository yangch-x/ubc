package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"time"
)

func PollyTextToSpeech(accessKey, secretKey, region string, params *polly.SynthesizeSpeechInput) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return "", err
	}
	s := polly.New(sess)
	req, _ := s.SynthesizeSpeechRequest(params)
	url, err := req.Presign(10 * time.Minute)
	if err != nil {
		return "", err
	}
	return url, nil
}
