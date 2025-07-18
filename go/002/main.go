package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter various conditions")

	// Scan input text 3-lines
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for range 3 {
		scanner.Scan()
		lines = append(lines, scanner.Text())
	}

	fmt.Println("The rows read are")
	for i, line := range lines {
		fmt.Printf("%d: %s\n", i+1, line)
	}

	// Check various conditions
	size, err := checkSizeNumber(lines[0])
	if err != nil {
		fmt.Println("line1 input error: ", err)
		return
	}
	elements, err := checkArrayElements(lines[1], size)
	if err != nil {
		fmt.Println("line2 input error: ", err)
		return
	}
	target, err := checkTargetNumber(lines[2])
	if err != nil {
		fmt.Println("line3 input error: ", err)
		return
	}

	// Check to see if the pair exists
	if isPairExists(elements, target) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func checkSizeNumber(str string) (int, error) {
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New("invalid format")
	}

	if n < 2 || n > 1000 {
		return 0, errors.New("out of range value (2 to 1000)")
	}

	return n, nil
}

func checkArrayElements(str string, size int) ([]int, error) {
	var result []int

	// Check array size
	sp := strings.Split(str, " ")
	if len(sp) != size {
		return nil, errors.New("array size is unmatched (non-prescribed character)")
	}

	// Check if numbers and non-white space are included
	for _, s := range sp {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.New("include non-numeric character")
		}

		if i < -1000 || i > 1000 {
			return nil, errors.New("out of range value (-1000 to 1000)")
		}

		result = append(result, i)
	}

	return result, nil
}

func checkTargetNumber(str string) (int, error) {
	t, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New("invalid target number")
	}

	if t < -2000 || t > 2000 {
		return 0, errors.New("out of range value (-2000 to 2000)")
	}

	return t, nil
}

func isPairExists(elements []int, target int) bool {
	sort.Ints(elements)
	// fmt.Println("sorted:", elements)
	for i := range len(elements) - 1 {
		for j := i + 1; j < len(elements); j++ {
			sum := elements[i] + elements[j]
			// fmt.Printf("sum: %d = %d + %d: \n", sum, elements[i], elements[j])
			if sum == target {
				return true
			}
		}
	}

	return false
}
