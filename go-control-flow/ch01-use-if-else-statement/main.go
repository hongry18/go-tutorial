package main

import "fmt"

func main() {
	ex01()
	ex02()
	ex03()
}

func ex01() {
	x := 27
	if x%2 == 0 {
		fmt.Println(x, "is even")
	}
}

func givemeanumber() int {
	return -1
}

func ex02() {
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digit")
	}
}

func somenumber() int {
	return -7
}

func ex03() {
	if num := somenumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digit")
	}

	//fmt.Println(num)
}
