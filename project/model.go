package Project

import (
	Customer "Brocolie/customer"
)

// ca sert a rien
const (
	PROSPECT          string = "prospect"
	SENT_ESTIMATE     string = "send_estimate"
	ACCEPTED_ESTIMATE string = "accepted_estimate"
	STARTED           string = "started"
	FINISHED          string = "finished"
	CANCELED          string = "canceled"
)

type Model struct {
	Id       int
	Name     string
	Customer Customer.Model
	Status   string
	Price    int
}
