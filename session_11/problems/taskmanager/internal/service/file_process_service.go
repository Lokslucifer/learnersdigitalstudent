package service

import (
	"errors"
	"fmt"
	"os"
)

type FileProcessingTask struct {
	Operation string
	Filename  string
	Data      []byte
}

func NewFileProcessingTask(filename, oper string, data []byte) *FileProcessingTask {
	return &FileProcessingTask{Filename: filename, Operation: oper, Data: data}

}

func recoverFile(err *error) {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
		*err = errors.New(fmt.Sprint(r))
	}
}

func (f *FileProcessingTask) Run() (err error) {
	defer recoverFile(&err)
	if f.Operation == "read" {
		data, err := os.ReadFile(f.Filename)
		if err != nil {

			err = fmt.Errorf("read operation failed:%w", err)
			panic(err)
		}
		fmt.Println(data)

	} else if f.Operation == "write" {
		err := os.WriteFile(f.Filename, f.Data, 0644)
		if err != nil {
			fmt.Println("error writing to file")
			err = fmt.Errorf("write operation failed:%w", err)
			panic(err)
		}

	} else {
		baseErr := errors.New("invalid operation")
		err = fmt.Errorf("file processing failed:%w", baseErr)
		panic(err)
	}
	return nil

}
