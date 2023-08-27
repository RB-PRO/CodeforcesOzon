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
256 42
1000 1000
-1000 1000
-1000 1000
20 22
`
	// Ожидаемый вывод
	answer := `298
2000
0
0
42
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
