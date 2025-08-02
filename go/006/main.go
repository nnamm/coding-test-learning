package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	row, col, distance int
}

func main() {
	// 1. 地形データの作成
	fmt.Print("Input Rows & Cols: ")
	var H, W int
	fmt.Scan(&H, &W)

	maze := make([][]string, H)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < H; i++ {
		scanner.Scan()
		line := scanner.Text()

		row := make([]string, W)
		for j, c := range line {
			if j >= W {
				break
			}
			row[j] = string(c)
		}
		maze[i] = row
	}

	// 初期化
	queue := []Point{}           // 探索待ち地点
	visited := make([][]bool, H) // 訪問済みフラグ
	for i := 0; i < H; i++ {
		row := make([]bool, W)
		for j := 0; j < W; j++ {
			row[j] = false
		}
		visited[i] = row
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上、下、左、右

	// 2. BFS(幅優先探索)
	// 開始地点をキューに登録
	var startRow, startCol int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if maze[i][j] == "S" {
				startRow, startCol = i, j
			}
		}
	}
	startPoint := Point{startRow, startCol, 0}
	queue = append(queue, startPoint)

	// キューにデータがある間、以下の処理を行う
	for len(queue) > 0 {
		currentPoint := queue[0]
		queue = queue[1:]

		// 現在位置がゴールか判定
		if isGoal(currentPoint.row, currentPoint.col, maze) {
			fmt.Println(currentPoint.distance)
			return
		}

		for _, direction := range directions {
			newRow := currentPoint.row + direction[0]
			newCol := currentPoint.col + direction[1]

			if canMove(newRow, newCol, maze) && !isVisited(newRow, newCol, visited) {
				newPoint := Point{
					row:      newRow,
					col:      newCol,
					distance: currentPoint.distance + 1,
				}
				queue = append(queue, newPoint)
				visited[newRow][newCol] = true

				log.Printf("newPoint: %v\n", newPoint)
				log.Printf("queue: %v\n", queue)
				log.Printf("visited: %v\n", visited)

				// 探索位置がゴールか判定
				if isGoal(newRow, newCol, maze) {
					fmt.Println(newPoint.distance)
					return
				}
			}
		}
	}

	// 3. 終了処理
	fmt.Println(-1)
}

func isGoal(curtRow, curtCol int, maze [][]string) bool {
	return maze[curtRow][curtCol] == "G"
}

func canMove(newRow, newCol int, maze [][]string) bool {
	// 最新位置が配列要素範囲内か確認
	var rows, cols int
	rows = len(maze)
	for _, row := range maze {
		cols = len(row)
	}
	if newRow > rows-1 || newCol > cols-1 || newRow < 0 || newCol < 0 {
		return false
	}

	// 最新情報の地形データを確認
	if maze[newRow][newCol] == "1" {
		return false
	}

	// "0", "G"とみなす
	return true // 上記以外ならfalseとする
}

func isVisited(newRow, newCol int, visited [][]bool) bool {
	// 最新位置が配列要素範囲内か確認
	var rows, cols int
	rows = len(visited)
	for _, row := range visited {
		cols = len(row)
	}
	if len(visited) > 0 && (newRow > rows-1 || newCol > cols-1 || newRow < 0 || newCol < 0) {
		return false
	}

	if !visited[newRow][newCol] {
		return false
	}

	return true
}
