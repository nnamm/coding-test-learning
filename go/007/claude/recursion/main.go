package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. 入力値の読み取りと初期化
	scanner := bufio.NewScanner(os.Stdin)

	// 行数と列数の読み取り
	var h, w int
	if _, err := fmt.Scan(&h, &w); err != nil {
		fmt.Fprintf(os.Stderr, "error reading dimensions: %v\n", err)
		return
	}

	// 入力検証
	if h <= 0 || w <= 0 || h > 300 || w > 300 {
		fmt.Fprintf(os.Stderr, "invalid dimensions: h=%d, w=%d\n", h, w)
		return
	}

	// グリッドを[][]byteとして定義（メモリ効率、アクセス速度向上）
	grid := make([][]byte, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]byte, w)

		// 各行の読み取り
		if scanner.Scan() {
			line := scanner.Text()
			// 行の長さを調整（短い場合は'0'でパディング、長い場合は切り詰め）
			for j := 0; j < w; j++ {
				if j < len(line) {
					grid[i][j] = line[j]
				} else {
					grid[i][j] = '0' // デフォルトは水'0'
				}
			}
		} else {
			// 入力が不足している場合はすべて水で埋める
			for j := 0; j < w; j++ {
				grid[i][j] = '0'
			}
		}
	}

	// スキャナエラーのチェック
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		return
	}

	// 島の数を数えて出力
	islandCount := countIslands(grid, h, w)
	fmt.Println(islandCount)
}

// countIslands ハグリッド内の島の数を数える
// BFSの実装と比較しやすいよう、アルゴリズムの責任を明確に分離
func countIslands(grid [][]byte, h, w int) int {
	// 訪問済みしフラグの初期化
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	islandCount := 0

	// 前グリッドをスキャンして未発見の島を探索
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			// 未訪問の陸地を発見した場合、新しい島として探索開始
			if grid[i][j] == '1' && !visited[i][j] {
				exploreIsland(grid, visited, i, j, h, w)
				islandCount++
			}
		}
	}

	return islandCount
}

// exploreIsland は指定された位置から連結している陸地をすべて探索する
// DFSを使用して再帰的に隣接する陸地を訪問済みにマーク
func exploreIsland(grid [][]byte, visited [][]bool, row, col, h, w int) {
	// 現在位置を訪問済みにマーク
	visited[row][col] = true

	// 上下左右の4方向への移動パターンを定義
	// （時計回りでも反時計回りでもOK）
	directions := []struct{ dr, dc int }{
		{-1, 0}, // 上
		{1, 0},  // 下
		{0, -1}, // 左
		{0, 1},  // 右
	}

	// 各方向に対して探索を実行
	for _, dir := range directions {
		newRow, newCol := row+dir.dr, col+dir.dc

		// 境界チェック、陸地チェック、未訪問チェックを一括実行
		if isValidPosition(newRow, newCol, h, w) && grid[newRow][newCol] == '1' && !visited[newRow][newCol] {
			// 条件を満たす場合、再帰的に探索を継続
			exploreIsland(grid, visited, newRow, newCol, h, w)
		}
	}
}

// isValidPosition は指定された座標がグリッドの境界内にあるかチェック
// 境界チェックロジックと関数を分離することで可読性と再利用性を向上
func isValidPosition(row, col, h, w int) bool {
	return row >= 0 && row < h && col >= 0 && col < w
}
