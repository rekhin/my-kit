package main

import (
	"fmt"
)

func main() {
	items := make([]int, 100)
	for i := 0; i < 100; i++ {
		items[i] = i
	}
	
	Slice[int](items, 6, func(chunk []int) {
		fmt.Println(chunk)
	})
}

func Slice[T any](all []T, size int, cb func(chunk []T)) {
	for ;len(all) > size; all = all[size:len(all)] {
		cb(all[0:size])
	}
	if len(all) > 0 {
		cb(all)
	}
}
