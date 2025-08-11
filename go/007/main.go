package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. 入力値の読み取り（地形データの作成）
	var h, w int
	if _, err := fmt.Scan(&h, &w); err != nil {
		return
	}

	grid := make([][]string, h)
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < h; i++ {
		input.Scan()
		line := input.Text()
		if len(line) < w {
			padding := make([]byte, w-len(line))
			for i := range padding {
				padding[i] = '0'
			}
			line += string(padding)
		}
		grid[i] = strings.Split(line, "")
	}

	// 2. DFS用データ構造の初期化
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	islandCount := 0

	// 3. DFS処理
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == "1" && !visited[i][j] {
				dfsExploreIsland(grid, visited, i, j, h, w)
				islandCount++
			}
		}
	}

	// 4. 島数の出力
	fmt.Println(islandCount)
}

func dfsExploreIsland(grid [][]string, visited [][]bool, row, col, maxRow, maxCol int) {
	visited[row][col] = true

	directions := [...]struct{ dr, dc int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range directions {
		nr, nc := row+dir.dr, col+dir.dc
		// 境界チェック、島チェック、訪問済みチェック
		if nr >= 0 && nr < maxRow && nc >= 0 && nc < maxCol && grid[nr][nc] == "1" && !visited[nr][nc] {
			dfsExploreIsland(grid, visited, nr, nc, maxRow, maxCol)
		}
	}
}
