package main

import (
	"bufio"
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

// Решение
func Work(in *bufio.Reader) (outputStr string) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	users := make(map[int][]int)

	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		var b int
		fmt.Fscan(in, &b)

		users[a] = append(users[a], b)
		users[b] = append(users[b], a)
	}

	var outs = make([][]int, n)
	for k, v := range users { // все друзья
		m := make(map[int]int, 0) // потенциальные друзья
		for _, vv := range v {    // друзья
			for _, vvv := range users[vv] { // друзья друзей
				if vvv != k && !containsElement(v, vvv) {
					m[vvv]++
				}
			}
		}
		fmt.Println(k, m)

		// cout,outs:=CoutFriends()
		var max int
		var out []int
		for _, vv := range m {
			if vv > max {
				max = vv
			}
		}
		for kk, vv := range m {
			if vv == max {
				out = append(out, kk)
			}
		}
		sort.Ints(out)
		outs[k-1] = out
	}

	for _, v := range outs {
		for _, vv := range v {
			fmt.Print(vv)
			fmt.Print(" ")
		}
		if len(v) == 0 {
			fmt.Print(0)
		}
		fmt.Println()
	}
	return outputStr
}

// найти к-во обзих друзей между двумя юзерами
func CoutFriends(a, b []int) (cout int, ret []int) {
	m := make(map[int]int)
	var max int
	for _, v := range a {
		m[v]++
		if m[v] > max {
			max = m[v]
		}
	}
	for k, v := range m {
		if v == max {
			ret = append(ret, k)
		}
	}
	for _, v := range b {
		if _, ok := m[v]; ok {
			cout++
		}

	}
	return cout, ret
}
func containsElement(slice []int, element int) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}
