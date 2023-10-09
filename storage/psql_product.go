package storage

import (
	"database/sql"
	"fmt"
	"go-db/pkg/product"
)

const (
	psqlCreateProduct = ` CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`
	psqlInsertProduct = ` INSERT INTO products (name,observations, price,created_at) VALUES ($1,$2,$3,$4)
				RETURNING id
		
	`
)

// PsqlProduct use for work with postgres - product
type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct return a new pointer of PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate implements the interface product.Storage
func (p *PsqlProduct) Migrate() error {
	stament, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stament.Close()

	_, err = stament.Exec()
	if err != nil {
		return err
	}
	fmt.Println("migrate of product successfull")
	return nil
}

func (p *PsqlProduct) Create(m *product.Model) error {

	stament, err := p.db.Prepare(psqlInsertProduct)
	if err != nil {
		return err
	}
	defer stament.Close()

	err = stament.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreateAt,
	).Scan(&m.Id)
	if err != nil {
		return err
	}

	fmt.Println("Element was successfull")
	return nil
}
