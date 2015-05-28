package main  

import (
	"fmt"
	"time"
	)

func foo(out chan string) {
	time.Sleep(1*time.Second)
	out <- "Foo!"
}

func bar(out chan string) {
	time.Sleep(2*time.Second)
	out <- "Bar!"
}

func baz(out chan string, quit chan bool) {
	time.Sleep(3*time.Second)
	out <-"Baz!"
	quit <- true
}

func main() {

	results := make(chan string)

	quit := make(chan bool)
	
	go foo(results)
	go bar(results)
	go baz(results, quit)  // using go allows for concurrent fucntions
	
	fmt.Println("Waiting....")
	for {
		select {
		case message:= <- results:
			fmt.Println(message)
		case <- quit:
			fmt.Println("Done")
			return
		default:
			fmt.Println("Waiting")
			time.Sleep(500 * time.Millisecond)
		}

	}
	fmt.Println("Done.")
}

