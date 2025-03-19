package main

import "fmt"

func problem1() {
	
	arr := []int{1, 2, 3, 4}

	s1 := arr[:2]

	s2 := arr[1:3]

	s1 = append(s1, 10) // Modify slice s1

	fmt.Println(arr, s1, s2)

}

//output of problem will be [1,2,10,4],[1,2,10] ,[2,10] because creating slice dont create copy of array
//it create a pointer which allows us to access and assign elements in the array
func problem2() {

	arr := []int{1, 2, 3, 4, 5}

	s := arr[1:3]

	fmt.Println(len(s), cap(s)) // Length and capacity before appending

	s = append(s, 10, 20, 30)

	fmt.Println(arr, s)
	/*expected :[1,2,3,10,20] [2,3,10,20,30]
	//but got :[1,2,3,4,5] [2,3,10,20,30]
	//this happens because when you tried to add element
	//beyond capacity it create a next slice in dif location
	*/

}
func modifySlice(s []int) {

	s[0] = 100

}
func problem3() {

	arr := []int{1, 2, 3, 4, 5}

	s := arr[:3]

	modifySlice(s)

	fmt.Println(arr)
	//expected [100,2,3,4,5]
	//got [100,2,3,4,5]
	/* Slice doesnot create a array copy untill it exceeds
	the capacity size,it will point to same array,but new slice will have
	dif len and capacity*/

}

func problem4() {

	s1 := []int{1, 2, 3, 4}

	s2 := make([]int, len(s1))

	copy(s2, s1)

	s1[0] = 100

	fmt.Println(s1, s2)

	s1 = append(s1, 200)

	fmt.Println(s1, s2)
	//expected [100,2,3,4] []
	//got [100,2,3,4,5]
	/* Slice doesnot create a array copy untill it exceeds
	the capacity size,it will point to same array,but new slice will have
	dif len and capacity*/

}

func problem5() {

	s := make([]int, 2, 3)

	s[0], s[1] = 10, 20

	s1 := append(s, 30)

	s2 := append(s, 40)

	fmt.Println(s, s1, s2)
	//expected [10,20] [10,20,30] [10,20,40]
	//got [10,20] [10,20,40] [10,20,40] because
	/* Slice doesnot create a array copy untill it exceeds
	the capacity size,it will point to same array,but new slice will have
	dif len and capacity*/

}
func main() {
	
	problem3()

}
