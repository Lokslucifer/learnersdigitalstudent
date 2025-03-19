package service

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

var log_store = make(map[string]int)
var mutex sync.Mutex

func LogGeneration(logs chan string, t_stamp time.Time, level string, msg string) {
	t := t_stamp.Format("2006-01-02 15:04:05")
	log := fmt.Sprintf("[%v] %s : %s", t, level, msg)
	logs <- log
	fmt.Println("send")

}

func ProcessLog(logs chan string, proccessed_logs chan LogEntry, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(proccessed_logs)

	}()
	for {
		log, ok := <-logs
		fmt.Println("received")
		if !ok {
			fmt.Println("Log channel closed, exiting process log goroutine.")
			return
		}

		fmt.Println("processing-", log)

		start := strings.Index(log, "[")
		end := strings.Index(log, "]")
		if start == -1 || end == -1 || start >= end {
			fmt.Println("Invalid log format")
			return
		}

		t_stamp := log[start+1 : end]

		parts := strings.SplitN(log[end+2:], " : ", 2)
		if len(parts) < 2 {
			fmt.Println("Invalid log format")
			return
		}

		level := parts[0]
		msg := parts[1]

		layout := "2006-01-02 15:04:05"
		parsedTime, err := time.Parse(layout, t_stamp)
		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
			return
		}

		logentry := LogEntry{Timestamp: parsedTime, Level: level, Message: msg}
		fmt.Println(logentry)
		fmt.Println(proccessed_logs)
		proccessed_logs <- logentry

		fmt.Println("process log send")
		// StoreProcessedLogs(proccessed_logs)

	}

}

func StoreProcessedLogs(processed_logs chan LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	for processed_log :=range processed_logs{



	fmt.Println("process recieved")
	mutex.Lock()
	_, found := log_store[processed_log.Level]
	if !found {
		log_store[processed_log.Level] = 1
	} else {
		log_store[processed_log.Level] += 1
	}
	mutex.Unlock()

	fmt.Println(log_store)}

}
func DisplayStorelog() {

	fmt.Println("Displaying log store")

	for key, value := range log_store {
		fmt.Println(key, "-", value)

	}
}
