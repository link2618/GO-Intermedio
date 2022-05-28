package main

import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

// go mod init <NombreModulo> -> crear modulo
// go get <NombreModulo> -> traer modulo
// go mod tidy -> eliminar modulos que no se estan usando
// go mod download -json -> información de los módulos cacheados
func main() {
	hellomod.SayHello()
	hellomod2.SayHello("Version 2")
}
