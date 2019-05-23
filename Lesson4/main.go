package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

const (
	Select = "select"
	Create = "create"
	Insert = "insert"
	Delete = "delete"
)

func main() {
	var login string
	var password string
	var ID int =1
	var logined bool
	db := map[string]*[]Student{}
	var exit bool

	readedPassHash := readPasswordHash()

	for !logined {
		fmt.Println("Enter username: ")
		fmt.Scan(&login)
		fmt.Println("Enter pass:")
		fmt.Scan(&password)

		if login == "root" && checkHash(password, readedPassHash) {
			logined = true
			fmt.Println("Hello, root")
		} else {
			fmt.Println("Invalid credentials")
		}
	}

	for !exit {
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		command = strings.Trim(command, " ")
		command = strings.ToLower(command)
		commandStruct := strings.Split(command, " ")

		switch commandStruct[0] {
		 case Select:
			tableName := commandStruct[3]
			tableName=strings.Trim(tableName," ")

			if db[tableName] == nil {
				fmt.Print(db[tableName])
				fmt.Println("table not exits")
			} else {
				columns := commandStruct[1]
				switch columns {
				case "*":
					if len(commandStruct)>4{
						if commandStruct[4]=="Where"{
							switch commandStruct[6] {
							case "==":
								switch commandStruct[5]{
								case "id":
									// tableName := commandStruct[3]
									arrayOfStudents := *db[tableName]
									for _, row := range arrayOfStudents {
										if string(row.ID)==commandStruct[7]{
											fmt.Printf("  ID|               Fname|Age|Average| \n")
											fmt.Println("------------------------------------")
											
											fmt.Printf("%4d|%20s|%3d|%2.2f\n", row.ID, row.Fname, row.Age, row.Average)
											
										}
									}
								}
								
							}
						}else{
							fmt.Println(strings.Repeat("-",50))
							fmt.Println("| argument of command not recognize |")
							fmt.Println(strings.Repeat("-",50))
						}
					}else{
						fmt.Printf("  ID|               Fname|Age|Average|Experience \n")
						fmt.Println("------------------------------------")
						arrayOfStudents := *db[tableName]

						for _, row := range arrayOfStudents {
							fmt.Printf("%4d|%20s|%3d|%4.2f|%2d\n", row.ID, row.Fname, row.Age, row.Average,row.Experience)
						}
						fmt.Println(strings.Repeat("-",50))
						fmt.Println("| rows returned: ", len(arrayOfStudents))
						fmt.Println(strings.Repeat("-",50))
					}
					break
				
				default:
					arrayOfStudents := *db[tableName]
					colls:=strings.Split(commandStruct[1],",")
					for _, row := range arrayOfStudents {
						for _,coll:= range colls{
							coll=strings.ToLower(coll)
							switch coll {
							case "id":
								// fmt.Printf("  ID|")
								fmt.Printf("%4d|", row.ID)
								break
							case "fname":
								// fmt.Printf("               Fname|")
								fmt.Printf("%20s|",row.Fname)
								break
							case "age":
								// fmt.Printf("Age|")
								fmt.Printf("%3d|",row.Age)
								break
							case "average":
								// fmt.Printf("Average ")
								fmt.Printf("%2.2f",row.Average)
								break
							default: fmt.Print("Ooops somefing wrong!")
							}
							
							// fmt.Printf("%4d|%20s|%3d|%2.2f\n", row.ID, row.Fname, row.Age, row.Average)
						}
						fmt.Println("")
					}
					
					break
				}

			}

			break
		case Create:
			tableName := commandStruct[2]
			emptySlice := []Student{
				Student{
					ID:         1,
					Age:        54,
					Fname:      "Bill Gates",
					Experience: 30,
				},
			}
			db[tableName] = &emptySlice
			fmt.Println("table created: " + tableName)
			break
		
		case Insert:
			args:=commandStruct[1]
			arg:=strings.Split(args,",")
			tableName:=commandStruct[3]
			emptySlice := []Student{
				Student{
					ID:         1,
					Age:        54,
					Fname:      "Bill Gates",
					Experience: 30,
				},
			}
			age,_:=strconv.Atoi(arg[1])
			isStudent,_:=strconv.ParseBool(arg[2])
			exp,_:=strconv.Atoi(arg[3])
			if len(arg)==4{
				if db[tableName] == nil {
					fmt.Println("table not exits")
				}else{
					ID++
					emptySlice=append(emptySlice,Student{
						ID:         ID,
						Age:        age,
						Fname:      arg[0],
						IsStudent:   isStudent,
						Experience: exp,
					},)
					db[tableName] = &emptySlice
				}
			}else{
				fmt.Println(strings.Repeat("-",50))
				fmt.Println("| check value of coll or count of coll |")
				fmt.Println(strings.Repeat("-",50))
			}

		case Delete:
			tableName := commandStruct[1]
				if commandStruct[2]=="where"{
					arrayOfStudents := *db[tableName]
					b,_:=strconv.Atoi(commandStruct[5])
					for _, row := range arrayOfStudents {
						if row.ID==b{
							arrayOfStudents[b]=Student{}
						}
					}
				}else{
					fmt.Println(strings.Repeat("-",50))
					fmt.Println("| check where argument |")
					fmt.Println(strings.Repeat("-",50))
				}
			
			
		default:
			fmt.Println(strings.Repeat("-",25))
			fmt.Println("| command not recognize |")
			fmt.Println(strings.Repeat("-",25))
		}
	}

}
