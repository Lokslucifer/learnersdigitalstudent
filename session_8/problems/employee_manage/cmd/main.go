package main

import (
	"employee_manage/internal/services"
)

func main() {
	emp_service := services.Employee_service{}
	emp_service.Load_employees("data.json")
	emp_service.Sort_employee()
	emp_service.Increase_salary()
	_ = emp_service.Filter_employee("Finance")
	emp_service.Save_data("upd.json")
	emp_service.Generate_report()
}
