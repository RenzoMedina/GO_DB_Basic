package invoiceheader

import "time"

//model of invoiceheader
type Model struct {
	Id       uint
	Client   string
	CreateAt time.Time
	UpdateAt time.Time
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
