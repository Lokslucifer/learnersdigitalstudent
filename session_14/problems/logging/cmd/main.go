package main

import (
	"fmt"
	"logging/internals/service"
	"math/rand"
	"sync"
	"time"
)

func simulate_rand_process(id int, logs chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	sleepTime := time.Duration(rand.Intn(5)+1) * time.Second
	fmt.Printf("Worker %d sleeping for %v...\n", id, sleepTime)
	time.Sleep(sleepTime)

	levels := []string{"INFO", "ERROR", "WARN"}
	level := levels[rand.Intn(3)]
	var msg string
	if level == "INFO" {
		msg = "Completed successfully"

	} else if level == "ERROR" {
		msg = "Error has occured"
	} else {
		msg = "Unexpected warning"
	}
	currentTime := time.Now()
	fmt.Println("generating log-", id)
	service.LogGeneration(logs, currentTime, level, msg)
	fmt.Printf("Worker %d finished work\n", id)

}
func main() {
	logs := make(chan string,1)
	processed_logs := make(chan service.LogEntry,1)

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	wg1.Add(3)
	wg2.Add(2)
	go simulate_rand_process(0, logs, &wg1)
	go simulate_rand_process(1, logs, &wg1)
	go simulate_rand_process(2, logs, &wg1)
	go service.ProcessLog(logs, processed_logs, &wg2)
	go service.StoreProcessedLogs(processed_logs,&wg2)

	wg1.Wait()

	close(logs)
	fmt.Println("log closed")

	wg2.Wait()
	service.DisplayStorelog()

}
