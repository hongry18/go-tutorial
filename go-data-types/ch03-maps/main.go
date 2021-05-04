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

	studentsAge := map[string]int{
		"john": 32,
		"bob":  32,
	}

	fmt.Println(studentsAge)
}

func ex02() {
	fmt.Printf("\n# %s\n\n", runFuncName())

	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	fmt.Println(studentsAge)
}

func ex03() {
	fmt.Printf("\n# %s\n\n", runFuncName())

	//studentsAge := map[string]int
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	fmt.Println(studentsAge)
}

func ex04() {
	fmt.Printf("\n# %s\n\n", runFuncName())

	//studentsAge := map[string]int
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	fmt.Println("Bob's age is", studentsAge["bob"])
}

func ex05() {
	fmt.Printf("\n# %s\n\n", runFuncName())

	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	fmt.Println("Christy's age is", studentsAge["christy"])
}

func ex06() {
	fmt.Printf("\n# %s\n\n", runFuncName())

	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	if age, exist := studentsAge["christy"]; exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}

}

func ex07() {
	fmt.Printf("\n# %s\n\n", runFuncName())

	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	delete(studentsAge, "bob")

	fmt.Println(studentsAge)
}

func ex08() {
	fmt.Printf("\n# %s\n\n", runFuncName())

	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	delete(studentsAge, "christy")

	fmt.Println(studentsAge)
}

func ex09() {
	fmt.Printf("\n# %s\n\n", runFuncName())

	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}

func ex10() {
	fmt.Printf("\n# %s\n\n", runFuncName())

	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	for _, age := range studentsAge {
		fmt.Printf("Ages %d\n", age)
	}
}

func ex11() {
	fmt.Printf("\n# %s\n\n", runFuncName())

	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	for name := range studentsAge {
		fmt.Printf("Names %s\n", name)
	}
}

func main() {
	ex01()
	ex02()
	ex03()
	ex04()
	ex05()
	ex06()
	ex07()
	ex08()
	ex09()
	ex10()
	ex11()
}
