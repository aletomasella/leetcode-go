package main

import (
	"fmt"

	"github.com/aletomasella/leetcode-go/utils"
)

func main() {
	var nums = []int{2, 7, 11, 15}

	for i := 0; i < 100_000_000; i++ {
		nums = append(nums, i)
	}

	res := utils.TwoSum(nums, 100_000_000)

	fmt.Println("DONE")

	fmt.Println(res)

}
