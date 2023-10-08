package invoice

import "time"

//model of invoice
type Model struct {
	Id         uint
	Invoive_Id uint
	Product_Id uint
	CreateAt   time.Time
	UpdateAt   time.Time
}

//slice of model
type Models []*Model

//Storage interface
type Storage interface {
	Migrate() error
}

//Service of product
type Service struct {
	storage Storage
}

//NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

//Migrate use for migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
