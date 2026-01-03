package practiceset

import "fmt"

type Vehicle interface {
	speed() int
}

type Car struct {
	Model string
}
type Bike struct {
	Brand string
}

func (c Car) speed() int {
	return 120
}

func (c Bike) speed() int {
	return 90
}

func VehicleSpeed() {
	car := Car{Model: "Maruti"}
	bike := Bike{Brand: "Yamaha"}

	vehicle := []Vehicle{car, bike}

	for _, val := range vehicle {
		fmt.Printf("%T speed:%d\n", val, val.speed())
	}
}
