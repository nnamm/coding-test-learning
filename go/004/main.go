package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("階段数を入力してください：")

	var s string
	fmt.Scan(&s)

	var i int
	i, _ = strconv.Atoi(s)
	if i < 1 {
		fmt.Println("入力エラー：", i)
		return
	}

	// Version1
	// f := fibonacciV1()
	// fmt.Println("入力：", i)
	// fmt.Println("出力：", f(i))

	// Version2
	fmt.Println("入力：", i)
	fmt.Println("出力：", fibonacciV2(i))
}

// Version1
//	func fibonacciV1() func(i int) int {
//		return func(i int) int {
//			if i <= 2 {
//				return i
//			}
//			f := fibonacci()
//			return f(i-1) + f(i-2)
//		}
//	}

// Version2
func fibonacciV2(n int) int {
	if n <= 2 {
		return n
	}
	return fibonacciV2(n-1) + fibonacciV2(n-2)
}
