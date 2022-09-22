package main

import "fmt"

func main() {
	defer fmt.Println("Mehedi Hasan")
	defer fmt.Println("Taslim Hasan")
	defer myDeferData(4)
	defer fmt.Println("Sudipto Hasan")
	fmt.Println("Mehedi Hasan")
	myDefer(6)
}

func myDefer(n int) int {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
	return int(n)
}

func myDeferData(n int) int {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
	return int(n)
}
