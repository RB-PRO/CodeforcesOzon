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
	var sl int
	fmt.Fscan(in, &sl)
	s := make([]string, sl)
	for i := 0; i < sl; i++ {
		var str string
		fmt.Fscan(in, &str)
		if str != "" {
			s[i] = str
		}
	}

	var tl int
	fmt.Fscan(in, &tl)
	t := make([]string, tl)
	for i := 0; i < tl; i++ {
		var str string
		fmt.Fscan(in, &str)
		if str != "" {
			t[i] = str
		}
	}

	for _, tt := range t {
		var max int
		var maxS string
		for _, ss := range s {
			// tt - слово
			// ss - словарь
			var tecal int
			for i := 0; i < len(ss) && i < len(tt) && tt != ss; i++ {
				if tt[len(tt)-1-i] != ss[len(ss)-1-i] {
					break
				}
				tecal++
			}
			if tecal > max {
				max = tecal
				maxS = ss
			}
		}
		// fmt.Println(tt, maxS, max)
		if maxS == "" {
			maxS = s[0]
		}
		fmt.Println(maxS)
	}

	return outputStr
}
