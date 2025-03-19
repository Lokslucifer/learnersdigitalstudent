package main

import (
	"fmt"
	"time"
)

func foo() {
	fmt.Println("Hello")
	return
}

func main() {
	go foo()  //it is not printing because it is not waiting for child process to finish
	time.Sleep(1)
	fmt.Println("hi")
	fmt.Println("hi")

}