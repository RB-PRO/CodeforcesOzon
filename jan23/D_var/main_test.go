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
	question := `6
4
20 10 20 30
3
5 7 6
5
6 3 4 3 1
5
200 10 100 11 200
1
1000000000
11
13 8 12 1 7 10 1 8 10 2 17
`
	// Ожидаемый вывод
	answer := `2 1 2 4 
1 1 1 
5 2 2 2 1 
4 1 3 1 4 
1 
9 4 9 1 4 7 1 4 7 1 11 
`

	// Записываем в канал Stdin информацию входную
	inputBuf := bufio.NewReader(strings.NewReader(question))

	// Запуск тестируемой функции
	outPutWork := Work(inputBuf)

	// Проверка на ошибку
	if outPutWork != answer {
		t.Fatal("\nloss:\n"+fmt.Sprint(outPutWork), "\nanswer\n"+answer+"\n")
	}
}
