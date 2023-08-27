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

	var testCount int
	fmt.Fscan(in, &testCount)
	// Цикл по всем тестам
	for i := 0; i < testCount; i++ {

		var CoutData int
		fmt.Fscan(in, &CoutData)

		var coast []int
		for j := 0; j < CoutData; j++ {
			var InpputData int
			fmt.Fscan(in, &InpputData)

			coast = append(coast, InpputData)
		}

		fmt.Println(CalcMinSub(coast))
	}

	return outputStr
}

func CalcMinSub(s []int) (price int) {
	prices := make(map[int]int)
	for i := 0; i < len(s); i++ {
		prices[s[i]]++
	}
	for k, v := range prices {
		price += int(2*int(v/3)+int(v%3)) * k
	}

	return price
}
