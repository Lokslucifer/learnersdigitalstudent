package main

import (
	"fmt"
	"log_analyzer/internal/service"
)

func ProccessLogFile() {
	filepath := ".\\sample.log"
	lineperchunk := 1
	loganalyser := service.NewLogAnalyzer()
	agglevelmap, err := loganalyser.ProcessFile(filepath, lineperchunk)
	if err != nil {
		fmt.Println(err)

	}
	for level, count := range agglevelmap {
		fmt.Println(level, "-", count)

	}

}

func main() {

	ProccessLogFile()

}
