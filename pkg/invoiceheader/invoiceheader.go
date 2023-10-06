package invoiceheader

import "time"

//model of invoiceheader
type Model struct {
	Id       uint
	Client   string
	CreateAt time.Time
	UpdateAt time.Time
}
