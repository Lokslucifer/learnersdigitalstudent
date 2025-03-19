package service

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	Id        string
	UserId    string
	Itemcount map[string]int
	Total     float64
	Address   string
	Status    string
}

func NewOrder(id, user_id string, item_map map[string]int, total float64, address string) Order {

	return Order{Id: id, UserId: user_id, Itemcount: item_map, Total: total, Address: address}
}

type Ecommerce struct {
	Orderqueue         chan *Order
	ShippingOrderqueue chan *Order
	DeliveryOrderqueue chan *Order
	Inventory          map[string]int
	PlacedOrders       map[string]*Order
	UniqueId           map[string]bool
}

var InvenMutex sync.Mutex
var IdMutex sync.Mutex

func NewEcommerce(inventory map[string]int) *Ecommerce {
	return &Ecommerce{Orderqueue: make(chan *Order), ShippingOrderqueue: make(chan *Order), DeliveryOrderqueue: make(chan *Order), Inventory: inventory,
		PlacedOrders: make(map[string]*Order), UniqueId: make(map[string]bool)}

}

func (e *Ecommerce) generateUniqueID() string {
	rand.Seed(time.Now().UnixNano()) // Seed random generator
	new_id := fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(10000))
	// _, exist := e.UniqueId[new_id]
	// for exist {
	// 	new_id = fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(10000))
	// 	_, exist = e.UniqueId[new_id]

	// }
	return new_id
}

func (e *Ecommerce) PlaceOrder(user_id string, item_map map[string]int, total float64, address string, wg *sync.WaitGroup) {
	defer wg.Done()

	IdMutex.Lock()
	new_id := e.generateUniqueID()
	_, exist := e.UniqueId[new_id]
	for exist {
		new_id := e.generateUniqueID()
		_, exist = e.UniqueId[new_id]
	}
	e.UniqueId[new_id] = true
	// fmt.Println(new_id)
	IdMutex.Unlock()
	new_order := NewOrder(new_id, user_id, item_map, total, address)
	new_order.Status = "Processing"
	e.PlacedOrders[new_order.Id] = &new_order

	e.Orderqueue <- &new_order

}

func (e *Ecommerce) Order_Proccessing(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(e.ShippingOrderqueue)
	}()
	for order := range e.Orderqueue {
		InvenMutex.Lock()
		placed := true
		// fmt.Println("Inventory before-",e.Inventory)
		// fmt.Println("Order-",order)
		for item, count := range order.Itemcount {
			// fmt.Println(item,"-",count,"-",e.Inventory[item])
			if e.Inventory[item] < count {

				placed = false
				break
			}
		}
		// fmt.Println(placed)
		if placed {
			for item, count := range order.Itemcount {
				if e.Inventory[item] >= count {
					e.Inventory[item] -= count
				} else {
					log.Fatal("Race condition")
				}
			}
			order.Status = "Processing"

		} else {
			order.Status = "Rejected"
		}
		// fmt.Println(placed)

		// fmt.Println(order)
		// fmt.Println(e.Inventory)

		InvenMutex.Unlock()

		if order.Status != "Rejected" {

			e.ShippingOrderqueue <- order
		}

	}

}

func (e *Ecommerce) Shipping(wg *sync.WaitGroup) {

	defer func() {
		wg.Done()
		close(e.DeliveryOrderqueue)
	}()
	for order := range e.ShippingOrderqueue {
		log.Println(order.Id, "-shipped successfully")
		order.Status = "Shipped"
		e.DeliveryOrderqueue <- order

	}

}

func (e *Ecommerce) Delivery(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	for order := range e.DeliveryOrderqueue {
		log.Println(order.Id, "with items ", order.Itemcount, "-delivery successfully-", order.UserId)
		order.Status = "Completed"

	}

}
