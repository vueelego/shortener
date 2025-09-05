package dbrepo

type Repository struct {
	UserRepo  UserRepo
	EntryRepo EntryRepo
	ClickRepo ClickRepo
}
