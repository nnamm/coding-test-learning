package main

import (
	"bufio"
	"fmt"
	"os"
)

// Point構造体は探索中の地点を表現する
// rowとcolで座標を、distanceでスタートからの歩数を管理
type Point struct {
	row, col, distance int
}

func main() {
	// 1. 入力の読み取りと地形データの構築
	var H, W int
	fmt.Scan(&H, &W)

	// 迷路データを文字列スライスとして保存
	// [][]stringではなく[]stringを使用することで、メモリ効率と処理速度を向上
	maze := make([]string, H)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < H; i++ {
		scanner.Scan()
		maze[i] = scanner.Text()
	}

	// 2. BFS用データ構造の初期化
	queue := []Point{}                                      // 探索待ちの地点を管理するキュー
	visited := make([][]bool, H)                            // 訪問済みフラグの二次元配列
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上、下、左、右の移動方向

	// visitedの二次元配列を適切に初期化
	// make([][]bool, H)では一次元目のみが作成されるため、各行を個別に初期化
	for i := 0; i < H; i++ {
		visited[i] = make([]bool, W)
	}

	// 3. スタート地点の探索と初期設定
	var startRow, startCol int
	found := false
	for i := 0; i < H && !found; i++ {
		for j := 0; j < W; j++ {
			if maze[i][j] == 'S' {
				startRow, startCol = i, j
				found = true // 早期終了でパフォーマンス向上
				break
			}
		}
	}

	// スタート地点をキューに追加し、訪問済みとしてマーク
	// 論理的一貫性のため、キューへの追加と同時に訪問済みフラグを設定
	startPoint := Point{startRow, startCol, 0}
	queue = append(queue, startPoint)
	visited[startRow][startCol] = true

	// 4. BFS本体：キューが空になるまで探索を継続
	for len(queue) > 0 {
		// キューから最初の要素を取り出し（FIFO操作）
		currentPoint := queue[0]
		queue = queue[1:]

		// 現在地点がゴールかどうかを最初に判定
		// この時点で最短距離が確定しているため、ここで判定するのが最適
		if maze[currentPoint.row][currentPoint.col] == 'G' {
			fmt.Println(currentPoint.distance)
			return
		}

		// 上下左右の隣接地点を探索
		for _, direction := range directions {
			newRow := currentPoint.row + direction[0]
			newCol := currentPoint.col + direction[1]

			// 移動可能で未訪問の地点のみをキューに追加
			if canMove(newRow, newCol, H, W, maze) && !visited[newRow][newCol] {
				// 新しい地点を作成し、距離を1増加させる
				newPoint := Point{
					row:      newRow,
					col:      newCol,
					distance: currentPoint.distance + 1,
				}

				// キューに追加し、同時に訪問済みとしてマーク
				// 重複探索を防ぐため、キューに追加する時点で訪問済みフラグを設定
				queue = append(queue, newPoint)
				visited[newRow][newCol] = true
			}
		}
	}

	// 5. ゴールに到達できない場合
	fmt.Println(-1)
}

// canMove関数：指定された座標が移動可能かどうかを判定
// 境界チェックと地形チェックを一箇所で実行することで、コードの重複を排除
func canMove(row, col, maxRow, maxCol int, maze []string) bool {
	// 境界チェック：配列の範囲内かどうかを確認
	if row < 0 || row >= maxRow || col < 0 || col >= maxCol {
		return false
	}

	// 地形チェック：壁（'1'）でなければ移動可能
	// '0'、'S'、'G'はすべて移動可能として扱う
	return maze[row][col] != '1'
}
