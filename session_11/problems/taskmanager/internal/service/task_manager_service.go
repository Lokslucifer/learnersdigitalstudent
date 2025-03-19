package service

import (
	"errors"
	"log"
	"os"
)

type TaskManager struct {
	Tasklist       []Task
	FailedTasklist []Task
	Maxtry         int
}

func NewTaskManager(maxtry int) *TaskManager {
	return &TaskManager{Tasklist: make([]Task, 0), FailedTasklist: make([]Task, 0), Maxtry: maxtry}

}

func (t *TaskManager) AddTask(task Task) error {
	t.Tasklist = append(t.Tasklist, task)
	return nil

}

func (t *TaskManager) Execute() {
	for _, task := range t.Tasklist {
		count := 0
		success := false
		for count < t.Maxtry {
			err := task.Run()
			
			if errors.Is(err, os.ErrNotExist) {
				break
			} else if errors.Is(err, errors.New("not valid")) {
				break
			}
			if err != nil {
				count += 1
			} else {
				success = true
				break
			}
		}
		if !success {
			t.FailedTasklist = append(t.FailedTasklist, task)
		}
		PrintLog(task, count+1, success)

	}

}

func PrintLog(task Task, try int, success bool) {
	switch t := task.(type) {
	case *DataValidationTask:

		log.Println(t)
	case *FileProcessingTask:

		log.Println(t)
	case *APICallTask:

		log.Println(t)
	default:
		log.Println("Unknown tasl")

	}

	log.Println("Tried:", try)
	if success {
		log.Println("Success")
	} else {
		log.Println("Failed")
	}

}
