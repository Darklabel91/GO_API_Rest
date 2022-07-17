package main

import (
	"fmt"
	"github.com/Darklabel91/GO_API_Rest/database"
	"github.com/Darklabel91/GO_API_Rest/routes"
)

func main() {
	database.DBConnect()
	fmt.Println("Iniciando")
	routs.HandleRequest()
}
