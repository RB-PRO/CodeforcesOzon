package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

// Тест №1
func TestMain(t *testing.T) {
	// Ввод
	question := `3
6
2 1 3 1 1 4
2
5 5
8
1 4 2 5 4 2 6 3
`
	// Ожидаемый вывод
	answer := `1 2
3 6
4 5

1 2

1 3
2 5
4 7
6 8
`

	// Записываем в канал Stdin информацию входную
	inputBuf := bufio.NewReader(strings.NewReader(question))

	// Запуск тестируемой функции
	outPutWork := Work(inputBuf)

	// Проверка на ошибку
	if outPutWork != answer {
		t.Fatal("\nOUTPUT:\n" + fmt.Sprint(outPutWork) + "\nTrue:\n" + answer)
	}
}

func TestMinResp(t *testing.T) {
	var inputMapTestMinResp [3]map[int]int
	inputMapTestMinResp[0] = map[int]int{1: 2, 2: 1, 3: 3, 4: 1, 5: 1, 6: 4}
	inputMapTestMinResp[1] = map[int]int{3: 3, 4: 1, 5: 1, 6: 4}
	inputMapTestMinResp[2] = map[int]int{4: 1, 5: 1}

	answer := [3]int{2, 6, 5}

	for index, mapas := range inputMapTestMinResp {
		responseFunc := MinResp(&mapas)
		if responseFunc != answer[index] {
			t.Fatal(mapas, "Ответ", responseFunc, "True", answer[index])
		}
	}
}
