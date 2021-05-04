package main

import (
	"fmt"
	"math"
)

func fizzbuzz() {
	for i := 1; i < 101; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func negativePanic() {
	val := 0
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &val)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover negativePanic", r)
		}
	}()
	if val < 0 {
		panic("negative number")
	}
	fmt.Println("You entered: ", val)
}

func challenge02() {
	fmt.Println(sqrt(25))
	fmt.Println(math.Sqrt(25))
}

func sqrt(num float64) float64 {
	currguess := 1.0
	prevguess := 0.0

	for count := 1; count <= 10; count++ {
		prevguess = currguess
		currguess = prevguess - (prevguess*prevguess-num)/(2*prevguess)
		if currguess == prevguess {
			break
		}
		fmt.Println("A guess for square root is ", currguess)
	}
	return currguess
}

func main() {
	fizzbuzz()
	negativePanic()

	challenge02()
}
