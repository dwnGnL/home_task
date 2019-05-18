package main

import (
	"os"
    "regexp"
	"fmt"
	"time"
	"strconv"
	)

func main(){
	var what bool = true
	var kompany []string
		var v_kompany string
		var number string
		var kom string
		var again string
		var sum int
		var kompany_ready bool = true
		var num_ready bool = true
		var sum_ready bool = true
		var stop bool = true
		
		kompany = append(kompany, "Megafon","Beeline","Babilon","Tcell")
	
		for what{
			kompany_ready=true
			num_ready=true
			sum_ready = true
			stop = true
		for kompany_ready{
		fmt.Print("\n")
		for i, value :=range kompany{
			fmt.Print(i+1)
			fmt.Printf(" - %s \n",value)
		}
		fmt.Scan(&v_kompany)
		
		switch v_kompany {
		case "1": kompany_ready=false
		case "2": kompany_ready=false
		case "3": kompany_ready=false
		case "4": kompany_ready=false
		case "esc": os.Exit(1)
			
		}
		
		}
		
			for num_ready{
			fmt.Print("\nEnter your number (Operator): ")
			fmt.Scan(&number)

			if checkNumber(number){
				num_ready = false
			}else if number=="esc"{
				os.Exit(1)
			}else{
				fmt.Print("Please, enter right operator number")
			}
			}
		
			
				for sum_ready{
				fmt.Print("Enter sum:")
				fmt.Scan(&sum)

				// exit, _ := strconv.Atoi(sum)
				
				// 	if sum=="esc" {
				// 		os.Exit(1)
				// 	}
				// 	if exit<=0 {
				// 		fmt.Println("Please, enter right sum > 0") 
				// 	}
				// 	if exit>0{
				// 		sum_ready=false
				// 	}
			

				
				 if sum<=0 {
					fmt.Print("Please, enter right sum > 0")
				}else{
					sum_ready=false
				}
				}
			vibor, err := strconv.Atoi(v_kompany)
			if err == nil {
				for i, value :=range kompany{
					if (i==(vibor-1)){
						kom=value
					}
					
				}
			}
			
			fmt.Println(kom)
			printCheck(kom,number,sum)
			for stop{
			fmt.Print("0 - if you want pay again, 9 - if you want exit")
			fmt.Scan(&again)
			if again == "0"{
			what = true
			stop=false
			}else if again=="9"{
				os.Exit(1)
			}else{
				fmt.Println("Please choose 1 or 9")
			}
		}
		}
	}
		
	


func checkNumber(num string) bool{
	if (len(num)==9){
		
		matched,err:= regexp.Match("[9]|[5][0-9]*",[]byte(num))
		check(err)
		if matched{
			
			return true	
		}else{
			fmt.Println("Please, enter right operator number")
			return false
		}
		
	}
	return false
}

func check(err error) {
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}

func printCheck(name string,num string,sum int){
	fmt.Println("--------------------")
	fmt.Printf("Operator: %s \n",name)
	fmt.Printf("Number: %s\n",num)	
	fmt.Printf("Sum: %d\n",sum)
// Тут дата
	dt := time.Now()
    fmt.Println("Current date and time is: ", dt.Format("2006-01-02 15:04:05"))
	fmt.Println("OperationStatus: Successful")
	fmt.Println("--------------------")
}