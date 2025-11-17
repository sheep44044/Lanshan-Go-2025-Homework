package main

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type SearchResult struct {
	FilePath string
	LineNum  int
	Content  string
}

type Task struct {
	FilePath string
	Keyword  string
}

type WorkerPool struct {
	taskChan   chan Task
	resultChan chan SearchResult
	wg         sync.WaitGroup
}

func getAllFiles(dirPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

func NewWorkerPool(workerCount int) *WorkerPool {
	pool := &WorkerPool{
		taskChan:   make(chan Task, 100),
		resultChan: make(chan SearchResult, 100),
	}

	for i := 0; i < workerCount; i++ {
		pool.wg.Add(1)
		go pool.worker()
	}

	return pool
}

func (p *WorkerPool) worker() {
	defer p.wg.Done()

	for task := range p.taskChan {
		p.searchInFile(task.FilePath, task.Keyword)
	}
}

func (p *WorkerPool) searchInFile(filePath, keyword string) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	lineNum := 1

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if strings.Contains(line, keyword) {
			cleanedLine := strings.TrimSpace(line)
			result := SearchResult{
				FilePath: filePath,
				LineNum:  lineNum,
				Content:  cleanedLine,
			}
			p.resultChan <- result
		}

		if err == io.EOF {
			break
		}
		lineNum++
	}
}

func (p *WorkerPool) Submit(task Task) {
	p.taskChan <- task
}

func (p *WorkerPool) WaitAndClose() {
	close(p.taskChan)
	p.wg.Wait()
	close(p.resultChan)
}

func (p *WorkerPool) Results() <-chan SearchResult {
	return p.resultChan
}
