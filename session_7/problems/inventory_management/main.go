package main

import (
	"inventory_management/inventory"
)

func main() {
	inven := inventory.Inventory_constructor()
	inven.AddProduct(1, "vivo y73", 12000, 5, "android mobile phone", "Electronics")
	inven.ShowInventory()
	inven.AddProduct(2, "vivo y5", 59000, 2, "android mobile phone", "Electronics")
	inven.ShowInventory()
	inven.AddProduct(3, "bajaj iron box", 1000, 10, "iron box use for ironing clothes", "Appliances")
	inven.ShowInventory()
	inven.DeleteProduct(1)
	inven.ShowInventory()
	inven.UpdateProduct_Price(2, 30000)
	inven.ShowInventory()
	inven.UpdateProduct_Quantity(2, 300)
	inven.ShowInventory()
	_ = inven.SearchByname("vivo y5")
	_ = inven.SearchBycategory("Appliances")

}
