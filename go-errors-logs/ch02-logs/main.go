package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func ex01() {
	log.Print("Hey, I'm a log!")
}

func ex02() {
	log.Fatal("Hey, I'm an error log!")
	fmt.Print("Can you see me?")
}

func ex03() {
	log.Panic("Hey, I'm an error log!")
	fmt.Print("Can you see me?")
}

func ex04() {
	log.SetPrefix("main(): ")
	log.Print("Hey, I'm a log!")
	log.Fatal("Hey, I'm an error log!")
}

func ex05() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Hey, I'm a log!")
}

func ex06() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zlog.Print("Hey!! I'a log message")
}

func ex07() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zlog.Debug().Int("EmployeeId", 1001).Msg("Getting employee information")

	zlog.Debug().Str("Name", "John").Send()


}

func main() {
	// ex01()
	// ex02()
	// ex03()
	// ex04()
	// ex05()
	// ex06()
	ex07()
}
