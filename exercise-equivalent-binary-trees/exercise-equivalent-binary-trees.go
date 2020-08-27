package main

import (
	"fmt"
	"reflect"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()
	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()
	var s1, s2 []int
	for i := range ch1 {
		s1 = append(s1, i)
	}
	for i := range ch2 {
		s2 = append(s2, i)
	}
	sort.Ints(s1)
	sort.Ints(s2)
	fmt.Println(s1)
	fmt.Println(s2)
	return reflect.DeepEqual(s1, s2)
}

func main() {
	ok := Same(tree.New(1), tree.New(1))
	fmt.Println(ok)
	ok = Same(tree.New(1), tree.New(2))
	fmt.Println(ok)
}
