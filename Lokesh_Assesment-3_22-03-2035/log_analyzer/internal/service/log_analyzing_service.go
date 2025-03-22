package service

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

type LogEntry struct {
	Level   string
	Message string
}

func NewLogEntry(level string, msg string) LogEntry {
	return LogEntry{Level: level, Message: msg}
}

type LogAnalyzer struct {
	AggregateLevelMap map[string]int
	LogMutex          sync.Mutex
	LevelMapChan      chan map[string]int
	Logs              []LogEntry
	invaliderror      error
}

func NewLogAnalyzer() *LogAnalyzer {
	return &LogAnalyzer{AggregateLevelMap: make(map[string]int), LogMutex: sync.Mutex{}, LevelMapChan: make(chan map[string]int, 5), invaliderror: errors.New("invalid log")}
}

func (l *LogAnalyzer) ProcessFile(filepath string, lineperchunk int) (map[string]int, error) {

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return make(map[string]int), err
	}

	strdata := string(data)
	splitdata := strings.Split(strdata, "\n")
	n := len(splitdata)
	end := 0
	wg2.Add(1)

	go l.AggregateLevel(&wg2)

	for i := 0; i < n-1; i = i + lineperchunk {
		if i+lineperchunk >= n {
			end = n
		} else {
			end = i + lineperchunk
		}

		logs := splitdata[i:end]
		wg1.Add(1)
		go l.ProcessLogs(logs, &wg1)
	}

	wg1.Wait()
	close(l.LevelMapChan)
	wg2.Wait()

	return l.AggregateLevelMap, nil

}

func (l *LogAnalyzer) ProcessLogs(logs []string, wg *sync.WaitGroup) {

	levelmap := make(map[string]int)
	defer func() {
		l.LevelMapChan <- levelmap
		wg.Done()
	}()

	for _, log := range logs {

		splitlog := strings.SplitN(log, " ", 2)

		if len(splitlog) < 2 {

			fmt.Println(log, "-", l.invaliderror)
			continue

		}
		level := splitlog[0]
		msg := splitlog[1]
		newlog := NewLogEntry(level, msg)

		l.LogMutex.Lock()
		l.Logs = append(l.Logs, newlog)
		l.LogMutex.Unlock()

		levelmap[level] += 1

	}

}
func (l *LogAnalyzer) AggregateLevel(wg *sync.WaitGroup) {

	defer wg.Done()

	for levelmap := range l.LevelMapChan {
		for level, count := range levelmap {
			_, exist := l.AggregateLevelMap[level]
			if exist {
				l.AggregateLevelMap[level] += count
			} else {
				l.AggregateLevelMap[level] = count
			}
		}

	}
}
