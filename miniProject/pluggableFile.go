package miniproject

import (
	"fmt"

	"os"
)

type Logger interface {
	Log(message string)
}
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

type FileLogger struct {
	FileName string
}

func (f FileLogger) Log(message string) {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("FileLogger error opening file", err)
		return
	}
	defer file.Close()

	_, err2 := file.WriteString(message + "\n")
	if err2 != nil {
		fmt.Println("FileLogger error writing to file", err2)
	}
}

func ProcessTask(task string, logger Logger) {
	result := "Processed Task:" + task
	logger.Log(result)
}

func PluggableFile() {
	consoleLogger := ConsoleLogger{}
	fileLogger := FileLogger{}

	ProcessTask("Send Email:", consoleLogger)
	ProcessTask("Generate report:", fileLogger)
	fmt.Println("Check system.Logs for File logs.")
}
