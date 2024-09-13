package main

import (
	"fmt"
	"neverquiz/utils"
)

func main() {
	fmt.Println(utils.CountSmileyFace([]string{":)", ";(", ";}", ":-D"}))
	fmt.Println(utils.CountSmileyFace([]string{";D", ":-(", ":-)", ";~)"}))
	fmt.Println(utils.CountSmileyFace([]string{";]", ":[", ";*", ":$", ";-D"}))
}
