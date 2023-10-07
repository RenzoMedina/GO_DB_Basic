package main

import (
	"go-db/pkg/product"
	"go-db/storage"
	"log"
)

func main() {

	storage.NewPostgreDB()

	storageProd := storage.NewPsqlProduct(storage.Pool())
	serviceProd := product.NewService(storageProd)

	if err := serviceProd.Migrate(); err != nil {
		log.Fatalf("product.Migrate(): %v", err)
	}

}
