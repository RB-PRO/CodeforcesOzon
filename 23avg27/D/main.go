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
	var t int
	fmt.Fscan(input, &t)
	for ii := 0; ii < t; ii++ {
		var k, n, m int
		fmt.Fscan(input, &k)
		fmt.Fscan(input, &n)
		fmt.Fscan(input, &m)

		// fmt.Println("k, n, m", k, n, m)
		rr := make([][]rune, n)
		for i := range rr {
			rr[i] = make([]rune, m)
		}

		for ik := 0; ik < k; ik++ {
			for in := 0; in < n; in++ {
				var str string
				fmt.Fscan(input, &str)
				strRune := []rune(str)

				// fmt.Println(ik, in, str)

				for im := 0; im < m; im++ {

					// rr[in][im] 0
					// rr[in][im] 46 .
					// rr[in][im] 47 /
					// rr[in][im] 88 X
					// rr[in][im] 92 \
					if (rr[in][im]) == 0 || rr[in][im] == 46 {
						rr[in][im] = strRune[im]
					}

				}
			}
		}

		for in := 0; in < n; in++ {
			outputStr += string(rr[in]) + "\n"
		}
		outputStr += "\n"
	}
	return outputStr
}
