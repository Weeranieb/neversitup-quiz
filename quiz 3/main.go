package main

import (
	"fmt"
	"neverquiz/utils"
)

func main() {
	fmt.Println(utils.OddInt([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
	fmt.Println(utils.OddInt([]int{0, 1, 0, 1, 0}))
	fmt.Println(utils.OddInt([]int{1, 1, 2}))
	fmt.Println(utils.OddInt([]int{7}))
}
