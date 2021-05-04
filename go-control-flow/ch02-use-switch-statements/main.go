package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func main() {
	ex01()
	ex02()
	ex03()
	ex04()
	ex05()
	ex06()
}

func ex01() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)

	switch i {
	case 0:
		fmt.Println("zero ...")
	case 1:
		fmt.Println("one ...")
	case 2:
		fmt.Println("two ...")
	default:
		fmt.Println("no match...")
	}

	fmt.Println(i, "ok")
}

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"

	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"

	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}

	return region, continent
}

func ex02() {
	region, continent := location("Irvine")
	fmt.Printf("John works in %s, %s\n", region, continent)
}

func ex03() {
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's weekend, time to rest!")
	}

	fmt.Println(time.Now().Weekday().String())
}

func ex04() {
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "foo@bar.com"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is an phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}
}

func ex05() {
	rand.Seed(time.Now().Unix())
	r := rand.Float64()

	switch {
	case r > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}

	fmt.Println("r is :", r)
}

func ex06() {
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is less than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200\n", num)
	}
}
