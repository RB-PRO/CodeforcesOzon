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

4 3
3 4 1
2 2 5
2 4 2
2 2 1
3
2 1 3

3 1
100
9
10
2
1 1

3 3
2 11 72
99 11 13
2 8 13
5
2 3 2 1 2
`
	// Ожидаемый вывод
	answer := `2 2 1
3 4 1
2 4 2
2 2 5

9
10
100

2 8 13
2 11 72
99 11 13

`

	// Записываем в канал Stdin информацию входную
	inputBuf := bufio.NewReader(strings.NewReader(question))

	// Запуск тестируемой функции
	outPutWork := Work(inputBuf)

	// Проверка на ошибку
	if !strings.Contains(outPutWork, answer) {
		t.Fatal("\nloss:", fmt.Sprint(outPutWork)+"\n")
	}
}
