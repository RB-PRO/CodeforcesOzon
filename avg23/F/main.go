package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() { // Главная функция
	inputBuf := bufio.NewReader(os.Stdin)   // Считывание
	outputBuf := bufio.NewWriter(os.Stdout) // Вывод
	fmt.Fprintln(outputBuf, Work(inputBuf)) // Записываем в вывод
	outputBuf.Flush()                       // Опустошаем вывод
}

type tt struct {
	ss int
	h  int
	m  int
	s  int
}
type ttt struct {
	t1 tt
	t2 tt
}

// Решение
func Work(in *bufio.Reader) (outputStr string) {

	var t int
	fmt.Fscan(in, &t)
	// Цикл по всем тестам
	for i := 0; i < t; i++ {

		var n int
		fmt.Fscan(in, &n)

		// board := make([]int, 0, n)
		// for j := 0; j < n; j++ {
		// 	var InpputData int
		// 	fmt.Fscan(in, &InpputData)
		// 	board = append(board, InpputData)
		// }
		// fmt.Println(working(board, n))

		// strs := make([]string, 0, n)
		tts := make([]ttt, n)
		for j := 0; j < n; j++ {
			var InpputData string
			fmt.Fscan(in, &InpputData)
			var src_buff bytes.Buffer
			fmt.Fprint(&src_buff, InpputData)
			fmt.Fscanf(&src_buff, "%d:%d:%d-%d:%d:%d", &tts[j].t1.h, &tts[j].t1.m, &tts[j].t1.s, &tts[j].t2.h, &tts[j].t2.m, &tts[j].t2.s)
		}
		// fmt.Println(tts)
		fmt.Println(work(tts))
	}

	return outputStr
}
func work(tts []ttt) string {
	for i := range tts {
		if tts[i].t1.h >= 24 || tts[i].t2.h >= 24 {
			return "NO"
		}
		if tts[i].t1.m >= 60 || tts[i].t2.m >= 60 {
			return "NO"
		}
		if tts[i].t1.s >= 60 || tts[i].t2.s >= 60 {
			return "NO"
		}

		tts[i].t1.ss = 10000*tts[i].t1.h + 100*tts[i].t1.m + tts[i].t1.s
		tts[i].t2.ss = 10000*tts[i].t2.h + 100*tts[i].t2.m + tts[i].t2.s

		// проверка что выполняется требование "от-до"
		if tts[i].t1.ss > tts[i].t2.ss {
			return "NO"
		}
	}
	if len(tts) == 1 {
		return "YES"
	}

	// Сортируем
	sort.Slice(tts, func(i, j int) bool {
		return tts[i].t1.ss < tts[j].t1.ss
	})

	// Локика такая. Если начало последующего не ближе текущего, то блокируем
	for i := 0; i < len(tts)-1; i++ {
		if tts[i].t2.ss >= tts[i+1].t1.ss {
			return "NO"
		}
	}

	return "YES"
}

// func working(arr []int, n int) string {
// 	m := make(map[int]int)
// 	var back int
// 	for _, v := range arr {
// 		// fmt.Println(i, arr, v, back)
// 		if _, ok := m[v]; ok {
// 			if back != v {
// 				return "NO"
// 			}
// 		} else {
// 			m[v] = 0
// 		}
// 		back = v
// 	}
// 	// fmt.Println(m)
// 	return "YES"
// }
