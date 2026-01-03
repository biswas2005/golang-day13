package practiceset

import "fmt"

func TypeAssertion() {
	var i interface{}

	i = 10
	if val, s := i.(int); s {
		fmt.Println("Int value:", val)
	} else {
		fmt.Println("Not an Int")
	}

	i = "Virat"
	if val, ok := i.(string); ok {
		fmt.Println("String value:", val)
	} else {
		fmt.Println("Not a String")
	}

	i = true
	if val, ok := i.(bool); ok {
		fmt.Println("Bool value:", val)
	} else {
		fmt.Println("Not a Bool")
	}

}