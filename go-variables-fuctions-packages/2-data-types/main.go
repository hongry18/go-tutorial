package main

import (
	"math"
	"strconv"
)

func ex1() {
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807

	println(integer8, integer16, integer32, integer64)
}

func ex2() {
	var integer16 int16 = 127
	var integer32 int32 = 32767
	// panic
	//println(integer16 + integer32)
	println(integer16, integer32)
}

func ex3() {
	rune := 'G'
	println(rune)
}

func ex4() {
	//var integer32 int32 = 2147483648
	var integer32 int = 2147483648
	println(integer32)
}

func ex5() {
	//var integer uint = -10
	var integer int = -10
	println(integer)
}

func ex6() {
	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807
	println(float32, float64)
}

func ex7() {
	println(math.MaxFloat32, math.MaxFloat64)
}

func ex8() {
	const e = 2.71828
	const Avogadro = 6.022114129e23
	const Planck = 6.62606957e-34
}

func ex9() {
	var featureFlag bool = true
	println(featureFlag)
}

func ex10() {
	var firstName string = "John"
	lastName := "Doe"
	println(firstName, lastName)
}

func ex11() {
	fullName := "John Doe \t(alias \"Foo\")\n"
	println(fullName)
}

func ex12() {
	// default variables
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)
}

func ex13() {
	// convert type
	var integer16 int16 = 127
	var integer32 int32 = 32767
	println(int32(integer16) + integer32)
}

func ex14() {
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(i)
	println(i, s)
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
	ex9()
	ex10()
	ex11()
	ex12()
	ex13()
	ex14()
}
