package main

import (
	"first/calculation"
	"fmt"
)

func main() {
	fmt.Println("first print wandi")
	// testPrint := firstUtil()
	// fmt.Println(testPrint)

	checkAddition := calculation.Add(8, 2)

	fmt.Println(checkAddition)
}
