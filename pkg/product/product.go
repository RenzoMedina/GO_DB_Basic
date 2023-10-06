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
