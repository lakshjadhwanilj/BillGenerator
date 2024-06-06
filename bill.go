package main

import "fmt"

type Bill struct {
	CustomerName string
	Items        map[string]float64
	Tip          float64
}

// Generate new bills
func GenerateBill(customerName string) Bill {
	newBill := Bill{
		CustomerName: customerName,
		Items:        map[string]float64{},
		Tip:          0,
	}
	return newBill
}

// Format the bill
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

// Update tip
func (billInstance *Bill) UpdateTip(tipAmount float64) {
	billInstance.Tip = tipAmount
}

// Add items to the bill
func (billInstance *Bill) AddItem(itemName string, itemPrice float64) {
	billInstance.Items[itemName] = itemPrice
}
