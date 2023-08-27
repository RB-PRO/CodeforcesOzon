package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() { // Главная функция
	inputBuf := bufio.NewReader(os.Stdin)   // Считывание
	outputBuf := bufio.NewWriter(os.Stdout) // Вывод
	fmt.Fprintln(outputBuf, Work(inputBuf)) // Записываем в вывод
	outputBuf.Flush()                       // Опустошаем вывод
}

// Решение
func Work(in *bufio.Reader) (outputStr string) {
	var testCount int

	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)

		outputStr += strconv.Itoa(n+m) + "\n"
	}

	return outputStr
}
