package main

import (
	"fmt"
	"strings"
	"time"
	"strconv"
)

func PrintMainScreen() {
	fmt.Println("Welcome to our pay terminal!")
	fmt.Println("Choose your operator:")
	fmt.Println("1 - Megafon")
	fmt.Println("2 - Babilon")
	fmt.Println("3 - Tcell")
	fmt.Println("4 - Beeline")
}

// Need to modify
func CheckNumber(operatorID int, number string) bool {
	

	var checkPrefix bool
	operatorPreffix := map[int]string{
		1: "90,88",
		2: "918,98",
		3: "93,50,92",
		4: "919,917",
	}

	if _, err := strconv.Atoi(number); err == nil {
		if len(number) != 9 {
			return false
		}
	
		prefix := operatorPreffix[operatorID]
		arrPrefix := strings.Split(prefix, ",")
	
		for _, _prefixValue := range arrPrefix {
			checkPrefix = strings.HasPrefix(number, _prefixValue)
	
			if checkPrefix {
				return true
			}
		}
	
		return checkPrefix
	}else{
		return false
	}
	
}

func printCheck(name string,num string,sum string){
	fmt.Println(strings.Repeat("-",50))
	fmt.Printf("Operator: %s \n",name)
	fmt.Printf("Number: %s\n",num)	
	fmt.Printf("Sum: %s\n",sum)
// Тут дата
	dt := time.Now()
    fmt.Println("Current date and time is: ", dt.Format("2006-01-02 15:04:05"))
	fmt.Println("OperationStatus: Successful")
	fmt.Println(strings.Repeat("-",50))
}
