package inventory

import (
	"fmt"
	"inventory_management/product"
)

type Inventory struct {
	inven_map map[int]*product.Product
}

func Inventory_constructor() Inventory {
	return Inventory{
		inven_map: make(map[int]*product.Product),
	}

}

func (inven *Inventory) AddProduct(id int, name string, price int, quantity int, des string, category string) {
	_, found := inven.inven_map[id]
	if found {
		fmt.Println("Product id already exists")
		return
	}
	newproduct := product.Product_Constructor(id, name, price, quantity, des, category)
	inven.inven_map[id] = newproduct
	fmt.Println("Product successfully added")

}

func (inven *Inventory) DeleteProduct(id int) {
	_, found := inven.inven_map[id]
	if !found {
		fmt.Println("Product id not exists")
		return
	}
	delete(inven.inven_map, id)

	fmt.Println("Product successfully deleted")

}

func (inven *Inventory) UpdateProduct_Price(id int, price int) {
	curproduct, found := inven.inven_map[id]
	if !found {
		fmt.Println("Product not exists")
		return
	}
	curproduct.Update_price(price)
	fmt.Println("Product price updated successfully")

}

func (inven *Inventory) UpdateProduct_Quantity(id int, quantity int) {
	curproduct, found := inven.inven_map[id]
	if !found {
		fmt.Println("Product not exists")
		return
	}
	curproduct.Update_quantity(quantity)
	fmt.Println("Product quantity updated successfully")

}

func (inven *Inventory) ShowInventory() {
	for id, curproduct := range inven.inven_map {
		fmt.Println("ID-", id)
		curproduct.Display()
		fmt.Println()
	}
}

func (inven Inventory) SearchByname(name string) []product.Product {
	fmt.Println("searching product with name:", name)
	temp := make([]product.Product, 0)
	for _, curproduct := range inven.inven_map {
		if curproduct.Check_name(name) {
			curproduct.Display()
			temp = append(temp, *curproduct)
		}

	}
	return temp

}

func (inven Inventory) SearchBycategory(category string) []product.Product {
	fmt.Println("searching product with category:", category)
	temp := make([]product.Product, 0)
	for _, curproduct := range inven.inven_map {
		if curproduct.Check_category(category) {
			curproduct.Display()
			temp = append(temp, *curproduct)

		}

	}
	return temp
}
