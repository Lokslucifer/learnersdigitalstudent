package utils

import (

	"log"
	"net/http"
	"os"
	"strings"
)

func DirReader(dirpath string) []os.DirEntry {
	filelst, err := os.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
		return []os.DirEntry{}
	}
	return filelst

}

func FileReader(filepath string) string {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(data)
}
func LineFilter(line string, filterword string) bool {
	wordlst := strings.Split(line, " ")
	for _, word := range wordlst {
		if word == filterword {
			return true
		}
	}
	return false

}

func WordCounter(line string) int {
	wordlst := strings.Split(line, " ")
	return len(wordlst)
}

func APICaller(line string, api string, trycount int) string {

	
	

	for trycount >= 0 {

		resp, err := http.Post(api, "txt", strings.NewReader(line))
		if err != nil {
			log.Println(err)
			trycount--
			continue
		}
		return resp.Status
	}
	return ""

}
