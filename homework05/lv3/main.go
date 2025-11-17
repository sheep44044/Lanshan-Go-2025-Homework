package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("用法: %s [目录路径] [搜索关键词]\n", os.Args[0])
		os.Exit(1)
	}

	dirPath := os.Args[1]
	keyword := os.Args[2]

	files, err := getAllFiles(dirPath)
	if err != nil {
		fmt.Printf("错误: 无法读取目录 %s: %v\n", dirPath, err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Printf("在目录 %s 中没有找到文件\n", dirPath)
		return
	}

	fmt.Printf("在目录 %s 中找到 %d 个文件，开始搜索关键词: %s\n\n", dirPath, len(files), keyword)

	pool := NewWorkerPool(8)

	var results []SearchResult
	var resultWg sync.WaitGroup
	resultWg.Add(1)

	go func() {
		defer resultWg.Done()
		for result := range pool.Results() {
			results = append(results, result)
		}
	}()

	for _, file := range files {
		task := Task{
			FilePath: file,
			Keyword:  keyword,
		}
		pool.Submit(task)
	}

	pool.WaitAndClose()

	resultWg.Wait()

	if len(results) == 0 {
		fmt.Printf("没有找到包含关键词 '%s' 的内容\n", keyword)
		return
	}

	fmt.Printf("找到 %d 个匹配结果:\n\n", len(results))
	for _, result := range results {
		fmt.Printf("文件: %s\n", result.FilePath)
		fmt.Printf("行号: %d\n", result.LineNum)
		fmt.Printf("内容: %s\n", result.Content)
		fmt.Println("---")
	}
}
