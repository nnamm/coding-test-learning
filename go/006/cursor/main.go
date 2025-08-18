package main

import (
	"bufio"
	"fmt"
	"os"
)

// pos represents a coordinate in the maze and the distance from the start.
type pos struct {
	r, c int // row, col
	dist int // distance from the start (BFS level)
}

func main() {
	// --- 1. 入力読み取り ----------------------------------------------------
	var h, w int
	if _, err := fmt.Scan(&h, &w); err != nil {
		// 入力不正時は何も出力せず終了
		return
	}

	maze := make([]string, h)
	in := bufio.NewScanner(os.Stdin)
	for i := 0; i < h; i++ {
		in.Scan()
		line := in.Text()
		// 入力が不足している場合は右側を壁として補完
		// （面接環境の入力保証があるなら不要だが、安全側に倒す）
		if len(line) < w {
			padding := make([]byte, w-len(line))
			for i := range padding {
				padding[i] = '1'
			}
			line += string(padding)
		}
		maze[i] = line
	}

	// --- 2. スタート地点の検索 ----------------------------------------------
	sr, sc := -1, -1
	for i := 0; i < h && sr == -1; i++ {
		for j := 0; j < w; j++ {
			if maze[i][j] == 'S' {
				sr, sc = i, j
				break
			}
		}
	}
	if sr == -1 {
		// スタートが無い場合は探索不能
		fmt.Println(-1)
		return
	}

	// --- 3. BFS 用データ構造の初期化 ----------------------------------------
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	dirs := [...]struct{ dr, dc int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// queue をスライス + head index で実装して O(1) pop
	q := make([]pos, 0, h*w)
	q = append(q, pos{sr, sc, 0})
	visited[sr][sc] = true

	// --- 4. 幅優先探索本体 ---------------------------------------------------
	head := 0 // 取り出し位置
	for head < len(q) {
		cur := q[head]
		head++ // pop

		// ゴール判定
		if maze[cur.r][cur.c] == 'G' {
			fmt.Println(cur.dist)
			return
		}

		// 隣接マスをチェック
		for _, d := range dirs {
			nr, nc := cur.r+d.dr, cur.c+d.dc
			if nr < 0 || nr >= h || nc < 0 || nc >= w {
				continue // 境界外
			}
			if visited[nr][nc] || maze[nr][nc] == '1' { // 既訪問 or 壁
				continue
			}
			visited[nr][nc] = true
			q = append(q, pos{nr, nc, cur.dist + 1})
		}
	}

	// --- 5. ゴールに到達出来なかった場合 ------------------------------------
	fmt.Println(-1)
}
