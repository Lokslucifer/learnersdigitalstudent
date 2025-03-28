package main

import "fileprocessor/internal/service"

const (
	dirpath    = "./test_files/"
	mode       = "Word Counter"
	filterword = "fine"
	api="httpbin.org/post"
	trycount=3
)

func Process() {
	fp := service.NewConcurrentFileProcessor(mode, filterword, trycount,api )
	fp.Start(dirpath)

}
func main() {
	Process()

}
