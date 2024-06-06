package main

import "fmt"

func main() {
	customerBill := GenerateBill("Mario's Bill")
	
	fmt.Println(customerBill.FormatBill())
}