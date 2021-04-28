package main

import (
    "github.com/myuser/calculator"
    "rsc.io/quote"
)

func main() {
    total := calculator.Sum(3, 5)
    println(total)
    println("Version: ", calculator.Version)
    println("Log Message : ", calculator.LogMessage)
    println(quote.Hello())
}
