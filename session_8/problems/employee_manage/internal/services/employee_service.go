package services

import (
	"employee_manage/internal/models"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Employee_service struct {
	employee_lst []models.Employee
}

func (e *Employee_service) Load_employees(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	var employee_lst []models.Employee
	json.Unmarshal(data, &employee_lst)
	e.employee_lst = employee_lst

}

func (e *Employee_service) Increase_salary() {
	fmt.Println("increase salary called:")
	for ind, emp := range e.employee_lst {
		if emp.Age < 30 {
			fmt.Println(emp, "-30")
			e.employee_lst[ind].Salary = emp.Salary * (1.15)
			fmt.Println(emp.Salary)

		} else if emp.Age > 50 {
			fmt.Println(emp, "-50")
			e.employee_lst[ind].Salary = emp.Salary * (1.1)
			fmt.Println(emp.Salary)
		}
	}
	fmt.Println(e.employee_lst)

}

func (e *Employee_service) Sort_employee() {

	sort.Slice(e.employee_lst, func(i int, j int) bool {
		return e.employee_lst[i].Salary > e.employee_lst[j].Salary
	})
	fmt.Println(e.employee_lst)
}

func (e Employee_service) Filter_employee(department string) []models.Employee {
	var filterlst []models.Employee
	for _, emp := range e.employee_lst {
		if emp.Department == department {
			filterlst = append(filterlst, emp)
		}
	}

	fmt.Println(filterlst)
	return filterlst
}

func (e Employee_service) Save_data(filename string) {

	upd_data, err := json.Marshal(e.employee_lst)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile(filename, upd_data, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		fmt.Println("employee lst is not updated")
	}

}
func (e Employee_service) Generate_report() {
	dep_count := make(map[string]int)

	for _, emp := range e.employee_lst {
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
