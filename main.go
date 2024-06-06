package main

import "fmt"

func main() {
	customerBill := GenerateBill("Mario's Bill")

	customerBill.AddItem("Onion Soup", 4.5)
	customerBill.AddItem("Veg Pie", 8.95)
	customerBill.AddItem("Toffee Pudding", 4.95)
	customerBill.AddItem("Coffee", 3.55)

	customerBill.UpdateTip(10)

	fmt.Println(customerBill.FormatBill())
}
