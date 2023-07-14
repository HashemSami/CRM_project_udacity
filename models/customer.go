package models

type Customer struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     uint64 `json:"phone,omitempty"`
	Contacted *bool  `json:"contacted,omitempty"`
}

var f = false
var t = true

var MockData = map[string]Customer{
	"1": {
		ID:        "1",
		Name:      "Hashem",
		Role:      "employee",
		Email:     "test@test.com",
		Phone:     056464567,
		Contacted: &t,
	},
	"2": {
		ID:        "2",
		Name:      "Sarah",
		Role:      "employee",
		Email:     "test1@test1.com",
		Phone:     55465478,
		Contacted: &f,
	},
	"3": {
		ID:        "3",
		Name:      "Billy",
		Role:      "employee",
		Email:     "test3@test3.com",
		Phone:     5554687964,
		Contacted: &t,
	},
}
