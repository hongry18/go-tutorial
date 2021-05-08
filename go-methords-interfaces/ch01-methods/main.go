package main

import (
	"fmt"
	"go-methods-interfaces/ch01-methods/geometry"
	"strings"
)

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

// 메소드 선언
func ex01() {
	t := triangle{size: 3}
	fmt.Println(t.perimeter())
}

type square struct {
	size int
}

func (s square) perimeter() int {
	return s.size * 4
}

func ex02() {
	t := triangle{3}
	s := square{4}

	fmt.Println(t.perimeter())
	fmt.Println(s.perimeter())
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

// 메서드의 포인터
func ex03() {
	t := triangle{3}
	t.doubleSize()
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter:", t.perimeter())
}

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

// 다른 형식에 대한 메서드 선언
func ex04() {
	s := upperstring("test")
	fmt.Println(s)
	fmt.Println(s.Upper())
}

type coloredTriangle struct {
	triangle        // triangle
	color    string // color string
}

// 메서드 포함
func ex05() {
	t := coloredTriangle{triangle{3}, "red"}
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter:", t.triangle.perimeter())
}

func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

// method overload
func ex06() {
	t := coloredTriangle{triangle{3}, "red"}
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter:", t.perimeter())

	fmt.Println("Perimeter colored:", t.perimeter())
	fmt.Println("Perimeter normal:", t.triangle.perimeter())
}

func ex07() {
	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Perimeter", t.Perimeter())
}

func ex08() {
	t := geometry.Triangle{}
	t.SetSize(3)
	// fmt.Println("Size", t.size)
	fmt.Println("Perimeter", t.Perimeter())
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
}
