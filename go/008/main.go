package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type children struct {
	left  int
	right int
}

func main() {
	// 1. 入力値の読み取り（木構造データ管理）
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Fprintf(os.Stderr, "error reading node count: %v\n", err)
		return
	}

	// 木構造の定義
	tree := make(map[int]children, n)

	// 各行の読み取り
	scanner := bufio.NewScanner(os.Stdin)
	for range n {
		scanner.Scan()
		text := strings.Split(scanner.Text(), " ")

		// 入力ノードを数値に変換
		var nodes []int
		for _, t := range text {
			node, err := strconv.Atoi(t)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
				return
			}
			nodes = append(nodes, node)
		}

		// 子孫ノードを調整
		// mapのkeyとなる数字は入力必須が前提
		// 以降2-3文字目を左の子、右の子とする
		// 左右の子が入力されていない場合-1を設定
		switch len(nodes) {
		case 0:
			fmt.Fprintf(os.Stderr, "key not input\n")
			return
		case 1: // keyのみ
			tree[nodes[0]] = children{left: -1, right: -1}
		case 2: // key、左の子
			tree[nodes[0]] = children{left: nodes[1], right: -1}
		default: // key、左の子、右の子、それ以降は木構造対象外
			tree[nodes[0]] = children{left: nodes[1], right: nodes[2]}
		}
	}

	// スキャナエラーのチェック
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
	}

	// 2. 木構造走査
	preOrder(1, tree)
	fmt.Println("")
	inOrder(1, tree)
	fmt.Println("")
	postOrder(1, tree)
	fmt.Println("")
}

// 3. 前順走査処理
func preOrder(current int, tree map[int]children) {
	// ノードが存在しない場合は何もしない
	if current == -1 {
		return
	}

	fmt.Printf("%d ", current)
	preOrder(tree[current].left, tree)
	preOrder(tree[current].right, tree)
}

// 4. 中順走査処理
func inOrder(current int, tree map[int]children) {
	// ノードが存在しない場合は何もしない
	if current == -1 {
		return
	}

	inOrder(tree[current].left, tree)
	fmt.Printf("%d ", current)
	inOrder(tree[current].right, tree)
}

// 5. 後順走査処理
func postOrder(current int, tree map[int]children) {
	// ノードが存在しない場合は何もしない
	if current == -1 {
		return
	}

	postOrder(tree[current].left, tree)
	postOrder(tree[current].right, tree)
	fmt.Printf("%d ", current)
}
