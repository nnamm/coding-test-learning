// 木構造走査の模範解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Node は各ノードの子ノード情報を表す
type Node struct {
	left  int
	right int
}

func main() {
	// 入力データの読み取りと木構造の構築
	tree, err := buildTree()
	if err != nil {
		fmt.Fprintf(os.Stderr, "入力エラー: %v\n", err)
		return
	}

	// 各走査結果を出力
	fmt.Print("Preorder: ")
	preOrder(1, tree)
	fmt.Println()

	fmt.Print("Inorder: ")
	inOrder(1, tree)
	fmt.Println()

	fmt.Print("Postorder: ")
	postOrder(1, tree)
	fmt.Println()
}

// buildTree は標準入力から木構造を読み取り、mapとして構築する
func buildTree() (map[int]Node, error) {
	scanner := bufio.NewScanner(os.Stdin)

	// ノード数の読み取り
	if !scanner.Scan() {
		return nil, fmt.Errorf("ノード数の読み取りに失敗")
	}

	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, fmt.Errorf("ノード数の変換に失敗: %v", err)
	}

	// 木構造を格納するmap
	tree := make(map[int]Node, n)

	// 各ノードの情報を読み取り
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf("ノード情報の読み取りに失敗")
		}

		// 行を空白で分割
		parts := strings.Fields(scanner.Text())
		if len(parts) == 0 {
			return nil, fmt.Errorf("空の行が入力されました")
		}

		// ノード番号の取得
		nodeNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("ノード番号の変換に失敗: %v", err)
		}

		// 左の子と右の子の情報を取得（デフォルトは-1）
		left, right := -1, -1

		if len(parts) > 1 {
			if parts[1] != "-1" {
				left, err = strconv.Atoi(parts[1])
				if err != nil {
					return nil, fmt.Errorf("左の子の変換に失敗: %v", err)
				}
			}
		}

		if len(parts) > 2 {
			if parts[2] != "-1" {
				right, err = strconv.Atoi(parts[2])
				if err != nil {
					return nil, fmt.Errorf("右の子の変換に失敗: %v", err)
				}
			}
		}

		// ノード情報をmapに格納
		tree[nodeNum] = Node{left: left, right: right}
	}

	// スキャナーエラーのチェック
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("入力読み取りエラー: %v", err)
	}

	return tree, nil
}

// preOrder は前順走査を実行する（ルート → 左の子 → 右の子）
func preOrder(current int, tree map[int]Node) {
	// ベースケース：ノードが存在しない場合は処理終了
	if current == -1 {
		return
	}

	// 現在のノードを出力
	fmt.Printf("%d ", current)

	// 左の子ノードを再帰的に処理
	preOrder(tree[current].left, tree)

	// 右の子ノードを再帰的に処理
	preOrder(tree[current].right, tree)
}

// inOrder は中順走査を実行する（左の子 → ルート → 右の子）
func inOrder(current int, tree map[int]Node) {
	// ベースケース：ノードが存在しない場合は処理終了
	if current == -1 {
		return
	}

	// 左の子ノードを再帰的に処理
	inOrder(tree[current].left, tree)

	// 現在のノードを出力
	fmt.Printf("%d ", current)

	// 右の子ノードを再帰的に処理
	inOrder(tree[current].right, tree)
}

// postOrder は後順走査を実行する（左の子 → 右の子 → ルート）
func postOrder(current int, tree map[int]Node) {
	// ベースケース：ノードが存在しない場合は処理終了
	if current == -1 {
		return
	}

	// 左の子ノードを再帰的に処理
	postOrder(tree[current].left, tree)

	// 右の子ノードを再帰的に処理
	postOrder(tree[current].right, tree)

	// 現在のノードを出力
	fmt.Printf("%d ", current)
}
