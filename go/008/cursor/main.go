package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Node represents the indices of the left and right child of a binary tree node.
// A value of -1 means the child does not exist.
type Node struct {
	left, right int
}

func main() {
	tree, root, err := buildTree()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Preorder
	pre := make([]int, 0)
	preorder(root, tree, &pre)
	printTraversal("Preorder", pre)

	// Inorder
	in := make([]int, 0)
	inorder(root, tree, &in)
	printTraversal("Inorder", in)

	// Postorder
	post := make([]int, 0)
	postorder(root, tree, &post)
	printTraversal("Postorder", post)
}

// buildTree reads the input and returns the constructed tree and its root index.
func buildTree() (map[int]Node, int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	// Allow reading long lines (default 64 K may be too small for large inputs).
	scanner.Buffer(make([]byte, 0, 64*1024), 1<<20)

	// --- read number of nodes ---
	if !scanner.Scan() {
		return nil, -1, fmt.Errorf("入力が空です")
	}
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return nil, -1, fmt.Errorf("ノード数が不正です: %v", err)
	}

	tree := make(map[int]Node, n)
	childrenSet := make(map[int]struct{}, n*2)

	// --- read each node line ---
	for read := 0; read < n; {
		if !scanner.Scan() {
			return nil, -1, fmt.Errorf("ノード情報が不足しています (期待 %d 行)", n)
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" { // skip empty lines
			continue
		}
		parts := strings.Fields(line)
		// first value is node id
		nodeID, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, -1, fmt.Errorf("ノード番号が不正です: %v", err)
		}

		left, right := -1, -1
		if len(parts) > 1 && parts[1] != "-1" {
			if left, err = strconv.Atoi(parts[1]); err != nil {
				return nil, -1, fmt.Errorf("左の子が不正です: %v", err)
			}
			childrenSet[left] = struct{}{}
		}
		if len(parts) > 2 && parts[2] != "-1" {
			if right, err = strconv.Atoi(parts[2]); err != nil {
				return nil, -1, fmt.Errorf("右の子が不正です: %v", err)
			}
			childrenSet[right] = struct{}{}
		}

		tree[nodeID] = Node{left: left, right: right}
		read++
	}

	if err := scanner.Err(); err != nil {
		return nil, -1, fmt.Errorf("入力読み取りエラー: %v", err)
	}

	// --- find root (node that never appears as a child) ---
	root := -1
	for id := range tree {
		if _, isChild := childrenSet[id]; !isChild {
			root = id
			break
		}
	}

	if root == -1 {
		return nil, -1, fmt.Errorf("ルートが特定できません")
	}

	return tree, root, nil
}

// preorder fills res with Preorder traversal results.
func preorder(cur int, tree map[int]Node, res *[]int) {
	if cur == -1 {
		return
	}
	node, ok := tree[cur]
	if !ok {
		return
	}
	*res = append(*res, cur)
	preorder(node.left, tree, res)
	preorder(node.right, tree, res)
}

// inorder fills res with Inorder traversal results.
func inorder(cur int, tree map[int]Node, res *[]int) {
	if cur == -1 {
		return
	}
	node, ok := tree[cur]
	if !ok {
		return
	}
	inorder(node.left, tree, res)
	*res = append(*res, cur)
	inorder(node.right, tree, res)
}

// postorder fills res with Postorder traversal results.
func postorder(cur int, tree map[int]Node, res *[]int) {
	if cur == -1 {
		return
	}
	node, ok := tree[cur]
	if !ok {
		return
	}
	postorder(node.left, tree, res)
	postorder(node.right, tree, res)
	*res = append(*res, cur)
}

// printTraversal prints the slice of ints in the required format.
func printTraversal(label string, nums []int) {
	fmt.Print(label + ": ")
	for i, v := range nums {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
