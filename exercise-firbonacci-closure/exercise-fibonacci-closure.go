package main

import "fmt"

func fibonacci() func() int {
	n, t1, t2 := 0, 0, 1
	return func() int {
		n++
		if n == 1 {
			return 0
		} else if n == 2 {
			return 1
		} else {
			sum := t1 + t2
			t1 = t2
			t2 = sum
			return sum
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
