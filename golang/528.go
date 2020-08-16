package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binarySearch(target int, accum []int) int {
	var (
		lo = 0
		hi = len(accum)-1
	)
	mid := lo + (hi-lo)/2
	for lo < hi {
		if target < accum[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	var (
		arr = [2]int{1,3}
		//index = -1
	)
	abcd := []int{1,2,3}
	fmt.Println(abcd)
	// 1, 2, 3 -> 1 | 3 | 6
	rand.Seed(time.Now().UnixNano())
	var accum []int
	cum := 0
	// generates the cumulative probability array.
	for _, v := range arr {
		cum += v
		accum = append(accum, cum)
	}
	rng := rand.Intn(accum[len(accum)-1])
	res := binarySearch(rng, accum)
	fmt.Println(arr[res])
}
