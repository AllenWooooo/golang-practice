package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func WalkPrint(t *tree.Tree, depth int) (sum int) {
	for i := 0; i < depth*3; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%3d\n", t.Value)
	sum += t.Value

	if t.Right != nil {
		sum += WalkPrint(t.Right, depth)
	}
	if t.Left != nil {
		sum += WalkPrint(t.Left, depth+1)
	}

	return
}

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	go Walk(t.Left, ch)
	ch <- t.Value
	go Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < cap(ch1) && i < cap(ch2); i++ {
		v1 := <-ch1
		v2 := <-ch2

		if v1 != v2 {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)
	//WalkPrint(t1, 0)
	//WalkPrint(t2, 0)
	//WalkPrint(t3, 0)
	fmt.Println(Same(t1, t1))
	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t1, t3))
}
