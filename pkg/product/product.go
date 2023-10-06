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
type Storage interface {
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetById(uint) (*Model, error)
	Delete(uint) error
}
