package service

import "fmt"

type APICallTask struct {
	Call string
}

func NewAPICallTask(callname string) *APICallTask {
	return &APICallTask{Call: callname}
}

func (a *APICallTask) Run() error {
	fmt.Println("Calling ", a.Call)
	return nil
}
