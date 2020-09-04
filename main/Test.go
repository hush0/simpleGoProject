package main

import (
	"fmt"
	"math/rand"
	"os"
	"simpleGoProject/main/math"
	"time"
)

func producer(header string, channel chan<- string) {

	for {
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func consumer(channel <-chan string) {
	for {
		message := <-channel
		fmt.Println(message)
	}

}

func main() {

	fmt.Println(os.Args)

	if len(os.Args) > 1 {
		fmt.Println("hello go", os.Args[1])
	}

	fmt.Println(math.Add(1, 2))

	channel := make(chan string)
	go producer("dog", channel)
	go producer("cat", channel)
	consumer(channel)
	os.Exit(-1)
}
