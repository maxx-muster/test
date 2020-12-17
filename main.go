package main

import (
	"fmt"
)

func main() {
	DoThis(3)
}

func DoThis(num int) {
	num += 3
	fmt.Print(num * num)
}
