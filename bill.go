package main

import (
	"fmt"
	"os"
)

// A Struct reprsenting a bill for a customer
type Bill struct {
	CustomerName string
	Items        map[string]float64
	Tip          float64
}

// Create a new bill for a give customer
func GenerateBill(customerName string) Bill {
	newBill := Bill{
		CustomerName: customerName,
		Items:        map[string]float64{},
		Tip:          0,
	}
	return newBill
}

// Generate a formatted string represeting the bill
func (billInstance *Bill) FormatBill() string {
	billString := "Bill Breakdown: \n"
	var totalAmount float64

	// list items
	for itemName, itemPrice := range billInstance.Items {
		billString += fmt.Sprintf("%-25v ...$%v \n", itemName+":", itemPrice)
		totalAmount += itemPrice
	}

	// Tip
	totalAmount += billInstance.Tip
	billString += fmt.Sprintf("%-25v ...$%0.2f \n", "Tip:", billInstance.Tip)

	// Total
	billString += fmt.Sprintf("%-25v ...$%0.2f", "Total:", totalAmount)

	return billString
}

// Updates the tip amount on the bill
func (billInstance *Bill) UpdateTip(tipAmount float64) {
	billInstance.Tip = tipAmount
}

// Adds an item and its price to the bill
func (billInstance *Bill) AddItem(itemName string, itemPrice float64) {
	billInstance.Items[itemName] = itemPrice
}

// Saves the bill to a text file
func (billInstance *Bill) SaveBill() {
	formattedBill := []byte(billInstance.FormatBill())
	fileName := "bills/" + billInstance.CustomerName + ".txt"

	err := os.WriteFile(fileName, formattedBill, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Bill for %v saved to file.\n", billInstance.CustomerName)
}
