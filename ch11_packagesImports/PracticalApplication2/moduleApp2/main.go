/*
One of your colleagues has created a program composed of a unique main.go file :
You are asked to refactor the code to improve its maintainability. Propose a new code organization :

    Which packages should you create?
	Packages:
	- email
	- invoice
    Should you create new directories?
*/

// package-imports/application-refactor/problem/main.go
package main

import (
	"fmt"
	"moduleApp2/email"
	"moduleApp2/invoice"
)

func main() {

	// first reservation
	customerName := "Doe"
	customerEmail := "john.doe@example.com"
	var nights uint = 12
	emailContents := email.GetEmailContents("M", customerName, nights)
	fmt.Println(emailContents)
	email.SendEmail(emailContents, customerEmail)
	invoice.CreateAndSaveInvoice(customerName, nights, 145.32)
}

// // send an email
// func sendEmail(contents string, to string) {
// 	// ...
// 	// ...
// }

// // prepare email template
// func getEmailContents(title string, name string, nights uint) string {
// 	text := "Dear %s %s,\n your room reservation for %d night(s) is confirmed. Have a nice day !"
// 	return fmt.Sprintf(text,
// 		title,
// 		name,
// 		nights)
// }

// // create the invoice for the reservation
// func createAndSaveInvoice(name string, nights uint, price float32) {
// 	// ...
// }
