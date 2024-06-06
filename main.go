package main

import (
	"bufio"
	"fmt"
	"os"
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

func promptOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)

	choice, _ := GetUserInput("Choose an Option (a - Add Item, s - Save Bill, t - Add Tip): ", reader)

	switch choice {
	case "a":
		itemName, _ := GetUserInput("Item Name:", reader)
		itemPrice, _ := GetUserInput("Item Price:", reader)

		fmt.Println(itemName, itemPrice)
	case "t":
		tip, _ := GetUserInput("Enter Tip Amount ($):", reader)
		fmt.Println(tip)
	case "s":
		fmt.Println("You chose s")
	default:
		fmt.Println("Please choose a valid option...")
		promptOptions(b)
	}
}

// Entry point of the program.
func main() {
	customerBill := CreateNewBill()
	promptOptions(customerBill)
	fmt.Println(customerBill)
}
