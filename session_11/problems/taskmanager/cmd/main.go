package main

import (
	"fmt"
	"os"
	"taskmanager/internal/service"
)

func main() {
	t1 := service.NewDataValidationTask("lok8695esh@gmail.com", `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	t2 := service.NewFileProcessingTask("new.txt", "read", make([]byte, 0))
	t3 := service.NewAPICallTask("interupt")
	fmt.Println(os.ErrNotExist)

	taskmanger := service.NewTaskManager(3)
	taskmanger.AddTask(t1)
	taskmanger.AddTask(t2)
	taskmanger.AddTask(t3)
	taskmanger.Execute()

}
