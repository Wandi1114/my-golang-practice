package main

import (
	"first/calculation"
	"fmt"
)

func main() {
	num1, num2 := 8, 5
	action := "MULTIPLY"
	// action := "ADDITION";
	var result int
	switch action {
	case "MULTIPLY":
		fmt.Printf("Start Calculation of %s number %d and %d \n", action, num1, num2)

		result = calculation.Multiply(num1, num2)
		fmt.Printf("Multiply result : %d \n", result)
	case "ADDITION":
		fmt.Printf("Start Calculation of %s number %d and %d \n", action, num1, num2)
		result = calculation.Add(num1, num2)
		fmt.Printf("Addition result : %d \n", result)
	}

	if result > 100 {
		fmt.Println("the result is more than 100")
	} else if result >= 80 {
		fmt.Println("the result is more than 80")
	} else {
		fmt.Println("result is less than 80")
	}

	fmt.Printf("final result is : %s \n", result)

	for i := result; i > 0; i-- {
		fmt.Println("print result : ", i)
	}
}
