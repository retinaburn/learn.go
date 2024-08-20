package main

import "fmt"

func fibonacci() func() int {
	first := 0
	second := 1
	return func() int {
		val := first
		first = second
		second = val + first

		return val

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
