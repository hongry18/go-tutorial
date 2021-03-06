package main

import (
	"fmt"
	"io"
	"os"
)

func ex01() {
	for i := 1; i < 5; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}
}

func ex02() {
	f, err := os.Create("notes.txt")
	if err != nil {
		return
	}

	defer f.Close()

	if _, err = io.WriteString(f, "Learning Go!"); err != nil {
		return
	}

	f.Sync()
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic("Panic in g() (major)")
	}

	defer fmt.Println("defer in g()", i)
	fmt.Println("Printing in g()", i)
	g(i + 1)
}

func ex03() {
	// recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	g(0)
	fmt.Println("Program finished successfully!")
}
func main() {
	ex01()
	ex02()
	ex03()
}
