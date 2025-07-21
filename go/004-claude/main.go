package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	var i int
	i, _ = strconv.Atoi(s)
	if i < 1 {
		fmt.Println("入力エラー：", i)
		return
	}

	fmt.Println("入力：", i)
	fmt.Println("出力：", fibonacci(i))
}

func fibonacci(n int) int {
	if n <= 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
