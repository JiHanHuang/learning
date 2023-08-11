package main

import (
	"math/rand"
	"time"
)

func randArray(num uint) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, num)
	for i := 0; uint(i) < num; i++ {
		//-10^6 ~ 10^6
		arr[i] = rand.Intn(2000000) - 1000000
	}
	return arr
}
