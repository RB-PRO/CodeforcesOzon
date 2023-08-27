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
	question := `8 10
1 2
1 3
1 4
4 3
3 2
2 4
1 8
5 6
7 6
5 7
`
	// Ожидаемый вывод
	answer := `0
8
8
8
0
0
0
2 3 4
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

func TestCoutFriends(t *testing.T) {
	fmt.Println(CoutFriends([]int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}))
	fmt.Println(CoutFriends([]int{1, 2, 3, 4, 5, 77}, []int{4, 5, 6, 7, 8}))
}
