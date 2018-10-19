package main

import (
	"fmt"
	"math/rand"
	"time"
)

//goroutine make the command run concurrent
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	//	go f(0)

	var input string
	fmt.Scanln(&input)
}