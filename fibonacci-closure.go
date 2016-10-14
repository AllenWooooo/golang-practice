package main

import "fmt"

func fibonacci() func() int {
	value1 := 0
	value2 := 1

	return func() int {
		result := value1

		value1 = value2
		value2 = value1 + result

		return result
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
