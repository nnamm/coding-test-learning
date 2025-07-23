package main

import (
	"fmt"
	"strconv"
	"time"
)

var fibMap map[int]int

func main() {
	fmt.Println("階段数を入力してください：")

	var s string
	fmt.Scan(&s)

	// 入力数値チェック
	var i int
	i, _ = strconv.Atoi(s)
	if i < 1 || i > 50 {
		fmt.Println("入力エラー: ", s)
		return
	}

	// グローバルマップの初期化
	if fibMap == nil {
		fibMap = make(map[int]int)
	}

	// 処理時間計測（開始）
	start := time.Now()

	// 実処理部分
	fmt.Printf("入力：%d\n", i)
	fmt.Printf("出力：%d\n", fibonacciOptimazed(i))

	// 処理時間計測（終了）
	elapsed := time.Since(start)
	fmt.Printf("実行時間：%s\n", elapsed)
}

func fibonacciOptimazed(n int) int {
	// 階段数が2以下の場合、入力値がそのまま解となる
	if n <= 2 {
		fibMap[n] = n
		return n
	}

	// 階段数が3以上の場合、メモ化して処理
	if val, ok := fibMap[n]; ok {
		return val
	} else {
		fibMap[n] = fibonacciOptimazed(n-1) + fibonacciOptimazed(n-2)
		return fibMap[n]
	}
}
