package main

import (
	"fmt"
	"net/http"
	"time"
)

func send(ch chan string, message string) {
	ch <- message
}

func ex01() {
	size := 2
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	go send(ch, "three")
	go send(ch, "four")
	fmt.Println("All data send to the channel...")

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")
}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func ex02() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string, 10)

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds\n", elapsed.Seconds())
}

func send2(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

func read2(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}

func ex03() {
	ch := make(chan string, 1)
	send2(ch, "Hello world")
	read2(ch)
}

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done replicating!"
}

// 멀티플렉싱
func ex04() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go process(ch1)
	go replicate(ch2)

	for i := 0; i < 2; i++ {
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}
}

func main() {
	// ex01()
	// ex02()
	// ex03()
	ex04()
}
