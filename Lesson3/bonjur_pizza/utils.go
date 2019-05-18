package main

import (
	"fmt"
	"strings"
	"time"
	"crypto/md5"
	"strconv"
	
)

func PrintMainScreen() {
	fmt.Println("Welcome to our pay pizza restorant")
	fmt.Println("How many slices you choose?")
}



func printCheck(count int,size []int,types []string){
	var cash int
	var total int=0
	
	fmt.Println(strings.Repeat("-",100))
	fmt.Printf("customer: %s \n","Vitaliy")
	fmt.Printf("cashier: %s\n","Barzu")	
	for i:=0;i<count;i++{
		if size[i]==1{
			cash=1
		}else{
			cash=2
		}
		switch types[i] {
		case "1":fmt.Printf("Peperoni size - %d ---------%d usd \n",size[i],(10*cash))
		total+=10*cash
		case "2":fmt.Printf("Meat size - %d ---------%d usd \n",size[i],(15*cash))	
		total+=15*cash
		case "3":fmt.Printf("Country size - %d ---------%d usd \n",size[i],(12*cash))	
		total+=12*cash
		}
	}
	fmt.Println(strings.Repeat("-",100))
	fmt.Printf("Total price: %d usd\n",total)
// Тут дата
	dt := time.Now()
	fmt.Println("Current date and time is: ", dt.Format("2006-01-02 15:04:05"))
	md6 := []byte(strings.Join([]string{strconv.Itoa(total)}, dt.Format("2006-01-02 15:04:05")))
	
	fmt.Printf("Hash: %x \n",md5.Sum(md6))

	fmt.Println("OperationStatus: Successful")
	fmt.Println(strings.Repeat("-",100))
}
