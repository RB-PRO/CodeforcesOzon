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
1
02:46:00-03:14:59
2
23:59:59-23:59:59
00:00:00-23:59:58
2
23:59:58-23:59:59
00:00:00-23:59:58
2
23:59:59-23:59:58
00:00:00-23:59:57
6
17:53:39-20:20:02
10:39:17-11:00:52
08:42:47-09:02:14
09:44:26-10:21:41
00:46:17-02:07:19
22:42:50-23:17:46
1
24:00:00-23:59:59
`
	// Ожидаемый вывод
	answer := `YES
YES
NO
NO
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
