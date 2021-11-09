package Line

import "container/list"

type Model struct {
	Id       list.Element
	Name     string
	Price    int
	Quantity int
	About    string
}
