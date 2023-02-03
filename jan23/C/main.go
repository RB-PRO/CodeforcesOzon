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
func FindCoast(devs map[int]int) (OutPutStr string) {
	//fmt.Println(devs)

	for key := range devs {
		if len(devs) < 2 {
			break
		}
		nextMinIndexKey := MinResp(devs)
		OutPutStr += fmt.Sprintf("%d %d\n", key, nextMinIndexKey)

		fmt.Println(key, nextMinIndexKey, devs)

		delete(devs, nextMinIndexKey)
		delete(devs, key)

	}
	return OutPutStr
}

// Получить номер элемента массива, который можем считать минимальным и стоит после текущего разраба(т.е со второго разраба)
// Возвращает положение напарника
func MinResp(devs map[int]int) (indexKey int) {
	fmt.Println(devs)
	var startInt, min, counter int
	for key, val := range devs {
		if counter == 0 {
			startInt = val
		} else if counter == 1 {
			indexKey = key
			min = val
		} else {
			if abs(startInt-val) < min {
				min = abs(startInt - val)
				indexKey = key
			}
		}
		//fmt.Printf("counter %#v. index %#v. min %#v. startInt %#v. key %#v. abs(startInt-val) %#v.\n", counter, index, min, startInt, key, abs(startInt-val))
		counter++
	}
	//fmt.Println()
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
