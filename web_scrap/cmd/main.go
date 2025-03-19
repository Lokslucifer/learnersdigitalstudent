package main

import (
	"fmt"
	"runtime/debug"
	"sync"
	"web_scrap/internals/service"
)

func main() {
	defer debug.PrintStack() // <-- Prints full call stack
	var wg sync.WaitGroup
	baseurl := "https://usf-cs272-s25.github.io/top10/"
	webscrap := service.NewWebScrapper(baseurl)
	wg.Add(2)
	go webscrap.Crawl(&wg)
	go webscrap.Process(&wg)
	var query string
	var yes int
	for {
		fmt.Print("Do you want to search? (0 for No, 1 for Yes): ")
		fmt.Scanln(&yes) 
		if yes != 1 {
			break
		}

		fmt.Print("Enter search query: ")
		fmt.Scanln(&query) 
		fmt.Println("Search query:", query)
		webscrap.Search(query) 
	}
	wg.Wait()

}
