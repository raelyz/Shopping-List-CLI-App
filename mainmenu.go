package main

import (
	"fmt"
)

var currListIndex int
var allLists = []shoppingList{}
var currentList = shoppingList{}

func main() {
	var on bool = true
	var input string

	for on {
		fmt.Println(" _________________________________")
		fmt.Println("|                                 |")
		fmt.Println("|    Shopping List Application    |")
		fmt.Println("|_________________________________|")
		fmt.Println("")

		if len(allLists) == 0 {
			fmt.Printf("         No List Available        \n")
			fmt.Printf("     Press Enter to Create one    \n")
			fmt.Println("         Key in 0 to exit")
			fmt.Scanln(&input)
			if input == "0" {
				break
			}
			createShoppingList()
			listMenu()
		} else {
			fmt.Printf("     1. Create New List\n")
			fmt.Printf("     2. Load Existing List\n")
			fmt.Println("     0. Quit")
			fmt.Scanln(&input)
		}

		if input == "1" {
			createShoppingList()
			listMenu()
		} else if input == "2" {
			loadShoppingList()
		} else if input == "0" {
			break
		} else {
			fmt.Println("Please key in a valid option")
		}

	}
	fmt.Println("=========================")
	fmt.Println("Application Closed")

}
