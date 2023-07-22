package main

// Account TODO: add a fields and contructure for this.
type Account struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Routes    []Route `json:"routes,omitempty"`
}

type Route struct {
	ID          int    `json:"id"`
	Number      int64  `json:"number"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
	}
}
