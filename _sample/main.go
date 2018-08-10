package main

import (
	"time"
	"fmt"

	"github.com/aakira/multiticker"
)

func main() {

	ticker := multiticker.NewMultiTicker(map[string]time.Duration{
		"a": 4 * time.Second,
		"b": 2 * time.Second,
		"c": 6 * time.Second,
		"d": 10 * time.Second,
	})

	defer ticker.Stop() // you must call stop function

	for c := range ticker.C {
		switch c.Key {
		case "a":
			fmt.Printf("receive: %s, time: %v\n", c.Key, c.Tick)
		case "b":
			fmt.Printf("receive: %s, time: %v\n", c.Key, c.Tick)
		case "c":
			fmt.Printf("receive: %s, time: %v\n", c.Key, c.Tick)
		default:
			fmt.Printf("receive: %s, time: %v\n", c.Key, c.Tick)
			return
		}
	}
}
