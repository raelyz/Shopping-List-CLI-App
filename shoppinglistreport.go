package main

import "fmt"

func generateReport() {
	var on bool = true
	for on {
		var input string
		fmt.Printf("\nGenerate Report\n")
		fmt.Println("=========================")
		fmt.Println("1. Total Cost of each category.")
		fmt.Println("2. List of item by category.")
		fmt.Println("3. Main Menu.")
		fmt.Println("")
		fmt.Println("Choose your report:")
		fmt.Scanln(&input)
		fmt.Println("=========================")
		switch input {
		case "3":
			on = false
			break
		case "1":
			costOfEachCategory()
			break
		case "2":
			listByCategory()
			break
		default:
			fmt.Println("Key in a valid input from 1 to 3")
		}

	}
}

func listByCategory() {
	fmt.Printf("\nItem List By Category\n")
	fmt.Println("=========================")
	itemMap := currentList.itemData
	categories := currentList.categories
	for index, category := range categories {
		for item, itemInfo := range itemMap {
			if index == itemInfo.Category {
				fmt.Println("Category: "+category+" - Item: "+item+" Quantity: ", itemInfo.Quantity, " Unit Cost: ", itemInfo.UnitCost)
			}
		}

	}
	fmt.Println("Press Enter to go back to previous Menu")
	fmt.Scanln()
}

func costOfEachCategory() {
	fmt.Printf("\nTotal Cost By Category\n")
	fmt.Println("=========================")
	itemMap := currentList.itemData
	categories := currentList.categories
	for index, category := range categories {
		var cost float64
		for _, itemInfo := range itemMap {
			if index == itemInfo.Category {
				cost += float64(itemInfo.Quantity) * itemInfo.UnitCost
			}
		}
		fmt.Println(category+" cost: ", cost)
	}
	fmt.Println("Press Enter to go back to previous Menu")
	fmt.Scanln()
}
