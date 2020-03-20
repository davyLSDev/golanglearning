package main

import (
	"fmt"
)

type Car struct {
	make   string
	model  string
	engine string
}

/*
 displayCarFeatures() method has Car as the receiver type
*/
func (c Car) displayFeatures() {
	fmt.Printf("This %s is made by %s and has a %s engine. \n", c.model, c.make, c.engine)
}

func main() {
	car1 := Car{
		make:   "Toyota",
		model:  "Camry",
		engine: "3 litre V6",
	}
	car1.displayFeatures() //Calling displayFeatures() method of Car type
	car2 := Car{
		make:   "GMC",
		model:  "Safari",
		engine: "4.3 litre V8",
	}
	car2.displayFeatures()
}
