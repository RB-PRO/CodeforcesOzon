package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() { // Главная функция
	inputBuf := bufio.NewReader(os.Stdin)   // Считывание
	outputBuf := bufio.NewWriter(os.Stdout) // Вывод
	fmt.Fprintln(outputBuf, Work(inputBuf)) // Записываем в вывод
	outputBuf.Flush()                       // Опустошаем вывод
}

// Решение
func Work(in *bufio.Reader) (outputStr string) {
	var t int

	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ { // Цикл по набору данных
		var k int
		fmt.Fscan(in, &k)

		var str string
		fmt.Fscan(in, &str)

		outputStr += stranici(k, str) + "\n"
	}

	return outputStr
}

func stranici(k int, str string) string {
	var pages []int
	strs := strings.Split(str, ",")
	for ind := range strs {
		if strings.Contains(strs[ind], "-") { // Если по диапазону
			strsTires := strings.Split(strs[ind], "-")
			strsTiresStart, _ := strconv.Atoi(strsTires[0])
			strsTiresEnd, _ := strconv.Atoi(strsTires[1])
			for i := strsTiresStart; i <= strsTiresEnd; i++ {
				if len(pages) > 0 {
					if pages[len(pages)-1] < i {
						pages = append(pages, i)
					}
				} else {
					pages = append(pages, i)
				}
			}
		} else {
			pInt, _ := strconv.Atoi(strs[ind])
			if len(pages) > 0 {
				if pages[len(pages)-1] < pInt {
					pages = append(pages, pInt)
				}
			} else {
				pages = append(pages, pInt)
			}
		}
	}

	// теперь делаем массив в обратную стрпону
	outs := make([]int, k)
	for i := 0; i < len(pages); i++ {
		outs[pages[i]-1] = 1
	}

	// Ну и форматируем в норм вид
	var outString []string
	var isStart int
	var isDiap bool
	//fmt.Printf("%#v\n", outs)
	for i := 0; i < len(outs)-1; i++ {

		if outs[i] == 1 && isDiap == false {
			outString = append(outString, strconv.Itoa(isStart+1)+"-"+strconv.Itoa(i+1))
			isDiap = true
			isStart = i
		}
		if outs[i] == 1 && isDiap == true {
			outString[len(outString)] = strconv.Itoa(isStart+1) + "-" + strconv.Itoa(i+1)
			isStart = i
		}

		if outs[i] == 1 && outs[i+1] == 1 {
			isDiap = true
		}
		if outs[i] == 0 {
			isDiap = false
		}

		//if outs[i] == 1 && outs[i+1] == 0 {
		//	outString += strconv.Itoa(i) + ","
		//}

	}

	return strings.Join(outString, ",")
}
