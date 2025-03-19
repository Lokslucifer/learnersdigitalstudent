package models
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
	Address    Address `json:"address"`
}