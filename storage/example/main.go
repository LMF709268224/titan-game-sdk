package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/zscboy/titan-game-sdk/storage"
)

func main() {
	locatorURL := flag.String("locator-url", "https://locator.titannet.io:5000/rpc/v0", "locator url")
	apiKey := flag.String("api-key", "", "api key")

	// 解析命令行参数
	flag.Parse()

	if len(*locatorURL) == 0 {
		fmt.Println("locator-url can not empty")
		return
	}

	if len(*apiKey) == 0 {
		fmt.Println("api-key can not empty")
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("please input file path")
		return
	}
	filePath := args[0]

	storage, close, err := storage.NewStorage(*locatorURL, *apiKey)
	if err != nil {
		fmt.Println("NewSchedulerAPI error ", err.Error())
		return
	}
	defer close()

	//show upload progress
	progress := func(doneSize int64, totalSize int64) {
		fmt.Printf("upload %d, total %d\n", doneSize, totalSize)
		if doneSize == totalSize {
			fmt.Printf("upload success\n")
		}
	}

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file error ", err.Error())
		return
	}

	// If want to upload file, use storage.UploadFile interface
	root, err := storage.UploadStream(context.Background(), f, f.Name(), progress)
	if err != nil {
		fmt.Println("UploadFile error ", err.Error())
		return
	}

	if err := storage.Delete(context.Background(), root.String()); err != nil {
		fmt.Println("UploadFile error ", err.Error())
		return
	}

	fmt.Printf("delete %s success\n", root.String())
}
