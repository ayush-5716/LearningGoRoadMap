package main

import "fmt"

func main() {

	goat1 := goat{
		animal: animal{
			Walks: true,
			Talks: false,
			Baths: false,
		},
		Colour: "Maroorn",
	}

	goat2 := goat{
		animal: animal{
			Walks: true,
			Talks: true,
			Baths: false,
		},
		Colour: "Maroorn",
	}

	if goat1 == goat2 {
		fmt.Println("They are same ")
	} else {
		fmt.Println("They are different ")
	}
}

type goat struct {
	animal
	Colour string
}

type animal struct {
	Walks bool
	Talks bool
	Baths bool
}
