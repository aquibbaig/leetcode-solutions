package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var (
		index = -1
		count = 0
		arr = [5]int{1, 2, 3, 3, 3}
		target = 3
	)
	rand.Seed(time.Now().UTC().UnixNano())
	for idx, num := range arr {
		if num == target {
			count += 1
			if rand.Intn(count) == 0 {
				index = idx
			}
		}
	}
	fmt.Println(index)
}
