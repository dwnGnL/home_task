package main

import "fmt"
import "os"
import "strconv"

func main() {
	var exitTerminal bool
	var count string

	for !exitTerminal {
		var countCheck bool
		var commandCheck bool
		var sizeCheck bool
		var checkType bool
		var command string
		var count2 int
		var types string
		var typ int
		var t string
		var sizenumber int
		var size []int
		var types2 []string
		var err error

		PrintMainScreen()

		for !countCheck {
			fmt.Scan(&count)
			if count=="esc" {
				os.Exit(1)
			}else{
				if count2, err = strconv.Atoi(count); err == nil {
					if count2 >0{
						countCheck = true
					} else {
						fmt.Println("Please, choose count > 0")
					}
				} 
			}
		}

		for !sizeCheck {
			i:=1
			for i<=count2 {
				fmt.Printf("Choose size of %d-st slice: \n",i)
				fmt.Println("1-40cm")
				fmt.Println("2-65cm")
				i++
				fmt.Scan(&t)
				if t=="esc"{
					os.Exit(1)
				}else{
					sizenumber, err =strconv.Atoi(t)
					if (err==nil){
						if sizenumber>=1 && sizenumber<=2{
							size = append(size,sizenumber)
							sizeCheck=true
					}else{
						fmt.Println("Choose right size")
						i--
					}
					}else{
						fmt.Println("Choose right size")
						i--
					}
					
					
				}
				
			}
		}	

		for !checkType{
			i:=1
			for i<=count2 {
				fmt.Printf("Choose Type of %d-st slice: \n",i)
				fmt.Println("1-Peperoni")
				fmt.Println("2-meat")
				fmt.Println("3-country")
				i++
				fmt.Scan(&types)
				if types=="esc"{
					os.Exit(1)
				}else{
					typ, err =strconv.Atoi(types)
					if (err==nil){
						if typ>=1 && typ<=3{
							types2 = append(types2,types)
							checkType=true
					}else{
						fmt.Println("Choose right type")
						i--
					}
					}else{
						fmt.Println("Choose right type")
						i--
					}
				}
			}
		}	
		
		printCheck(count2,size,types2)
		

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
