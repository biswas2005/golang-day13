package practiceset2

import (
	"fmt"
	"strings"
)

type Processor interface {
	Process(data string) string
}

type UpperCaseProcessor struct{}

func (a UpperCaseProcessor) Process(data string) string {
	return strings.ToUpper(data)
}

type LowerCaseProcessor struct{}

func (b LowerCaseProcessor) Process(data string) string {
	return strings.ToLower(data)
}

func FileProcessor() {
	upper := UpperCaseProcessor{}
	lower := LowerCaseProcessor{}

	process := []Processor{upper, lower}

	data := "Hello, Welcome to Golang"

	for _, val := range process {
		result := val.Process(data)
		fmt.Printf("%T Processed Data:%s\n", val, result)
	}
}
