package main

import (
	"fmt"
	"os"
)

type CustomError struct {
	Msg string
}

func (c *CustomError) Error() string {
	return c.Msg
}

func readfile(){
	data,err:=os.ReadFile("new.txt")
	if err!=nil{
		panic(err)
	}
	fmt.Println(data)
}
func main() {
	readfile()
	

}