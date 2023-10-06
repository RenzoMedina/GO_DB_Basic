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
