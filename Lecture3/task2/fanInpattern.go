package main

import (
	"fmt"
	"sync"
)

//many channels to 1 channel

func send(c <-chan int, out chan int, wg *sync.WaitGroup) {
	for n := range c {
		out <- n
	}
	wg.Done()
}

func Merge(cs ...<-chan int) <-chan int {
	wg := &sync.WaitGroup{}
	wg.Add(len(cs))
	out := make(chan int)

	for _, c := range cs {
		go send(c, out, wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	go func() {
		for _, num := range []int{1, 2, 3, 4} {
			a <- num
		}
		close(a)
	}()
	go func() {
		for _, num := range []int{5, 6, 7, 8} {
			b <- num
		}
		close(b)
	}()
	go func() {
		for _, num := range []int{9, 10} {
			c <- num
		}
		close(c)
	}()

	for num := range Merge(a, b, c) {
		fmt.Println(num)
	}

}
