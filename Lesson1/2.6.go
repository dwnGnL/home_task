package main

import "fmt"

func main() {
	var sec int
	var hour int
	var min int
	fmt.Scan(&sec)
	hour=sec/3600
	min=sec/60
	fmt.Printf("Прошло часов %d \n",hour)
	fmt.Printf("Прошло минут %d \n",min)
	fmt.Printf("Прошло секунд %d \n",sec)
}