package main

import (
	"fmt"

	"github.com/azihsoyn/gollection"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res, _ := gollection.New(arr).Map(func(v int) int {
		return v * 2
	}).Result()
	fmt.Println("origin : ", arr)
	fmt.Println("ret    : ", res)
}
