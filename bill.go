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
		Items:        map[string]float64{"Pie": 5.99, "Cake": 3.99},
		Tip:          0,
	}
	return newBill
}

// Format the bill
func (billInstance Bill) FormatBill() string {
    billString := "Bill Breakdown: \n"
    var totalAmount float64

    // list items
    for itemName, itemPrice := range billInstance.Items {
        billString += fmt.Sprintf("%-25v ...$%v \n", itemName+":", itemPrice)
        totalAmount += itemPrice
    }

    // Total
    billString += fmt.Sprintf("%-25v ...$%0.2f", "total:", totalAmount)

    return billString
}