package main

import (
	"first/calculation"
	"fmt"
)

func main() {
	num1, num2 := 8, 5

	fmt.Printf("Start Calculation of number %d and %d \n", num1, num2)
	checkAddition := calculation.Add(num1, num2)
	fmt.Printf("Addition result : %d \n", checkAddition)

	checkMultiply := calculation.Multiply(num1, num2)
	fmt.Printf("Multiply result : %d \n", checkMultiply)
}
