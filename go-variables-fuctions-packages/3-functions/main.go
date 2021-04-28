package main

import (
	"os"
	"strconv"
)

/*
// custom functions
func name(params) (results) {
	body-content
}
*/

func main() {
	number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])
	println("Sum: ", number1+number2)

	println("Sum: ", sum(os.Args[1], os.Args[2]))
	println("Sum: ", sum2(os.Args[1], os.Args[2]))
	println(calc(os.Args[1], os.Args[2]))

	fName := "John"
	updateName(&fName)
	println(fName)
}

func sum(number1, number2 string) int {
	i1, _ := strconv.Atoi(number1)
	i2, _ := strconv.Atoi(number2)
	return i1 + i2
}

func sum2(n1, n2 string) (result int) {
	i1, _ := strconv.Atoi(n1)
	i2, _ := strconv.Atoi(n2)
	result = i1 + i2
	return
}

func calc(n1, n2 string) (sum int, mul int) {
	i1, _ := strconv.Atoi(n1)
	i2, _ := strconv.Atoi(n2)
	sum = i1 + i2
	mul = i1 * i2
	return
}

func updateName(n *string) {
	// pointer
	*n = "David"
}
