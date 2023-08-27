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
		var m int
		fmt.Fscan(in, &m)
		strs := make([]string, m)

		for j := 0; j < m; j++ {
			fmt.Fscan(in, &strs[j])
		}
		outputStr += strconv.Itoa(TovarZnak(strs)) + "\n"
	}

	return outputStr
}

func TovarZnak(strs []string) int {
	for indX := range strs {
		//var tecStr string = strs[indX]
		for indY := range strs {
			if indY != indX { // чтобы не попасть на туже самуж строку

			}
		}
	}
	return len(strs) + 1
}
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
