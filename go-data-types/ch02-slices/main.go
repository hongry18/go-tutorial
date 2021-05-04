package main

import (
	"fmt"
	"runtime"
)

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func ex01() {
	fmt.Printf("\n# %s\n\n", runFuncName())
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "Obtober", "November", "December"}
	fmt.Println(months)
	fmt.Println(len(months))
	fmt.Println(cap(months))
}

func ex02() {
	fmt.Printf("\n# %s\n\n", runFuncName())
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "Obtober", "November", "December"}
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))
}

func ex03() {
	fmt.Printf("\n# %s\n\n", runFuncName())
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "Obtober", "November", "December"}
	quarter2 := months[3:6]
	quarter2Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))
	fmt.Println(quarter2[3:6])
}

func ex04() {
	fmt.Printf("\n# %s\n\n", runFuncName())
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
}

func ex05() {
	fmt.Printf("\n# %s\n\n", runFuncName())
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	fmt.Println("Before", letters)
	letters[remove] = letters[len(letters)-1]
	fmt.Println("Middle", letters)
	letters = letters[:len(letters)-1]

	fmt.Printf("After: %v cap: %d\n", letters, cap(letters))
}

func ex06() {
	fmt.Printf("\n# %s\n\n", runFuncName())
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	s1 := letters[0:2]
	s2 := letters[1:4]

	s1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice 2", s2)
}

func ex07() {
	fmt.Printf("\n# %s\n\n", runFuncName())
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	s1 := letters[0:2]
	s2 := make([]string, 3)
	copy(s2, letters[1:4])

	s1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice 2", s2)
}

func main() {
	ex01()
	ex02()
	ex03()
	ex04()
	ex05()
	ex06()
	ex07()
}
