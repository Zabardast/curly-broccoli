package User

const (
	Business string = "Business"
	Person   string = "Person"
)

type Model struct {
	Id            int
	Name          *string
	Laste_name    *string
	Date_of_birth string
	Email         string
	Phone_number  string
	Turne_over    int
	Taxed_income  int
	Password      string
}
