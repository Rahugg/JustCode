package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	rwMutex       = &sync.RWMutex{}
	customSyncMap = make(map[string]int)
)

func readCustomSyncMap(ctx context.Context, channel chan int) {
	intCh := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("readCustomSyncMap is done")
			channel <- intCh
			return
		default:
			rwMutex.RLock()
			var _ int = customSyncMap["key"]
			rwMutex.RUnlock()
			intCh++
		}
	}
}

func writeCustomSyncMap(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writeCustomSyncMap is done")
			return
		default:
			rwMutex.Lock()
			customSyncMap["key"] += 1
			rwMutex.Unlock()
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	//defer
	channel := make(chan int, 1)
	customSyncMap["key"] = 0
	for i := 0; i < 20; i++ {
		go readCustomSyncMap(ctx, channel)
	}
	for i := 0; i < 20; i++ {
		go writeCustomSyncMap(ctx)
	}
	defer cancel()

	result := 0
	for i := 0; i < 20; i++ {
		result += <-channel
	}
	fmt.Println("Read map value:", result)
	fmt.Println("Write map value:", customSyncMap["key"])
}
