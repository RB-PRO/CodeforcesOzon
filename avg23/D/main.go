package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() { // Главная функция
	inputBuf := bufio.NewReader(os.Stdin)   // Считывание
	outputBuf := bufio.NewWriter(os.Stdout) // Вывод
	fmt.Fprintln(outputBuf, Work(inputBuf)) // Записываем в вывод
	outputBuf.Flush()                       // Опустошаем вывод
}

// Решение
func Work(in *bufio.Reader) (outputStr string) {

	var testCount int
	fmt.Fscan(in, &testCount)
	// Цикл по всем тестам
	for i := 0; i < testCount; i++ {

		var n, m int
		fmt.Fscan(in, &n)
		fmt.Fscan(in, &m)

		board := make([]int, 0, m*n) // board[i*m+j] = "abc" // like board[i][j] = "abc"
		for j := 0; j < m*n; j++ {
			var InpputData int
			fmt.Fscan(in, &InpputData)
			board = append(board, InpputData)
		}

		var coutWork int
		fmt.Fscan(in, &coutWork)
		workingMode := make([]int, 0, coutWork)
		for j := 0; j < coutWork; j++ {
			var InpputData int
			fmt.Fscan(in, &InpputData)
			workingMode = append(workingMode, InpputData)
		}

		working(board, m, workingMode)
		fmt.Println()
	}

	return outputStr
}

func working(arr []int, m int, workingMode []int) (out string) {
	// PrintArr(arr, m)
	// fmt.Println()

	// // fmt.Println()
	// PrintArr(SwapLine(arr, m, 0, 1), m)
	// fmt.Println()
	// PrintArr(RefreshSlice(arr, m, []int{2, 0, 1, 3}), m)

	for _, mode := range workingMode {
		data, numbers := collumn(arr, m, mode)
		// fmt.Println("data,numbers -", data, numbers)

		// Сортировка
		var i = 1
		for i < len(data) {
			var j = i
			for j >= 1 && data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
				j--
			}
			i++
		}

		// fmt.Println("data,numbers >", data, numbers)

		arr = RefreshSlice(arr, m, numbers) // пересобрать слайс

		//
	}
	PrintArr(arr, m) // Вывести полученную матрицу

	return out
}

////////////////////////////

// печать матрицы
func PrintArr(arr []int, m int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		if (i+1)%m == 0 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}

// Получить список колонок и порядковые номера
func collumn(arr []int, m int, ColCout int) (data []int, numbers []int) {

	// Порядковые номера колонок с нуля
	var n int = len(arr) / m
	for i := 0; i < n; i++ {
		numbers = append(numbers, i)
	}

	// колонки
	for i := 0; i < n; i++ {
		data = append(data, arr[i*m+ColCout-1])
	}

	return data, numbers
}

// Поменять местами строки в массиве
func SwapLine(arr []int, m int, a, b int) []int {
	arrA := make([]int, m)
	arrB := make([]int, m)
	copy(arrA, arr[a*m:(a+1)*m])
	copy(arrB, arr[b*m:(b+1)*m])
	// fmt.Println("arrA", arrA, "arrB", arrB)
	for i := 0; i < len(arrA); i++ {
		arr[a*m+i] = arrB[i]
	}
	for i := 0; i < len(arrB); i++ {
		arr[b*m+i] = arrA[i]
	}
	return arr
}

// функция принимает представление двумерного слайса и необходимый строковый порядок, который необходимо соблюсти.
// На выходе думерный слайс с порядком, в соответствии со строковым порядком
func RefreshSlice(arr []int, m int, numbers []int) (OutPut []int) {

	// Создать мапу, которая уже будет содержать
	// в значениях - номер строки из исходного слайса
	// а в индексах - номера строк в полученном слайсе
	// Ну оно бы было мапой, если бы в мапе все значения в индексах были бы по порядку! а пока что слайс
	// NumbersMap := make(map[int]int)
	// for i, v := range numbers {
	// 	NumbersMap[v] = i
	// }
	// NumbersMapSlise := make([]int, 0, len(arr)/m)
	// for i, v := range NumbersMap {
	// 	fmt.Println(i, v)
	// 	NumbersMapSlise[i] = v
	// }

	NumbersMapSlise := make([]int, len(arr)/m, len(arr)/m)
	for i, v := range numbers {
		NumbersMapSlise[i] = v
	}
	// fmt.Println(NumbersMapSlise)
	for _, v := range NumbersMapSlise {
		OutPut = append(OutPut, arr[v*m:(v+1)*m]...)
	}

	return OutPut
}
