package billing

import "container/list"

// billing status
const (
	Modified  string = "Modified"
	Sent      string = "Sent"
	Payed_for string = "Payed_for"
)

// Payment method
const (
	Check         string = "Check"
	Bank_transfer string = "Bank_transfer"
	Payapal       string = "Payapal"
	Other         string = "Other"
)

type Model struct {
	Billing_number   int
	Entitled         string
	Billing_status   string `validate:"Modified, Sent, Payed_for"`
	Creation_date    string
	Payment_deadline string
	Payment_method   string `validate:"Check, Bank_transfer, Payapal, Other"`
	Time_of_payment  string
	Details          string
	Billing_line_id  list.List
}
