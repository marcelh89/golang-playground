package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// direct call (blocking)
	f("direct")

	// as go routine (async call) - call f fuction
	go f("goroutine")

	// as go routine (async call) - define and call function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

}
