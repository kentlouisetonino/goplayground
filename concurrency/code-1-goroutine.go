package main

import "fmt"

func test(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go test(0)

	var input string
	fmt.Scanln(&input)
}
