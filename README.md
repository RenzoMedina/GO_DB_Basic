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
