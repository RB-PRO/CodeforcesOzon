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
8
Booble
yyyess
oooops
oooooopppss
Boooble
yyessss
oooops
ooooppssss
6
a
aa
MMM
mmm
yess
yyes
5
rrrrrrrr
rrrrrr
rrr
rr
r
`
	// Ожидаемый вывод
	answer := `4
6
2`

	// Записываем в канал Stdin информацию входную
	inputBuf := bufio.NewReader(strings.NewReader(question))

	// Запуск тестируемой функции
	outPutWork := Work(inputBuf)

	// Проверка на ошибку
	if outPutWork != answer {
		t.Fatal("\nloss:\n", fmt.Sprint(outPutWork)+"\nanswer\n"+answer+"\n")
	}
}
