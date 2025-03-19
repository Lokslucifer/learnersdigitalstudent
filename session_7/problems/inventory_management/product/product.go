package product

import "fmt"

type Product struct {
	id       int
	name     string
	price    int
	isavail  bool
	quantity int
	des      string
	category string
}

func Product_Constructor(id int, name string, price int, quantity int, des string, category string) *Product {
	p := Product{
		id:       id,
		name:     name,
		price:    price,
		quantity: quantity,
		category: category,
		des:      des,
	}
	if p.quantity > 0 {
		p.isavail = true
	}
	return &p

}
func (p *Product) Update_price(newprice int) {
	print("product updated")
	p.price = newprice
}

func (p *Product) Update_quantity(newquantity int) {
	p.quantity = newquantity
	if p.quantity > 0 {
		p.isavail = true
	}
}
func (p Product) Check_name(name string) bool {
	return p.name == name
}

func (p Product) Check_category(cate string) bool {
	return p.category == cate
}

func (p Product) Display() {

	fmt.Println("product id:", p.id)
	fmt.Println("Product name:", p.name)
	fmt.Println("Product price:", p.price)
	fmt.Println("Product description:", p.des)
	fmt.Println("Product Quantity avail:", p.quantity)
	fmt.Println("Product Category:", p.category)

}
