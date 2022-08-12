package main

import (
	"fmt"
)

func main() {

	one(two(three))
}

func one(f func()) {
	fmt.Println("ONE\n")
	f()
}

func two(f func()) func() {
	return func() {
		fmt.Println("TWO\n")
		f()
	}
}

func three(){
	fmt.Println("THREE\n")
}
