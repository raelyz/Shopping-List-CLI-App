package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addNewItem() bool {
	name := itemName()
	if itemExists(name) {
		fmt.Println("Item exists returning to Main Menu")
		fmt.Println("Select option 4 to update the Item")
		return false
	}
	category := categoryName()
	index, exists := categoryIndex(category)
	if !exists {
		index = createNewCategory(category)
	}
	units := numberOfUnits()
	cost := costPerUnit()

	newItem := item{
		Category: index,
		Quantity: units,
		UnitCost: cost,
	}
	currentList.itemData[name] = &newItem
	fmt.Println("Successfully added", units, "units of", name, "under", category, "at a cost of", cost, "per unit")
	return true
}

func stringSanitizer(input string) string {
	if input == "" {
		return ""
	}
	sanitizedString := strings.ToUpper(input[0:1]) + input[1:]
	return sanitizedString
}

type item struct {
	Category int
	Quantity int
	UnitCost float64
}

func itemExists(name string) bool {
	itemMap := currentList.itemData
	if _, found := itemMap[name]; !found {
		return false
	}
	return true
}

func itemName() string {
	var input string
	for input == "" {
		fmt.Println("What is the name of your item?")
		fmt.Scanln(&input)
		if input == "" {
			fmt.Println("Please key in a valid name")
			continue
		}
	}
	return stringSanitizer(input)
}

func numberOfUnits() int {
	var input string
	var units int64
	var err error
	for {
		fmt.Println("How many units are there?")
		fmt.Scanln(&input)
		units, err = strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("Please key in a number")
			continue
		}
		break
	}

	return int(units)
}

func costPerUnit() float64 {
	var input string
	var cost float64
	var err error
	for {
		fmt.Println("How much does it cost per unit?")
		fmt.Scanln(&input)
		cost, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Please key in a number")
			continue
		}
		break
	}

	return cost
}

func modifyItem() {
	var input string
	fmt.Println("Which item would you wish to modify?")
	fmt.Scanln(&input)
	name := stringSanitizer(input)
	if !itemExists(name) || name == "" {
		fmt.Println("Invalid Entry")
		fmt.Println("Returning to the Main Menu")
		return
	}
	itemMap := &currentList.itemData
	categoryAddress := &currentList.categories
	categories := *categoryAddress
	oldCategory := categories[(*itemMap)[name].Category]
	oldCost := (*itemMap)[name].UnitCost
	oldQuantity := (*itemMap)[name].Quantity
	fmt.Println("Current item name is", name, "- Category is", oldCategory, "- Quantity is ", oldQuantity, "- Unit Cost", oldCost)

	newName := modifiedName()
	newCategory, index := modifiedCategory(name)
	newCost := modifiedCost()
	newQuantity := modifiedQuantity()

	itemAddress := (*itemMap)[name]

	if newCategory == "" {
		fmt.Println("No changes to category made.")
	} else {
		itemAddress.Category = index
	}
	if newCost == 0 {
		fmt.Println("No changes to unit cost made.")
	} else {
		itemAddress.UnitCost = newCost
	}
	if newQuantity == 0 {
		fmt.Println("No changes to quantity made.")
	} else {
		itemAddress.Quantity = newQuantity
	}
	if newName == "" {
		fmt.Println("No changes to item name made.")
	} else {
		(*itemMap)[newName] = itemAddress
		delete(*itemMap, name)
	}
	fmt.Println("Press enter to return to the Main Menu")
	fmt.Scanln()
}

func modifiedCost() float64 {
	var input string
	fmt.Println("Enter new Unit Cost. Enter for no change")
	fmt.Scanln(&input)
	if input != "" {
		for {
			newCost, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Please key in a valid number")
				fmt.Scanln(&input)
				continue
			}
			return newCost
		}
	}
	return 0
}

func modifiedQuantity() int {
	var input string
	fmt.Println("Enter new Quantity. Enter for no change")
	fmt.Scanln(&input)
	if input != "" {
		for {
			quantity, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Please key in a valid number")
				fmt.Scanln(&input)
				continue
			}
			return int(quantity)
		}
	}
	return 0
}
func deleteItem() {
	var input string
	itemMap := &currentList.itemData
	fmt.Printf("\nDelete Item.\n")
	fmt.Println("Enter item name to delete:")
	fmt.Scanln(&input)
	item := stringSanitizer(input)
	if itemExists(item) {
		delete(*itemMap, item)
		fmt.Printf("\nDeleted %v \n", item)
		fmt.Printf("Returning to Main Menu\n")
		return
	}
	fmt.Printf("Item not found. Nothing to delete!\n")
}
func modifiedCategory(name string) (string, int) {
	var newCategory string
	fmt.Println("Enter new Category. Enter for no change")
	fmt.Scanln(&newCategory)
	for {
		if newCategory == "" {
			break
		}
		newCategory = stringSanitizer(newCategory)
		index, exists := categoryIndex(newCategory)
		if !exists {
			index = createNewCategory(newCategory)
			return newCategory, index
		}
		fmt.Println("Assigning over to index", index)
		return newCategory, index

	}
	return "", 0
}

func modifiedName() string {
	var input string
	fmt.Println("Enter new name. Enter for no change")
	fmt.Scanln(&input)
	if input != "" {
		input = stringSanitizer(input)
		for itemExists(input) {
			fmt.Println("Item already exists please choose another name Enter for no change")
			fmt.Scanln(&input)
			if input == "" {
				break
			}
		}
	}

	return input
}

func printData() {
	itemMap := &currentList.itemData
	for key, value := range *itemMap {
		fmt.Println(key, "-", value)
	}
}
