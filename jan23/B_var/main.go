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

/*
	четыре однопалубных корабля,
	три двухпалубных корабля,
	два трёхпалубных корабля,
	один четырёхпалубный корабль.
*/

// Решение
func Work(in *bufio.Reader) (outputStr string) {
	var testCount int

	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ { // Цикл по набору данных
		ship := make(map[int]int)
		for j := 0; j < 10; j++ {
			// буду делать через мапу и проверять в ней целостность данных. По тем критериям входных данных это приемлемо

			var m int
			fmt.Fscan(in, &m)
			ship[m]++
		}
		if isCorrectShips(ship) {
			outputStr += "YES\n"
		} else {
			outputStr += "NO\n"
		}
	}

	return outputStr
}

func isCorrectShips(ship map[int]int) bool {
	if ship[4] != 1 {
		return false
	}
	if ship[3] != 2 {
		return false
	}
	if ship[2] != 3 {
		return false
	}
	if ship[1] != 4 {
		return false
	}
	return true
}
