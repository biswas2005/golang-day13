package practiceset

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
	file, err := os.OpenFile(f.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
func ProcessData(data string, logger Logger) {
	result := "Processed:" + data
	logger.Log(result)
}
func LoggerSystem() {
	consoleLogger := ConsoleLogger{}
	ProcessData("Helllo World", consoleLogger)
	fileLogger := FileLogger{FileName: "app.log"}
	ProcessData("Hello File", fileLogger)
	fmt.Println("Check app.log for file logs.")
}
