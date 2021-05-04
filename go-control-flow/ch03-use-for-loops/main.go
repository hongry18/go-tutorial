package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ex01() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += 1
	}

	fmt.Println("sum of 1..100 is ", sum)
}

func ex02() {
	var num int64
	rand.Seed(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}
}

func ex03() {
	var num int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Print("Writting inside the loop...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num)
	}
}

func ex04() {
	sum := 0
	for num := 1; num < 101; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}

	fmt.Println(sum)
}

func main() {
	ex01()
	ex02()
	ex03()
	ex04()
}
