// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import (
	"fmt"
	"time"
)

func func1(ch chan string) {
	fmt.Println("first a")
	time.Sleep(1 * time.Second)
	ch <- "Welcome"
	fmt.Println("first b")
}

func func2(ch chan int) {
	fmt.Println("second a")
	time.Sleep(1 * time.Second)
	ch <- 1
	fmt.Println("second b")
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan int)

	go func1(ch1)
	go func2(ch2)

	fmt.Println("Fired")
	select {
	case val2 := <-ch2:
		fmt.Println(val2)
	case val1 := <-ch1:
		fmt.Println(val1)
	}

}
