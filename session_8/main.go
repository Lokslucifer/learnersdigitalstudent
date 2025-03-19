package main

import (
	"encoding/json"
	"fmt"
)



type Data struct {
	Name string `json:"name`
	Age  int `json:"age"`
}

func (d Data) DoubleAge() int {
	return 2 * d.Age
}
func main() {
	data := Data{
		Name: "John",
		Age:  23,
	}
	jsonData,err := json.Marshal(data)
	if  err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
	// fmt.Println(jsonData)
	data.Age = data.DoubleAge()
	fmt.Println(data.Age)

}