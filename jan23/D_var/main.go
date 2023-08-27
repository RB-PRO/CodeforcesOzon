package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	for i := 0; i < testCount; i++ { // Цикл по набору данных
		var m int
		// Берём массив данных в массив интов
		fmt.Fscan(in, &m)
		t := make([]int, m)
		for j := 0; j < m; j++ {
			var value int
			fmt.Fscan(in, &value)
			t[j] = value
		}

		// Получаем массив мест
		outputStr += Ranges(t) + "\n"
	}
	return outputStr
}

// ***
type rez struct {
	t     int // Время
	pos   int // Позиция в исходном массиве
	mesto int // Место на пьедестале
}

func Ranges(t []int) string {
	// Получить исходную структуру, с которой и буде работать
	rezult := make([]rez, len(t))
	for ind := range t {
		rezult[ind].t = t[ind]
		rezult[ind].pos = ind
	}

	// Сортируем
	sort.Slice(rezult, func(i, j int) (less bool) {
		return rezult[i].t < rezult[j].t
	})

	// Обработка мест
	var p int
	for ind := range rezult {
		rezult[ind].mesto = ind + 1
		if ind > 0 {
			// Если отличие по времени не более 1 секунды
			if abs(rezult[ind].t-rezult[ind-1].t) <= 1 {
				rezult[ind].mesto = rezult[ind-1].mesto
				p++
			} else if p != 0 {
				rezult[ind].mesto = ind + 1
				p = 0
			}

		}
	}

	// Вывод результатов
	outputStr := make([]string, len(t))
	for ind := range rezult {
		outputStr[rezult[ind].pos] = strconv.Itoa(rezult[ind].mesto)
	}
	return strings.Join(outputStr, " ")
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	} else {
		return i
	}
}
