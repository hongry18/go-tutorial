package main

import (
	"errors"
	"fmt"
	"time"
)

type Employee1 struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func getInformation1(id int) (*Employee1, error) {
	emp, err := apiCallEmployee1(1000)
	if err != nil {
		return nil, err
	}

	return emp, nil
}

func getInformation2(id int) (*Employee1, error) {
	emp, err := apiCallEmployee1(1000)
	if err != nil {
		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	}

	return emp, nil
}

func getInformation3(id int) (*Employee1, error) {
	for tries := 0; tries < 3; tries++ {
		emp, err := apiCallEmployee1(1000)
		if err == nil {
			return emp, nil
		}

		fmt.Println("Server is not responding, retrying ...")
		time.Sleep(time.Second * 2)
	}

	return nil, fmt.Errorf("server has failed to respond to get the employee information")
}

func apiCallEmployee1(id int) (*Employee1, error) {
	emp := Employee1{LastName: "Doe", FirstName: "John"}
	return &emp, nil
}

func ex01() {
	emp, err := getInformation1(1001)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(emp)
	}
}

var ErrNotFound = errors.New("employee not found!")

func getInformation4(id int) (*Employee1, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	emp := Employee1{LastName: "Doe", FirstName: "John"}
	return &emp, nil
}

func ex02() {
	emp, err := getInformation4(1000)
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("Not FOUND: %v\n", err)
	} else {
		fmt.Println(emp)
	}
}

func main() {
	ex01()
	ex02()
}
