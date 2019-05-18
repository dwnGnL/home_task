package main

import "fmt"
import "os"
import "strconv"

func main() {
	var exitTerminal bool
	var operator string
	var number string
	var sum string

	for !exitTerminal {
		var operatorCheck bool
		var numberCheck bool
		var sumCheck bool
		var commandCheck bool
		var command string
		var operator2 int
		var sum2 float64
		var err error

		PrintMainScreen()

		for !operatorCheck {
			fmt.Scan(&operator)
			if operator=="esc" {
				os.Exit(1)
			}else{
				if operator2, err = strconv.Atoi(operator); err == nil {
					if operator2 >= 1 && operator2 <= 4 {
						operatorCheck = true
					} else {
						fmt.Println("Please, choose right operator")
					}
				} 
				// if operator >= 1 && operator <= 4 {
				// 	operatorCheck = true
				// } else {
				// 	fmt.Println("Please, choose right operator")
				// }
			}
		}

		for !numberCheck {
			fmt.Println("Enter your number (Operator): ")
			fmt.Scan(&number)
			if number=="esc"{
				os.Exit(1)
			}else{
				if CheckNumber(operator2, number) {
					numberCheck = true
				} else {
					fmt.Println("Please, enter right number")
				}
			}
		}

		for !sumCheck {
			fmt.Println("Enter sum:")
			fmt.Scan(&sum)
			if sum=="esc" {
				os.Exit(1)
			}else{
				if sum2, err = strconv.ParseFloat(sum, 32); err == nil {
					if sum2 > 0 {
						sumCheck = true
					} else {
						fmt.Println("Please, enter right sum")
					}
				} 
			}
			
		}

		switch operator {
		case "1":printCheck("Megafon",number,sum)
		case "2":printCheck("Babilon",number,sum)
		case "3":printCheck("Tcell",number,sum)
		case "4":printCheck("Beeline",number,sum)	
		}

		fmt.Println("esc- to go exit, back - to go to terminal")

		for !commandCheck {
			fmt.Scan(&command)

			if command == "esc" {
				commandCheck = true
				exitTerminal = true
			} else if command == "back" {
				commandCheck = true
				continue
			} else {
				fmt.Println("command not recongize")
			}
		}
	}

	fmt.Println("application terminated")
}
