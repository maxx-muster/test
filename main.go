package main

import (
	"fmt"
)

func main() {
	DoThis(3)
}

func DoThis(num int) {
	fmt.Print(num * num)
}
