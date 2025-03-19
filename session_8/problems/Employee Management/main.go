package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"os"
)

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

func increase_salary(employee_lst []Employee) {
	fmt.Println("increase salary called:")
	for ind, emp := range employee_lst {
		if emp.Age < 30 {
			fmt.Println(emp,"-30")
			employee_lst[ind].Salary = emp.Salary * (1.15)
			fmt.Println(emp.Salary)

		} else if emp.Age > 50 {
			fmt.Println(emp,"-50")
			employee_lst[ind].Salary = emp.Salary * (1.1)
			fmt.Println(emp.Salary)
		}
	}
	fmt.Println(employee_lst)

}

func sort_employee(employee_lst []Employee) {
	sort.Slice(employee_lst, func(i int, j int) bool {
		return employee_lst[i].Salary > employee_lst[j].Salary
	})
	fmt.Println(employee_lst)
}

func filter_employee(employee_lst []Employee, department string) []Employee {
	var filterlst []Employee
	for _, emp := range employee_lst {
		if emp.Department == department {
			filterlst = append(filterlst, emp)
		}
	}

	fmt.Println(filterlst)
	return filterlst
}
func main() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var employee_lst []Employee
	json.Unmarshal(data, &employee_lst)
	sort_employee(employee_lst)
	increase_salary(employee_lst)
	_ = filter_employee(employee_lst, "Finance")
	upd_data, err := json.Marshal(employee_lst)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile("updated_employee.json", upd_data, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		fmt.Println("employee lst is not updated")
	}
	dep_count := make(map[string]int)

	for _, emp := range employee_lst {
		_, found := dep_count[emp.Department]
		if found {
			dep_count[emp.Department] += 1
		} else {
			dep_count[emp.Department] = 1
		}
	}
	fmt.Println("Summary report:")
	fmt.Println("Total Employees by Department:")
	for dep, count := range dep_count {
		fmt.Println(dep, ":", count)
	}

}
