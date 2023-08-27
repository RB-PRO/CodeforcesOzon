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

func Work(input *bufio.Reader) (outputStr string) {
	var n, m int
	fmt.Fscan(input, &n)
	fmt.Fscan(input, &m)
	if m < n {
		return "-1"
	}
	a := make([]int, n)
	for ii := 0; ii < n; ii++ {
		fmt.Fscan(input, &a[ii])
	}

	b := make([]int, n)

	// карточки
	mm := make([]int, m)
	for i := range mm {
		mm[i] = i + 1
	}

	// fmt.Println(n, m, a, "-", b)
	var min int
	if m > 0 {
		min = mm[0]
	}
	for ia := range a {

		if min > a[ia] {
			b[ia] = min
			mm = mm[1:]
			if len(mm) > 0 {
				min = mm[0]
			}
		}

		for im := range mm {
			if mm[im] > a[ia] {
				b[ia] = mm[im]
				mm = remove(mm, im)
				if len(mm) > 0 {
					min = mm[0]
				}
				break
			}
		}
		if b[ia] == 0 {
			return "-1"
		}
	}

	for _, mv := range b {
		outputStr = fmt.Sprintf("%s %v", outputStr, mv)
	}

	return outputStr
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
