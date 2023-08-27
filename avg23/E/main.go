package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() { // Главная функция
	inputBuf := bufio.NewReader(os.Stdin)   // Считывание
	outputBuf := bufio.NewWriter(os.Stdout) // Вывод
	fmt.Fprintln(outputBuf, Work(inputBuf)) // Записываем в вывод
	outputBuf.Flush()                       // Опустошаем вывод
}

// Решение
func Work(in *bufio.Reader) (outputStr string) {

	var t int
	fmt.Fscan(in, &t)
	// Цикл по всем тестам
	for i := 0; i < t; i++ {

		var n int
		fmt.Fscan(in, &n)

		board := make([]int, 0, n)
		for j := 0; j < n; j++ {
			var InpputData int
			fmt.Fscan(in, &InpputData)
			board = append(board, InpputData)
		}

		fmt.Println(working(board, n))
	}

	return outputStr
}

func working(arr []int, n int) string {

	m := make(map[int]int)
	var back int
	for _, v := range arr {
		// fmt.Println(i, arr, v, back)
		if _, ok := m[v]; ok {
			if back != v {
				return "NO"
			}
		} else {
			m[v] = 0
		}
		back = v
	}
	// fmt.Println(m)
	return "YES"
}
