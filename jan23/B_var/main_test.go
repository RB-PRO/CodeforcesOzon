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
	question := `5
2 1 3 1 2 3 1 1 4 2
1 1 1 2 2 2 3 3 3 4
1 1 1 1 2 2 2 3 3 4
4 3 3 2 2 2 1 1 1 1
4 4 4 4 4 4 4 4 4 4
`
	// Ожидаемый вывод
	answer := `YES
NO
YES
YES
NO
`

	// Записываем в канал Stdin информацию входную
	inputBuf := bufio.NewReader(strings.NewReader(question))

	// Запуск тестируемой функции
	outPutWork := Work(inputBuf)

	// Проверка на ошибку
	if outPutWork != answer {
		t.Fatal("\nloss:", fmt.Sprint(outPutWork)+"\n")
	}
}
