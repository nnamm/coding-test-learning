package main

import (
	"fmt"
	"strconv"
	"time"
)

// ./go/005/main.goの解答でも十分である。完全に動作し、期待される性能改善も達成しているため。
// 本プログラムはこれらは「より良い選択肢」としてあるものである。
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

	// メモ化版フィボナッチ（トップダウン）
	start := time.Now()
	result1 := fibonacciMemoized(i)
	elapsed := time.Since(start)
	fmt.Printf("入力（トップダウン）：%d\n", i)
	fmt.Printf("出力（トップダウン）：%d\n", result1)
	fmt.Printf("実装時間：%s\n", elapsed)

	// ボトムアップ方式のフィボナッチ
	start = time.Now()
	result2 := fibonacciBottomUp(i)
	elapsed = time.Since(start)
	fmt.Printf("入力（ボトムアップ）：%d\n", i)
	fmt.Printf("出力（ボトムアップ）：%d\n", result2)
	fmt.Printf("実装時間：%s\n", elapsed)
}

// メモ化を内容したフィボナッチ関数
func fibonacciMemoized(n int) int {
	memo := make(map[int]int)
	return fibHelper(n, memo)
}

// メモ化を使ったヘルパー関数
func fibHelper(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}

	if val, exists := memo[n]; exists {
		return val
	}

	memo[n] = fibHelper(n-1, memo) + fibHelper(n-2, memo)
	return memo[n]
}

// ボトムアップ方式の動的プログラミング版（参考例）
func fibonacciBottomUp(n int) int {
	if n <= 2 {
		return n
	}

	// 配列を使った反射的な計算
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
