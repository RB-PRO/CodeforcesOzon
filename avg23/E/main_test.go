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
5
1 2 3 4 5
4
1 2 3 1
8
2 3 4 8 5 5 5 5
5
1 1 3 2 2
5
1 1 2 3 2
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
	if !strings.Contains(outPutWork, answer) {
		t.Fatal("\nloss:", fmt.Sprint(outPutWork)+"\n")
	}
}
