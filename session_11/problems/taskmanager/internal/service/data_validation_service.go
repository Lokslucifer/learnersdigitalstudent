package service

import (
	"errors"
	"fmt"
	"regexp"
)

type DataValidationTask struct {
	Data    string
	Pattern string
}

func NewDataValidationTask(data, pattern string) *DataValidationTask {
	return &DataValidationTask{Data: data, Pattern: pattern}
}

func (d *DataValidationTask) Run() error {
	re, err := regexp.Compile(d.Pattern)
	if err != nil {
		fmt.Println("Invalid regex pattern:", err)
		return fmt.Errorf("validation Error:%w", err)
	}
	res := re.MatchString(d.Data)
	if res {
		return nil
	}
	baseErr := errors.New("not valid")
	return fmt.Errorf("validation Error:%w", baseErr)

}
