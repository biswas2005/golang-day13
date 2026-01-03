package practiceset

import "fmt"

func Type(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("Type: int, value:", v)
	case string:
		fmt.Println("Type: string, value:", v)
	case bool:
		fmt.Println("Type: bool, value:", v)
	default:
		fmt.Println("Invalid type.")
	}
}

func TypeSwitch() {
	Type("Virat")
	Type(23)
	Type(true)
	Type(29)
	Type(23.4)
}
