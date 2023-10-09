package product

import "time"

//model of product
type Model struct {
	Id           uint
	Name         string
	Observations string
	Price        int
	CreateAt     time.Time
	UpdateAt     time.Time
}

//slice of model
type Models []*Model

//Storage interface
type Storage interface {
	Migrate() error
	Create(*Model) error
	//Update(*Model) error
	//GetAll() (Models, error)
	//GetById(uint) (*Model, error)
	//Delete(uint) error
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

func (s *Service) Create(m *Model) error {
	m.CreateAt = time.Now()
	return s.storage.Create(m)
}
