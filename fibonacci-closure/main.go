package main

import "fmt"

func fibonacci() func() int {
	n1, n2 := 0, 1 // closures for n-1, n-2
	return func() (sum int) {
		sum = n2

		n1, n2 = n2, n1 + n2

		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
