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
func Work(in *bufio.Reader) string {
	var testCount int
	var outputStr string

	fmt.Fscan(in, &testCount) // количество наборов входных данных

	for i := 0; i < testCount; i++ {
		var n, dataInt int
		fmt.Fscan(in, &n)         // количество разработчиков
		devs := make(map[int]int) // Данные будем держать в мапе
		for j := 1; j < n; j++ {
			fmt.Fscan(in, &dataInt) // Прочитать мастерство разработчика
			devs[j] = dataInt
		}

		// Сохранение результата
		outputStr += FindCoast(devs) + "\n"
	}

	return outputStr
}

// Поиск пар-ключей
func FindCoast(devs map[int]int) string {
	var OutPutStr string

	for key := range devs {
		if len(devs) < 2 {
			break
		}
		nextMinIndexKey := MinResp(&devs)
		OutPutStr += fmt.Sprintf("%d %d\n", key, nextMinIndexKey)

		delete(devs, nextMinIndexKey)
		delete(devs, key)

	}
	return OutPutStr
}

// Получить номер элемента массива, который можем считать минимальным и стоит после текущего разраба(т.е со второго разраба)
// Возвращает положение напарника
func MinResp(devs *map[int]int) int {
	var startInt, min, counter, indexKey int = 0, 0, 0, 0
	for key, val := range *devs {
		if counter == 0 {
			startInt = val
		} else if counter == 1 {
			indexKey = key
			min = abs(startInt - val)
		} else {
			if abs(startInt-val) < min {
				min = abs(startInt - val)
				indexKey = key
			}
		}
		//fmt.Printf("counter %#v. indexKey %#v. min %#v. startInt %#v. key %#v. val %#v. abs(startInt-val) %#v-%#v. abs %#v\n", counter, indexKey, min, startInt, key, val, startInt, val, abs(startInt-val))
		counter++
	}
	return indexKey
}

// Модуль числа
func abs(inputInt int) int {
	if inputInt < 0 {
		return inputInt * (-1)
	} else {
		return inputInt
	}
}
