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
		t.Fatal("\nOUTPUT:\n" + fmt.Sprint(outPutWork) + "\nTrue:\n" + answer)
	}
}

func TestMinResp(t *testing.T) {
	// *** 2 1 3 1 1 4 ***
	inputMapTestMinResp := make([]Dev, 6)
	inputMapTestMinResp[0] = Dev{number: 1, rep: 2}
	inputMapTestMinResp[1] = Dev{number: 2, rep: 1}
	inputMapTestMinResp[2] = Dev{number: 3, rep: 3}
	inputMapTestMinResp[3] = Dev{number: 4, rep: 1}
	inputMapTestMinResp[4] = Dev{number: 5, rep: 1}
	inputMapTestMinResp[5] = Dev{number: 6, rep: 4}

	responseFunc := MinResp(inputMapTestMinResp)
	if responseFunc != 2 {
		t.Fatal("Ответ", responseFunc, "True", 2)
	}

	// *** 5 5 ***
	inputMapTestMinResp = make([]Dev, 4)
	inputMapTestMinResp[0] = Dev{number: 1, rep: 3}
	inputMapTestMinResp[1] = Dev{number: 2, rep: 1}
	inputMapTestMinResp[2] = Dev{number: 3, rep: 1}
	inputMapTestMinResp[3] = Dev{number: 4, rep: 4}

	responseFunc = MinResp(inputMapTestMinResp)
	if responseFunc != 4 {
		t.Fatal("Ответ", responseFunc, "True", 2)
	}

	// *** 5 5 ***
	inputMapTestMinResp = make([]Dev, 2)
	inputMapTestMinResp[0] = Dev{number: 2, rep: 1}
	inputMapTestMinResp[1] = Dev{number: 3, rep: 1}

	responseFunc = MinResp(inputMapTestMinResp)
	if responseFunc != 2 {
		t.Fatal("Ответ", responseFunc, "True", 2)
	}
}
