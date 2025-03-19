package main

import (
	"fmt"
	sub "session_revision/add"
)

func main() {
	fmt.Println("Hello world")
	arr:=make([]int,0,10)
	// var arr []int
	fmt.Println((arr))

	c := sub.Sub(2, 4)
	fmt.Println((c))
	fmt.Println(cap(arr))
	fmt.Println(len(arr))

}
