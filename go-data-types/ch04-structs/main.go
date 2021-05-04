package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func ex01() {
	runFuncName()

	employee := Employee{1001, "John", "Doe", "Doe's Street"}
	fmt.Println(employee)

	employee.ID = 1001
	fmt.Println(employee.FirstName)

	emp2 := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(emp2)
	emp3 := &emp2
	emp3.FirstName = "David"
	fmt.Println(emp2)
}

type Person2 struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee2 struct {
	Person2
	ManagerID int
}

type Contractor2 struct {
	Person2
	CompanyID int
}

func ex02() {
	runFuncName()

	employee := Employee2{
		Person2: Person2{FirstName: "John"},
	}

	employee.LastName = "Doe"
	fmt.Println(employee)
}

type Person3 struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee3 struct {
	Person3
	ManagerId int
}

type Contractor struct {
	Person3
	CompanyID int
}

func ex03() {
	employee := []Employee3{
		Employee3{
			Person3: Person3{
				LastName: "Doe", FirstName: "John",
			},
		},
		Employee3{
			Person3: Person3{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	data, _ := json.Marshal(employee)
	fmt.Printf("%s\n", data)

	var decoded []Employee3
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v\n", decoded)
}

func main() {
	ex01()
	ex02()
	ex03()
}
