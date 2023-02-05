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
		fmt.Fscan(in, &n) // количество разработчиков
		//devs := make(map[int]int) // Данные будем держать в мапе
		devs := make([]Dev, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &dataInt) // Прочитать мастерство разработчика
			devs[j].number = j + 1
			devs[j].rep = dataInt
		}

		// Сохранение результата
		outputStr += FindCoast(devs) + "\n"
	}

	return outputStr
}

// Поиск пар-ключей
func FindCoast(devs []Dev) string {
	var OutPutStr string

	//for indexDev := range devs {
	for indexDev := 0; indexDev < len(devs)/2; indexDev++ {
		fmt.Println(indexDev, len(devs)/2)
		if len(devs) < 2 {
			break
		}
		nextMinIndexKey := MinResp(devs)
		OutPutStr += fmt.Sprintf("%d %d\n", indexDev, nextMinIndexKey)

		fmt.Println(devs, "nextMinIndexKey", nextMinIndexKey-1)
		devs = removeDev(devs, nextMinIndexKey-1)
		devs = removeDev(devs, 0)

	}
	return OutPutStr
}

// Получить номер элемента массива, который можем считать минимальным и стоит после текущего разраба(т.е со второго разраба)
// Возвращает положение напарника
func MinResp(devs []Dev) int {
	var startInt, min, counter, indexKey int = 0, 0, 0, 0
	for _, val := range devs {
		if counter == 0 {
			startInt = val.rep
		} else if counter == 1 {
			indexKey = val.number
			min = abs(startInt - val.rep)
		} else {
			if abs(startInt-val.rep) < min {
				min = abs(startInt - val.rep)
				indexKey = val.number
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
func removeDev(slice []Dev, s int) []Dev {
	return append(slice[:s], slice[s+1:]...)
}

type Dev struct {
	number int // Порядковый номер сотрудника
	rep    int // Репутация сотрудника
}
