package main

import (
	"fmt"

	"github.com/JngMkk/golang/variables/currency/currency"
)

func main() {
	userInput := "15.93"

	pennies, err := currency.ConvertStringDollarsToPennies(userInput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User input converted to %d pennies\n", pennies)

	pennies += 15

	dollars := currency.ConvertPenniestoDollarString(pennies)

	fmt.Printf("Added 15 cents, new values is %s dollars\n", dollars)
}
