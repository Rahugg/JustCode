package main

import (
	"fmt"
	"time"
)

var cnt int

func increment() {
	for {
		cnt++
	}
}

func read() {
	for {
		var val int = cnt
		if val%100 == 0 {
			return
		}
	}
}

// race condition
func main() {
	go increment()
	go read()
	time.Sleep(10 * time.Second)
	fmt.Println(cnt)
}
