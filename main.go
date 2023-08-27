package main

import (
	"bufio"
	"encoding/json"
	"errors"
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
func Work(input *bufio.Reader) (outputStr string) {
	var t int
	fmt.Fscan(input, &t)
	for ii := 0; ii < t; ii++ {
		var n int
		fmt.Fscan(input, &n)
		var jsonstr string
		for iii := 0; iii <= n; iii++ {
			var str string
			// fmt.Fscanln(input, &str)
			str, _ = input.ReadString('\n')

			jsonstr += str
		}

		// fmt.Println(jsonstr)

		// Получить ответ
		var cat Cat
		json.Unmarshal([]byte(jsonstr), &cat)

		// var c Category3Base

		// c := &Category3Base{}
		var c *Category3Base = new(Category3Base)
		c.Cat3 = make(map[int]*Category3Base)
		c.ID = cat[0].ID
		c.Name = cat[0].Name
		for i := 1; i < len(cat); i++ {

			s := find(c, cat[i].Parent, cat[i].ID, cat[i].Name)
			fmt.Println("parent:", s.ID, s.Name)

			s.Cat3[cat[i].ID] = &Category3Base{ID: cat[i].ID, Name: cat[i].Name}
			s.Cat3[cat[i].ID].Cat3 = make(map[int]*Category3Base)
			s.Cat3[cat[i].ID].ID = cat[i].ID
			s.Cat3[cat[i].ID].Name = cat[i].Name

			fmt.Printf("%+#v\n", c)
			var prefix string = "-"
			for _, CatBase := range c.Cat3 {
				printCategory3Base(CatBase, prefix)
			}
		}
		// c.PrintCat3()
		var prefix string = "-"
		for _, CatBase := range c.Cat3 {
			printCategory3Base(CatBase, prefix)
		}

	}

	return outputStr
}
func find(c *Category3Base, idparentfind, newid int, newname string) *Category3Base {
	if idparentfind == 0 {
		return c
	}
	var s *Category3Base = new(Category3Base)
	for Id, CatBase := range c.Cat3 {
		if Id == idparentfind { // Если это этот ID
			s = CatBase
			c.Cat3[Id].ID = newid
			c.Cat3[Id].Name = newname
		}
		FindCat, ErrorFind := findCategory3Base(CatBase, idparentfind)
		if ErrorFind == nil {
			s = FindCat
			FindCat.ID = newid
			FindCat.Name = newname
		}
	}
	return s
}

type Cat []struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Parent int    `json:"parent,omitempty"`
}
type Category3Base struct {
	ID   int    `json:"id"`
	Name string `json:"name"` // Навазние категории

	// Системная информация
	Cat3 map[int]*Category3Base // Мапа потомков
}

type CATEG struct {
	Cat3 Category3Base
}

var ErrorNotFoundCatId error = errors.New("FindCat3: Not found Category3Base in Cat3")
var ErrorNotFoundCatfindCategory3Base error = errors.New("findCategory3Base: not found Category3Base in Cat3")

// Поиск по ID в дереве категорий
func (woo *Category3Base) FindCat3(FindId int) *Category3Base {
	//fmt.Println(" woo.Cat3", woo.Cat3)
	for Id, CatBase := range woo.Cat3 {
		if Id == FindId { // Если это этот ID
			return CatBase
		}
		//fmt.Println("CatBase", CatBase.Cat3)
		FindCat, ErrorFind := findCategory3Base(CatBase, FindId)
		//fmt.Println("CatBase!!!", FindCat, ErrorFind)
		if ErrorFind == nil {
			return FindCat
		}
	}
	return nil
}

// Поиск именно в ячейке
func findCategory3Base(cat *Category3Base, FindId int) (*Category3Base, error) {
	//fmt.Println("cat.Cat3", cat.Cat3)
	if FindId == 0 {
		return cat, nil
	}
	for Id, CatBase := range cat.Cat3 {
		//fmt.Println("ID", Id == FindId, Id, FindId)
		if Id == FindId { // Если это этот ID
			return CatBase, nil
		}
		//if len(CatBase.Cat3) != 0 {
		if CatBase.Cat3 != nil && len(CatBase.Cat3) != 0 {
			FindCat, ErrorFindCat := findCategory3Base(CatBase, FindId)
			//fmt.Println("---FindCat", FindCat, ErrorFindCat, " -- ", FindCat.Name, FindCat.Slug, FindCat.Parent)
			if ErrorFindCat == nil { // Если не нашёл
				return FindCat, nil
			}
		}
		//fmt.Println("Exit")
	}
	return &Category3Base{}, ErrorNotFoundCatfindCategory3Base
}
func (woo *Category3Base) PrintCat3() {
	var prefix string = "-"
	for _, CatBase := range woo.Cat3 {
		printCategory3Base(CatBase, prefix)
	}
}

// Обход всех потомков и вывод на экран
func printCategory3Base(cat *Category3Base, prefix string) {
	for Id, CatBase := range cat.Cat3 {
		fmt.Println(prefix, Id, CatBase.Name)
		printCategory3Base(CatBase, prefix+"-")
	}
}
