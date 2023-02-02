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
		fmt.Fscan(in, &n)         // количество разработчиков
		devs := make(map[int]int) // Данные будем держать в мапе
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &dataInt) // Прочитать мастерство разработчика
			devs[i] = dataInt
		}

		// Сохранение результата
		outputStr += FindCoast(devs) + "\n"
	}

	return outputStr
}

// Поиск пар-ключей
func FindCoast(devs map[int]int) string {

	return ""
}

// Получить номер элемента массива, который можем считать минимальным и стоит после текущего разраба(т.е со второго разраба)
// Возвращает положение напарника
func MinResp(devs map[int]int /**/) (index int) {
	startInt := devs[0]
	var min int = devs[0]
	//for i := 1; i < len(devs); i++ {
	for key, val := range devs {
		if abs(startInt-val) < min {
			min = abs(startInt - val)
			index = key
		}
	}
	return index + 1
}

// Модуль числа
func abs(inputInt int) int {
	if inputInt < 0 {
		return inputInt * (-1)
	} else {
		return inputInt
	}
}
