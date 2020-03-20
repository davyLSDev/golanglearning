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

// must use a pointer or else changes made inside this method will then be visible to the caller
func (c *Car) addFeatures(make, model, engine string) {
	c.make = make
	c.model = model
	c.engine = engine
}

func (c Car) addFeaturesByValue(make, model, engine string) {
	c.make = make
	c.model = model
	c.engine = engine
	fmt.Println("DEBUG: ", c.make, c.model, c.engine)
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

	fmt.Println("First use a procedure with a value as the sender")
	car3 := Car{}
	car3.addFeaturesByValue("Mazda", "Protege", "2 litre v4")
	car3.displayFeatures()

	fmt.Println("Now compare with the procedure which uses a pointer to the value as the sender")

	car3.addFeatures("Mazda", "Protege", "2 litre v4")
	car3.displayFeatures()
}
