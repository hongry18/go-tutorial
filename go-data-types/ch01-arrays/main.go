package main

import (
	"fmt"
)

func ex01() {
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])
}

func ex02() {
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities: ", cities)
}

func ex03() {
	q := [...]int{1, 2, 3}
	fmt.Println(q)
}

func ex04() {
	n := [...]int{99: -1}
	fmt.Println("First Position: ", n[0])
	fmt.Println("Last Position: ", n[99])
	fmt.Println("Length: ", len(n))
}

func ex05() {
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}

		fmt.Println("Row", i, twoD[i])
	}

	fmt.Println("\n All at once: ", twoD)
}

func ex06() {
	var threeD [3][5][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 2; k++ {
				threeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
			}
		}
	}
	fmt.Println("\nAll at once:", threeD)
}

func main() {
	ex01()
	ex02()
	ex03()
	ex04()
	ex05()
	ex06()
}
