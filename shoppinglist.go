package main

import (
	"fmt"
	"strconv"
)

func listMenu() {
	var listSelected bool = true
	var input string
	for listSelected {
		fmt.Println("\nCurrent List:", currentList.Name)
		fmt.Println("1. View Entire Shopping List.")
		fmt.Println("2. Generate Shopping List Report.")
		fmt.Println("3. Add Items.")
		fmt.Println("4. Modify Items.")
		fmt.Println("5. Delete Items.")
		fmt.Println("6. Add Category.")
		fmt.Println("7. Delete Category.")
		fmt.Println("8. Print Data.")
		fmt.Println("9. Save List.")
		fmt.Println("10. Back to Main Menu .")
		fmt.Println("Select your choice:")

		fmt.Scanln(&input)

		switch input {
		case "10":
			listSelected = false
			break
		case "1":

			viewList()
			break
		case "2":
			generateReport()
			break
		case "3":
			addNewItem()
			break
		case "4":
			modifyItem()
			break
		case "5":
			deleteItem()
			break
		case "6":
			addNewCategory()
			break
		case "7":
			deleteCategory()
			break
		case "8":
			printData()
			break
		case "9":
			saveList()
			break
		default:

		}
	}
}

func viewList() {
	if len(currentList.categories) == 0 {
		fmt.Println("List is Empty")
		fmt.Println("Please populate the Shopping List before Viewing!")
	} else {

		itemMap := currentList.itemData
		category := currentList.categories
		fmt.Println("")
		for itemName, item := range itemMap {

			fmt.Println("Category: "+category[item.Category]+" - Item: "+itemName+" Quantity: ", item.Quantity, " Unit Cost: ", item.UnitCost)
		}
		fmt.Println("Press Enter to go Back to the Shopping List Menu")
		fmt.Scanln()
	}

}

type shoppingList struct {
	Name       string
	categories []string
	itemData   map[string]*item
}

func loadShoppingList() {
	var input string
	var loadingList bool = true
	for loadingList {
		fmt.Printf("\nKey in the index of the list \nyou would like to load\n")

		for i, list := range allLists {
			fmt.Println(i+1, ". "+list.Name)
		}

		fmt.Println("\nKey in 0 to Exit")
		fmt.Scanln(&input)

		if input == "0" {
			break
		}

		currListIndex, _ = strconv.Atoi(input)
		currListIndex--
		if currListIndex >= 0 && currListIndex < len(allLists) {
			currentList = allLists[currListIndex]
			listMenu()
			return
		} else {
			fmt.Println("Please key in a valid option")
		}
	}

}

func createShoppingList() {
	var name string
	currListIndex = len(allLists)
	for name == "" {
		fmt.Println("Name your Shopping List")
		fmt.Scanln(&name)
		if name == "" {
			fmt.Println("Key in a List Title")
			continue
		}
		name = stringSanitizer(name)
		slice := []string{}
		currentList = shoppingList{
			Name:       name,
			categories: slice,
			itemData:   make(map[string]*item),
		}
	}
}

func saveList() {
	if currListIndex >= len(allLists) {
		allLists = append(allLists, currentList)
	} else {
		allLists[currListIndex] = currentList
	}
	fmt.Println("List Saved!")
}
