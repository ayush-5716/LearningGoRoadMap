package main

import (
	"errors"
	"fmt"
)

func main() {

	answer, newError := divide(4, 0)

	if newError == nil {
		fmt.Println(answer)
	} else {
		fmt.Println(newError)
	}

	newCar := car{
		Make:  "Hello",
		Model: "Bye",
		Dimension: dimension{
			Height: 234,
			Width:  834,
		},
	}

	msg := fmt.Sprintf(
		"The make of the car is %s, its model is %s and it of the height %d and width %d",
		newCar.Make, newCar.Model, newCar.Dimension.Height, newCar.Dimension.Width)

	fmt.Println(msg)
}

func divide(a, b int) (answer int, newError error) {
	if b == 0 {
		return 0, errors.New("Divisior cannot be zero")
	}

	return a / b, nil
}

type car struct {
	Make      string
	Model     string
	Dimension dimension
}

type dimension struct {
	Width  int
	Height int
}
