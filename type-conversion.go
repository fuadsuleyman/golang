package main

import "fmt"

func main () {
	numInt := 200
	numFloat := 2.5

	// below print 200.0000
	fmt.Printf("decimal numInt: %.4f \n", float64(numInt))

	// below print 2
	fmt.Printf("decimal numInt: %d \n", int(numFloat))

	// for get right answer  first convert int to float
	// then mutliple
	// then convert multiple result to int
	// in GO we can't multiple numbers with different types 
	result := int(float64(numInt) * numFloat)

	fmt.Printf("numFloat: %f \n", numFloat);

	fmt.Printf("This is result: %d \n", result);
}