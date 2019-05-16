package main

import "fmt"

func main() {
	x:=45 //какое то двухзначное число
	fmt.Println(x/10)// колличество десяток
	fmt.Println(x%10)// колличество единиц
	fmt.Println((x%10)+(x/10)) //Сумма цифр
	fmt.Println((x%10)*(x/10)) //произведение его цифр
}
