package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A Helper function that prompts the user with a given string and returns the user's input.
func GetUserInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	userInput, err := reader.ReadString('\n')
	return strings.TrimSpace(userInput), err
}

// Prompts the user for a customer name and creates a new bill for that customer.
func CreateNewBill() Bill {
	inputReader := bufio.NewReader(os.Stdin)

	username, _ := GetUserInput("Enter Customer Name: ", inputReader)

	newBill := GenerateBill(username)
	fmt.Println("Created Bill For -", newBill.CustomerName)

	return newBill
}

// Prompts the user with a menu of options
func PromptUserOptions(bill Bill) {
	inputReader := bufio.NewReader(os.Stdin)

	userChoice, _ := GetUserInput("\nChoose an Option (a - Add Item, s - Save Bill, t - Add Tip): ", inputReader)

	switch userChoice {
	case "a":
		itemName, _ := GetUserInput("Item Name: ", inputReader)
		itemPriceStr, _ := GetUserInput("Item Price ($): ", inputReader)

		itemPrice, err := strconv.ParseFloat(itemPriceStr, 64)
		if err != nil {
			fmt.Println("The price must be a number.")
			PromptUserOptions(bill)
		}

		bill.AddItem(itemName, itemPrice)
		fmt.Println("Item Added -", itemName, "$"+itemPriceStr)
		PromptUserOptions(bill)

	case "t":
		tipStr, _ := GetUserInput("Enter Tip Amount ($): ", inputReader)

		tip, err := strconv.ParseFloat(tipStr, 64)
		if err != nil {
			fmt.Println("The tip must be a number.")
			PromptUserOptions(bill)
		}

		bill.UpdateTip(tip)
		fmt.Println("Tip Added -", "$"+tipStr)
		PromptUserOptions(bill)

	case "s":
		bill.SaveBill()

	default:
		fmt.Println("Please choose a valid option...")
		PromptUserOptions(bill)
	}
}

// Entry point of the program.
func main() {
	customerBill := CreateNewBill()
	PromptUserOptions(customerBill)
}
