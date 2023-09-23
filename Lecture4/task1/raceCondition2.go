package main

import "time"

var sharedMap map[string]int = map[string]int{}

func runSimpleMapReader() {
	for {
		var _ int = sharedMap["key"]
	}
}

func runSimpleMapWriter() {
	for {
		sharedMap["key"] = sharedMap["key"] + 1
	}
}

func main() {
	sharedMap["key"] = 0

	go runSimpleMapReader()
	go runSimpleMapWriter()
	time.Sleep(10 * time.Second)
}
