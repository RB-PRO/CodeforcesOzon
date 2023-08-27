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
task
decide
id
6
flask
code
void
forces
id
ask
`
	// Ожидаемый вывод
	answer := `task
decide
id
task
decide
task
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
