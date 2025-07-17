package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// 1つ目の引数を回文チェック用の文字列とする
	str := os.Args[1]

	// 文字チェック
	if err := checkStrings([]rune(str)); err != nil {
		fmt.Printf("文字列チェックエラー: %v\n", err)
		return
	}

	// 回文用入力文字列を格納
	palindromeOrg := strings.Split(str, "")

	// 回文用入力文字列を反転して格納
	palindromeReverse := reverseStrings(palindromeOrg)

	// 回文チェック
	for i, v := range palindromeOrg {
		if v != palindromeReverse[i] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

func checkStrings(s []rune) error {
	// 桁数チェック
	if len(s) > 100 {
		return errors.New("入力文字数が100桁を超えています")
	}

	for _, v := range s {
		// 文字列チェック（大文字か）
		if unicode.IsUpper(v) {
			return errors.New("大文字が含まれてます")
		}

		// 文字列チェック（数字が含まれるか）
		if unicode.IsDigit(v) {
			return errors.New("数字が含まれてます")
		}
	}
	return nil
}

func reverseStrings(s []string) []string {
	n := len(s)
	reversed := make([]string, n)
	for i, char := range s {
		reversed[n-1-i] = char
	}
	return reversed
}
