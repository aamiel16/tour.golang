package main

import "fmt"

func fibonacci() func() int {
	n1, n2 := 0, 0 // closures for n-1, n-2
	return func() (sum int) {
		if n1 == 0 {
			sum = 1
			n1 += 1
			return
		}

		sum = n1 + n2
		n2 = n1
		n1 = sum

		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
