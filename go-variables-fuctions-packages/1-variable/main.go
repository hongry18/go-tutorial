package main

func main() {
	/*
		var firstName, lastName string

		var age int

		var (
			firstName, lastName string
			age                 int
		)

		var (
			firstName string = "Jhon"
			lastName string = "Doe"
			age	int = 32
		)

		var (
			firstName = "Jhon"
			lastName = "Doe"
			age = 32
		)

		var (
			firstName, lastName, age = "Jhon", "Doe", 32
		)
	*/

	firstName, lastName := "Jhon", "Doe"
	age := 32

	println(firstName, lastName, age)

	// constants
	const HTTPStatusOk = 200
	const (
		StatusOk              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)
}
