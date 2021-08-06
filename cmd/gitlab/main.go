package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/xanzy/go-gitlab"
)

func main() {
	client, _ := gitlab.NewClient(
		os.Getenv("GITLAB_TOKEN"),
		gitlab.WithBaseURL(os.Getenv("GITLAB_ENDPOINT")+"/api/v4"),
	)
	title := "hello"
	source := "dev"
	target := "master"
	mr, _, _ := client.MergeRequests.CreateMergeRequest(1107, &gitlab.CreateMergeRequestOptions{
		Title:        &title,
		SourceBranch: &source,
		TargetBranch: &target,
	})
	bs, _ := json.Marshal(mr)
	fmt.Println(string(bs))
	iid := mr.IID
	fmt.Println(mr.State)
	for {
		mr, _, _ := client.MergeRequests.GetMergeRequest(1107, iid, nil)
		fmt.Println(mr.State)
		time.Sleep(time.Second)
		bs, _ := json.Marshal(mr)
		fmt.Println(string(bs))
	}
}
