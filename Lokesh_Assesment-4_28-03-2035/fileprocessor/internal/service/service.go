package service

import (
	"fileprocessor/internal/utils"
	"fmt"
	"log"
	"strings"
	"sync"
)

type ConcurrentFileProcessor struct {
	Filechannel     chan string
	Proccesschannel chan string
	Mode            string
	RetryCount      int
	Api             string
	Filterword      string
}

func NewConcurrentFileProcessor(mode string, filterword string, retrycount int, api string) *ConcurrentFileProcessor {
	return &ConcurrentFileProcessor{Filechannel: make(chan string), Proccesschannel: make(chan string), Mode: mode, RetryCount: retrycount, Api: api, Filterword: filterword}
}

func (fp *ConcurrentFileProcessor) TextFilesExtractor(dirpath string, wg *sync.WaitGroup) {
	defer func() {
		close(fp.Filechannel)

		wg.Done()
	}()

	filelst := utils.DirReader(dirpath)
	log.Println(filelst)
	for _, file := range filelst {
		if !file.IsDir() {
			fname := file.Name()
			splitname := strings.SplitN(fname, ".", 2)
			if splitname[1] == "txt" {
				filepath := dirpath + "/" + fname

				fp.Filechannel <- filepath
			}
		}
	}
}

func (fp *ConcurrentFileProcessor) CreateWorker(wg *sync.WaitGroup) {

	defer func() {

		close(fp.Proccesschannel)
		wg.Done()

	}()

	var wg2 sync.WaitGroup

	for filepath := range fp.Filechannel {

		wg2.Add(1)
		go fp.FileProcessor(filepath, &wg2)
	}
	wg2.Wait()

}

func (fp *ConcurrentFileProcessor) Start(dirpath string) {
	var wg sync.WaitGroup
	wg.Add(3)

	go fp.TextFilesExtractor(dirpath, &wg)
	go fp.CreateWorker(&wg)
	go fp.Aggregator(&wg)
	wg.Wait()

}

func (fp *ConcurrentFileProcessor) Aggregator(wg *sync.WaitGroup) {

	defer wg.Done()

	if fp.Mode == "Line Filter" {
		log.Println("Line called")
		fp.LinefilterAggregator()
	} else if fp.Mode == "Word Counter" {
		fp.WordCountAggregator()
	} else if fp.Mode == "API Call" {
		fp.APIcallAggregator(0)
	} else {
		fp.APIcallAggregator(fp.RetryCount)

	}
}

func (fp *ConcurrentFileProcessor) FileProcessor(filepath string, wg *sync.WaitGroup) {
	defer wg.Done()
	data := utils.FileReader(filepath)
	lines := strings.Split(data, "\n")
	for _, line := range lines {

		fp.Proccesschannel <- line
	}
}

func (fp *ConcurrentFileProcessor) LinefilterAggregator() {
	res := make([]string, 0)
	for line := range fp.Proccesschannel {

		val := utils.LineFilter(line, fp.Filterword)
		log.Println(line, "-", val)
		if val {
			res = append(res, line)

		}
	}
	fmt.Println(fp.Mode)
	for _, r := range res {
		fmt.Println(r)
	}
}

func (fp *ConcurrentFileProcessor) WordCountAggregator() {
	res := 0
	for line := range fp.Proccesschannel {
		wordcount := utils.WordCounter(line)
		res += wordcount
	}
	fmt.Println(fp.Mode)
	fmt.Println(res)
}

func (fp *ConcurrentFileProcessor) APIcallAggregator(trycount int) {

	res := make([]string, 0)
	for line := range fp.Proccesschannel {
		status := utils.APICaller(fp.Api, line, trycount)
		log.Println(status)
		res = append(res, status)
	}
	fmt.Println(fp.Mode)
	fmt.Println(res)
}
