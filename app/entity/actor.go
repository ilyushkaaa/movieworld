package entity

type Actor struct {
	ID      uint64
	Name    string
	Surname string
	Films   []*Film
}
