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

	inputBuf := bufio.NewReader(strings.NewReader(question))

	outPutWork := Work(inputBuf) // Запуск тестируемой функции

	if outPutWork != answer { // Проверка на ошибку
		t.Fatal("\nloss:", fmt.Sprint(outPutWork)+"\n")
	}
}
