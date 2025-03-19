package main
import "fmt"
type Human struct {
	name string
	age int 
	isPlaced bool
}

func (h Human)Sleep(){
	fmt.Println("Good night",h.name)
}

func main(){
	h:=Human{
		name:"Lokesh",
		age:21,
		isPlaced: true,
	}
	hpointer:=&h
	fmt.Println(h,&hpointer)
	newhpointer:=new(Human)
	fmt.Println(newhpointer)
	hpointer.Sleep()
}