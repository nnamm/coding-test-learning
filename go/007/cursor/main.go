package main

import (
	"bufio"
	"fmt"
	"os"
)

// DFS により連結成分（島）の個数を数える。
// 入力: h, w の後に h 行の 0/1 文字列（行長が w 未満の場合は右側を 0 で埋める）。
// 出力: 島の個数
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var height, width int
	if _, err := fmt.Fscan(reader, &height, &width); err != nil {
		return
	}

	grid := make([][]byte, height)
	for row := 0; row < height; row++ {
		var line string
		if _, err := fmt.Fscan(reader, &line); err != nil {
			line = ""
		}

		if len(line) < width {
			padded := make([]byte, width)
			copy(padded, []byte(line))
			for i := len(line); i < width; i++ {
				padded[i] = '0'
			}
			grid[row] = padded
		} else {
			bytes := []byte(line)
			if len(bytes) > width {
				bytes = bytes[:width]
			}
			grid[row] = bytes
		}
	}

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		visited[r][c] = true
		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= height || nc < 0 || nc >= width {
				continue
			}
			if grid[nr][nc] == '1' && !visited[nr][nc] {
				dfs(nr, nc)
			}
		}
	}

	islandCount := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] == '1' && !visited[r][c] {
				dfs(r, c)
				islandCount++
			}
		}
	}

	fmt.Fprintln(writer, islandCount)
}
