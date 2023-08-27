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
R48FAO00OOO0OOA99OKA99OK
R48FAO00OOO0OOA99OKA99O
A9PQ
A9PQA
A99AAA99AAA99AAA99AA
AP9QA
`
	// Ожидаемый вывод
	answer := `R48FA O00OO O0OO A99OK A99OK
-
A9PQ
-
A99AA A99AA A99AA A99AA
-
`

	// Записываем в канал Stdin информацию входную
	inputBuf := bufio.NewReader(strings.NewReader(question))

	// Запуск тестируемой функции
	outPutWork := Work(inputBuf)

	// Проверка на ошибку
	if outPutWork != answer {
		t.Fatal("\nloss:\n", fmt.Sprint(outPutWork), "answer\n", answer+"\n")
	}
}

func TestNumberAvto(t *testing.T) {
	// Ввод
	question := "R48FAO00OOO0OOA99OKA99OK"
	// Ожидаемый вывод
	//answer := `R48FA O00OO O0OO A99OK A99OK`
	answer := make([]string, 5)
	answer[0] = "R48FA"
	answer[1] = "O00OO"
	answer[2] = "O0OO"
	answer[3] = "A99OK"
	answer[4] = "A99OK"

	// Запуск тестируемой функции
	outPutWork := NumberAvto(question)

	// Проверка на ошибку

	if len(outPutWork) != len(answer) {
		t.Fatal("\nРазные длины массивов", len(outPutWork), len(answer), "\n")
	}
	for index := range answer {
		if !strings.Contains(strings.TrimSpace(answer[index]), strings.TrimSpace(outPutWork[index])) {
			t.Fatal("\nloss:", answer[index]+" - "+outPutWork[index]+"\n")
		}
	}
}

func TestIsNumberOne(t *testing.T) { // Тест №2 для первого типа номеров
	// Ввод
	question := `R48FA`

	// Проверка на ошибку
	if !IsNumberOne(question) {
		t.Fatal("\nloss:", question+"\n")
	}
}

func TestIsNumberTwo(t *testing.T) { // Тест №2 для первого типа номеров
	// Ввод
	question := `T7RR`

	// Проверка на ошибку
	if !IsNumberTwo(question) {
		t.Fatal("\nloss:", question+"\n")
	}
}

func TestIsLetter(t *testing.T) { // Тест №3 для первого типа номеров
	// Проверка на ошибку
	if !IsLetter('a') {
		t.Fatal("\nloss: a\n")
	}
	if !IsLetter('A') {
		t.Fatal("\nloss: A\n")
	}
	if !IsLetter('F') {
		t.Fatal("\nloss: A\n")
	}
	if IsLetter('1') {
		t.Fatal("\nloss: 1\n")
	}
}
func TestIsFigure(tFig *testing.T) {
	if !IsFigure('1') {
		tFig.Fatal("\nloss: 1\n")
	}
	if !IsFigure('0') {
		tFig.Fatal("\nloss: 0\n")
	}
	if IsFigure('a') {
		tFig.Fatal("\nloss: a\n")
	}
	if IsFigure('A') {
		tFig.Fatal("\nloss: A\n")
	}
}
