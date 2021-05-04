package main

import (
	"fmt"
	"runtime"
)

func runFuncName() {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	fmt.Printf("\n### %s\n", f.Name())
}

func fibonacci(n int) []int {
	runFuncName()

	if n < 2 {
		return nil
	}

	result := make([]int, 0)

	for i := 0; i < n; i++ {
		if i == 0 {
			result = append(result, 1)
			continue
		}

		if i == 1 {
			result = append(result, 1)
			continue
		}

		result = append(result, (result[i-1] + result[i-2]))
	}

	return result
}

func conversionRoma(s string) {
	runFuncName()

	romaList := make(map[rune]int)
	romaList['M'] = 1000
	romaList['D'] = 500
	romaList['C'] = 100
	romaList['L'] = 50
	romaList['X'] = 10
	romaList['V'] = 5
	romaList['I'] = 1

	arabicVals := make([]int, len(s)+1)

	for i, digit := range s {
		if val, exist := romaList[digit]; exist {
			arabicVals[i] = val
		} else {
			fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", s, digit)
			return
		}
	}

	var sum int = 0
	for i := 0; i < len(s); i++ {
		if arabicVals[i] < arabicVals[i+1] {
			arabicVals[i] = -arabicVals[i]
		}

		sum += arabicVals[i]
	}

	fmt.Printf("total: %d", sum)
}

func main() {
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci(7))
	fmt.Println(fibonacci(8))

	conversionRoma("MCLX")
	conversionRoma("MCMXCIX")
	conversionRoma("MCMZ")
	conversionRoma("MCM")

}
