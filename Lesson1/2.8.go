package main

import "fmt"

func main() {
	var k int
	var d int
	fmt.Scan(&k)
	fmt.Println(k%7) //Если первый день понедельник
	fmt.Println(k%7+1) //Если первый день вторник
	fmt.Scan(&d)
	fmt.Print(k%7+(d-1)) //первый день выбрали сами от 1 до 7
}
// на все не тестил но вроде работает