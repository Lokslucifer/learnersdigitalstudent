package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type ProcessedData struct {
	Data    string
	Count   int
	WordMap map[string]int
}

func NewProcessData(data string, count int, wordmap map[string]int) *ProcessedData {
	return &ProcessedData{Data: data, Count: count, WordMap: wordmap}
}

type WebScrapper struct {
	BaseUrl  string
	Urlqueue []string
	CrawlMap map[string]ProcessedData
	Rawdata  chan PageContent
	Url_List []string
	Visited  map[string]bool
}

func NewWebScrapper(baseurl string) *WebScrapper {
	return &WebScrapper{BaseUrl: baseurl, Urlqueue: []string{baseurl}, CrawlMap: make(map[string]ProcessedData), Rawdata: make(chan PageContent),
		Url_List: make([]string, 0), Visited: make(map[string]bool)}
}

func (ws *WebScrapper) Search(query string) {
	fmt.Println("search called")
	querylst := cleanWords(strings.Fields(query))
	ws.DisplayTopResults(querylst)
	fmt.Println("dISPAL CALLED")

}

func (ws *WebScrapper) Crawl(wg *sync.WaitGroup) {

	defer func() {
		// fmt.Println("closed")

		close(ws.Rawdata)
		fmt.Println("closed-rawdata")
		wg.Done()
	}()

	for len(ws.Urlqueue) != 0 {
		// fmt.Println(len(ws.Urlqueue), "-url quere")

		cururl := ws.Urlqueue[0]
		ws.Urlqueue = ws.Urlqueue[1:]

		cleanurl := removeDecimalFromLastNumber(cururl)
		_, exists := ws.Visited[cleanurl]
		// fmt.Println(cururl,"-",exists)

		if exists {
			continue
		}
		ws.Visited[cleanurl] = true
		// fmt.Println(cururl, "-url")
		resp, err := http.Get(cururl)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			continue
		}
		docnode, err := html.Parse(resp.Body)

		content := PageContent{}
		content.Baseurl = cururl
		if err != nil {
			bodyBytes, err_read := io.ReadAll(resp.Body)
			log.Println(err, "-during parsing-", docnode, "-", bodyBytes, "-", err_read)
			continue
		}
		resp.Body.Close()
		extractText(docnode, &content, cururl)
		// fmt.Println(content)

		ws.Urlqueue = append(ws.Urlqueue, content.Links...)
		// fmt.Println("appending")
		ws.Rawdata <- content

	}

}

func (ws *WebScrapper) Process(wg *sync.WaitGroup) {

	defer func() {
		fmt.Println("Finished processing")
		wg.Done()
	}()

	for content := range ws.Rawdata {

		wordmap := make(map[string]int)

		data := ""
		wordlst := make([]string, 0)

		for _, sen := range content.Paragraphs {
			words := cleanWords(strings.Fields(sen))
			wordlst = append(wordlst, words...)
			data = data + sen + "\n"
		}

		for _, word := range wordlst {
			_, exist := wordmap[word]
			if exist {
				wordmap[word] += 1
			} else {
				wordmap[word] = 1
			}
		}
		processdata := *NewProcessData(data, len(wordlst), wordmap)
		ws.CrawlMap[content.Baseurl] = processdata
		ws.Url_List = append(ws.Url_List, content.Baseurl)

	}
}

func getTopKKeys(m map[string]int, k int) []string {
	type kv struct {
		Key   string
		Value int
	}
	var kvSlice []kv
	for key, value := range m {
		kvSlice = append(kvSlice, kv{key, value})
	}

	sort.Slice(kvSlice, func(i, j int) bool {
		return kvSlice[i].Value > kvSlice[j].Value
	})

	topK := make([]string, 0, k)
	for i := 0; i < k && i < len(kvSlice); i++ {
		topK = append(topK, kvSlice[i].Key)
	}

	return topK
}

func (ws *WebScrapper) DisplayTopResults(querylst []string) {
	time.Sleep(10 * time.Second)
	urlst := ws.Url_List
	count_map := make(map[string]int)
	for _, url := range urlst {
		datamap := ws.CrawlMap[url].WordMap
		count := 0
		for _, q := range querylst {
			c, exist := datamap[q]
			if exist {
				count += c
			}

		}
		count_map[url] = count

	}
	topurls := getTopKKeys(count_map, 10)
	output := ""
	for ind, url := range topurls {

		if count_map[url] > 0 {

			cur := fmt.Sprintf("Rank:%d-Count:%d-Url-%s\n %s", ind+1, count_map[url], url, ws.CrawlMap[url].Data)

			output = output + cur
		}

	}

	os.WriteFile("output.txt", []byte(output), 0644)

}
