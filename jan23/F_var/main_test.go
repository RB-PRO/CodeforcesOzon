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
	question := `7
8
7
8
1,7,1
8
1-5,1,7-7
10
1-5
10
1,2,3,4,5,6,8,9,10
3
1-2
100
1-2,3-7,10-20,100
`
	// Ожидаемый вывод
	answer := `8,1-6
2-6,8
6,8
6-10
7
3
9,21-99,8`

	// Записываем в канал Stdin информацию входную
	inputBuf := bufio.NewReader(strings.NewReader(question))

	// Запуск тестируемой функции
	outPutWork := Work(inputBuf)

	// Проверка на ошибку
	if outPutWork != answer {
		t.Fatal("\nloss:\n", fmt.Sprint(outPutWork)+"\nanswer\n"+answer+"\n")
	}
}
