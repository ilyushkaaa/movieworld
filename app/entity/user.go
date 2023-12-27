package entity

type User struct {
	ID             uint64
	Username       string
	password       string
	FavouriteFilms []*Film
	Reviews        []*Review
}
