package repository

type Authentication interface {

}

type Adverts interface {

}

type Users interface {

}

type Repository struct {
	Authentication
	Adverts
	Users
}

func NewRepository() *Repository {
	return &Repository{}
}
