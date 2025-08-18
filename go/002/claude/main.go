package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// サイズ読み取り
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// 配列読み取り
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	elements := make([]int, n)
	for i, part := range parts {
		elements[i], _ = strconv.Atoi(part)
	}

	// ターゲット読み取り
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())

	// Two Pointer法で解決
	if isPairExists(elements, target) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isPairExists(elements []int, target int) bool {
	sort.Ints(elements)
	left, right := 0, len(elements)-1

	for left < right {
		sum := elements[left] + elements[right]
		if sum == target {
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}
