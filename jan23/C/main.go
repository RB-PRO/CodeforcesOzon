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

	fmt.Fscan(in, &testCount) // количество наборов входных данных

	for i := 0; i < testCount; i++ {
		var n, dataInt int
		fmt.Fscan(in, &n) // количество разработчиков
		var arrInt []int
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &dataInt) // Прочитать мастерство разработчика
			arrInt = append(arrInt, dataInt)
		}

		// Сохранение результата
		outputStr += FindCoast(arrInt) + "\n"
	}

	return outputStr
}

// Поиск пар-ключей
func FindCoast(inp []int) string {

	return ""
}
