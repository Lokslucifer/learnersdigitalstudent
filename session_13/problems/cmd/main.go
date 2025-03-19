package main

import (
	"fmt"
	"sync"
)

func process(pid int, wg *sync.WaitGroup) {
	fmt.Println("Processing ",pid,":",pid * 2)
	wg.Done()
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	wg.Add(len(numbers))
	for _, num := range numbers {
		go process(num, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines completed")
}
