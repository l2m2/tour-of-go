package main

import "fmt"

func fibonacci() func() int {
	t1, t2 := 0, 1
	return func() int {
		t := t1
		t1, t2 = t2, t1+t2
		return t
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
