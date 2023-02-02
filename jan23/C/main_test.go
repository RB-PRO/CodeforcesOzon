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
		t.Fatal("\nOUTPUT:\n", fmt.Sprint(outPutWork)+"\n")
	}
}
