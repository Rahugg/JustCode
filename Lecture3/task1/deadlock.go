package main

import (
	"fmt"
	"sync"
)

//Deadlock in Channel in Golang is a condition that happens when a few goroutines are waiting
//for each other but no goroutine can proceed.

func main() {
	firstCase()
	secondCase()
	thirdCase()
	fourthCase()
	fifthCase()
}

// //reading from an empty channel
func firstCase() {
	ch := make(chan int)
	fmt.Println(<-ch) //reading empty channel

	go func(num chan int) {
		num <- 3
	}(ch)

	i := <-ch //3
	j := <-ch //here we will enter deadlock because the channel is empty
	fmt.Println(i, j)
}

// can't send to the fulled channel
func secondCase() {
	var i int
	num := make(chan int)
	i = 17
	num <- i
	num <- i + 18
	fmt.Println(num)
}

// when we are passing more values than it can store
func thirdCase() {
	ch := make(chan string, 3)

	ch <- "1"
	ch <- "2"
	ch <- "3"
	ch <- "4"

	for channel := range ch {
		fmt.Println(channel)
	}
}

// if you don't use wg.Done() for every your wg.Add() it will deadlock
func fourthCase() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			//defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait() // If you forget to call Done() inside goroutines, it will deadlock.

}

// deadlock because the channel is unbuffered channel
// An unbuffered channel will block until there is a receiver to receive the message.
func fifthCase() {
	var number = make(chan int)
	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup

	number <- 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&wg, i, number, mutex)
	}
	wg.Wait()
	fmt.Println(<-number) // expected output: 0+1+2+3+4 = 10
}

func worker(wg *sync.WaitGroup, id int, number chan int, mutex *sync.Mutex) {
	defer wg.Done()

	mutex.Lock()
	number <- id + <-number
	mutex.Unlock()
}
