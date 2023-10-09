package main

import (
	"fmt"
	"go-db/pkg/product"
	"go-db/storage"
	"log"
)

func main() {

	storage.NewPostgreDB()

	storageProd := storage.NewPsqlProduct(storage.Pool())
	serviceProd := product.NewService(storageProd)

	model := &product.Model{
		Name:         "Curso de moodle basic",
		Observations: "Creación de plugins ",
		Price:        80,
	}
	if err := serviceProd.Create(model); err != nil {
		log.Fatalf("product.Create(): %v", err)
	}

	fmt.Printf("%v\n", model)
}
