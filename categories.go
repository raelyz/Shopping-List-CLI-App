package main

import (
	"fmt"
	"strconv"
)

func addNewCategory() {
	var input string
	fmt.Println("Add New Category Name")
	fmt.Println("What is the New Category Name to add?")
	fmt.Scanln(&input)
	input = stringSanitizer(input)
	if input == "" {
		fmt.Println("No Input Found!")
		return
	}
	categoriesAddress := &currentList.categories
	categories := *categoriesAddress
	currentList.categories = append(categories, input)
	fmt.Println("Category:", input, "added at index", len(categories)-1)
}

func categoryIndex(newCategory string) (int, bool) {
	categoriesAddress := &currentList.categories
	categories := *categoriesAddress
	for i, categoryName := range categories {
		if newCategory == categoryName {
			fmt.Printf("Category: %v already exists at index %d \n", newCategory, i)
			return i, true
		}
	}
	return -1, false
}

func emptyCategory(category int) bool {
	itemMap := &currentList.itemData
	for _, itemInfo := range *itemMap {
		if category == itemInfo.Category {
			return false
		}
	}
	return true
}

func categoryName() string {
	var input string
	for input == "" {

		fmt.Println("What category does it belong to?")
		fmt.Scanln(&input)
		if input == "" {
			fmt.Println("Please key in a category")
			continue
		}

	}
	return stringSanitizer(input)
}

func createNewCategory(category string) int {
	fmt.Println("Category:", category, "does not exist")
	categoriesAddress := &currentList.categories
	categories := *categoriesAddress
	currentList.categories = append(categories, category)
	index := len(currentList.categories) - 1
	fmt.Println("Creating", category, "at index", index)
	return index
}

func lastItemInCategory(index int) bool {
	var count int
	itemMap := &currentList.itemData
	for _, itemInfo := range *itemMap {
		if index == itemInfo.Category {
			count++
			if count > 1 {
				return false
			}
		}
	}
	fmt.Println("No items will be left in the previous Category")
	return true
}

func deleteCategory() {
	var input string
	fmt.Println("Key in the number for the category to be deleted")
	itemMap := currentList.itemData
	categoriesAddress := &currentList.categories
	categories := *categoriesAddress
	for i, category := range categories {
		fmt.Println(i, ".", category)
	}
	fmt.Scanln(&input)
	index, err := strconv.ParseInt(input, 10, 64)
	if err != nil || int(index) >= len(categories) {
		fmt.Println("Invalid input")
		fmt.Println("Returning to Main Menu")
		return
	}
	fmt.Println("Deleting Category:", categories[index])
	if int(index) == len(categories)-1 {
		for itemName, itemInfo := range itemMap {
			if int(index) == itemInfo.Category {
				delete(itemMap, itemName)
			}
		}
		categories = categories[:index]
	} else {
		categories = append(categories[:index], categories[(index+1):]...)

		for itemName, itemInfo := range itemMap {
			if int(index) == itemInfo.Category {
				delete(itemMap, itemName)
			} else if itemInfo.Category > int(index) {
				currentList.itemData[itemName].Category--
			}
		}
	}
	currentList.categories = categories
}
