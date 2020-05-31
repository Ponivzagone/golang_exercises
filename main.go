package main

import (
    "fmt"
    "os"
    "bufio"
	"strconv"
	"flag"
	"strings"

	"./exercises/htree"
)

import "C"

//export HTreeClient
func HTreeClient() {
	fmt.Println(htree.Description())
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Длинна массива: ")
	value, _ := reader.ReadString('\n')
    n, err := strconv.Atoi(strings.Trim(value,"\t \n"))
	if err != nil { panic(err) }
	
	HT := htree.MakeHeightTree(n)
	fmt.Println("Введите массив переходов узлов дерева: ")
	value, _ = reader.ReadString('\n')
	HT.ReadFrom(strings.Trim(value,"\t \n"))

	fmt.Println("Переходы: ", HT.Tree())
	HT.ComputJump()
	fmt.Println("Число прыжков до корня из каждого узла: ", HT.Jump())

	fmt.Println("Высота дерева: ", HT.GetHeight())
}

//TODO Положить все будующие обработчики (клиентов для задачек) в массив функций и без параметров вызывать все функции

func main()  {
	HtreeActive := flag.Bool("htree", false, "Высота дерева")
	flag.Parse()

	if *HtreeActive == true {
		HTreeClient()
	}
}