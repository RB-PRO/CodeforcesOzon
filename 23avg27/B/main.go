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
	var s string
	fmt.Fscan(in, &s)
	sr := []rune(s)

	var n int
	fmt.Fscan(in, &n)
	for ii := 0; ii < n; ii++ {
		var start, end int
		var r string
		fmt.Fscan(in, &start)
		fmt.Fscan(in, &end)
		fmt.Fscan(in, &r)
		// fmt.Println(start, end, r)
		// for i := start; i < end; i++ {
		// 	s[i] = r[i]
		// }
		var srNew []rune
		rr := []rune(r)

		// fmt.Println("-", string(sr[:start-1]), string(sr[end:]))
		srNew = append(sr[:start-1], rr...) //sr[end:]... )
		srNew = append(srNew, sr[end:]...)

		sr = srNew
		// fmt.Println(string(sr))
	}

	return string(sr)
}
