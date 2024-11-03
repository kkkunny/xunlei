package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kkkunny/xunlei/dto"
	"github.com/kkkunny/xunlei/internal/api"
)

func main() {
	resp, err := api.CreateTask(context.Background(), "xxx", &api.CreateTaskRequest{
		PanAuth:  "xxx",
		Type:     string(dto.TaskTypeUserDownloadURL),
		Name:     "xxx",
		FileSize: 111,
		Space:    "xxx",
		Param: api.CreateTaskParam{
			Target: "xxx",
			URL:    "xxx",
		},
	})
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
