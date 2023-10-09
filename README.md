# Conexión a Base de datos con GO en PostgreSQL & MySQL

## PostgreSQL
### Paquete de instalación
Para mayor información del paquete que se utilizará para la conexión con este gestor de base datos visitar
[pq package](https://pkg.go.dev/github.com/lib/pq#section-documentation)
Dentro de la ruta para poder realizar la conexión debe reemplazar los siguientes, en este declaramos en constante y asi solo corregimos dentro ella
``` go
const (
	conPos = "postgres://(userdb):(password)@localhost:(puerto)/(nameschema)?sslmode=disable"
)
```
Para instanciar nuestra conexión se agrega el siguiente codigo
``` go
storage.NewPostgreDB()
```
#### Migrate de de product
Para poder correr la migración de product, implementamos el siguiente codigo
``` go
    storageProd := storage.NewPsqlProduct(storage.Pool())
	serviceProd := product.NewService(storageProd)

	if err := serviceProd.Migrate(); err != nil {
		log.Fatalf("product.Migrate(): %v", err)
	}
```

#### Insert en la tabla de product
Para poder creadar datos aplicamos el siguiente codigo, y reemplazar los valores que deseamos agregar a nuestra tabla
``` go
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
```