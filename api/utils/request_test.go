package utils

import (
	"testing"
)

func TestPost(t *testing.T) {
	url := "http://47.243.201.228:8000/feedback"
	b := `{"score":1,"run_id":"589b0fa2-5301-46d8-beea-ba7c511a8032","key":"user_score","feedback_id":"8124c8f4-d45042d9-8129-f7f54f86ab04","comment":"","source_info":{"is_explicit":true}}`

	post, err := Post(url, []byte(b))
	if err != nil {
		t.Errorf("post err:%v", err)
		return
	}

	t.Log(post)

}
