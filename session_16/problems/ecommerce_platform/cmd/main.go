package main

import (
	"ecommerce_platform/internal/service"
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	inventory := make(map[string]int)
	inventory["Item 1"] = 53
	inventory["Item 2"] = 100
	inventory["Item 3"] = 30
	inventory["Item 4"] = 300
	ecom := service.NewEcommerce(inventory)
	itemlst := []string{"Item 1", "Item 2", "Item 3", "Item 4"}
	ordercount := 100
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	wg2.Add(3)

	go ecom.Order_Proccessing(&wg2)
	go ecom.Shipping(&wg2)
	go ecom.Delivery(&wg2)

	for i := 0; i < ordercount; i++ {
		wg1.Add(1)
		curitemmap := make(map[string]int)
		itemcount := rand.Intn(4)+1
		for j := 0; j < itemcount; j++ {
			curitemmap[itemlst[rand.Intn(4)]] = rand.Intn(50)
		}

		total := rand.Float64() * 100
		go ecom.PlaceOrder(string(i), curitemmap, total, "Puducherry", &wg1)

	}
	wg1.Wait()
	close(ecom.Orderqueue)
	wg2.Wait()
	for key,val:=range ecom.PlacedOrders{
		fmt.Println(key,"-",val)
	}
	fmt.Println(ecom.Inventory)

}
