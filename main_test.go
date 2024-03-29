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
	question := `2
21
[
	{
		"id":0,
		"name":"all"
	},
	{
		"id":1,
		"name":"clothes",
		"parent":0
	},
	{
		"id":2,
		"name":"shoes",
		"parent":0
	},
	{
		"id":55,
		"name":"sneakers",
		"parent":2
	}
]
6
	[ {"parent":	0,"id":100,  "name":
	"x"},{

"name":"x","id":0}

]
`
	// Ожидаемый вывод
	answer := `[{
	"id": 0,
	"name": "all",
	"next": [{
		"id": 1,
		"name": "clothes",
		"next": []
	}, {
		"id": 2,
		"name": "shoes",
		"next": [{
			"id": 55,
			"name": "sneakers"
		}]
	}]
},
{"name":"x","id":0,"next":[{"id":100,"name":"x"}]}
]
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
