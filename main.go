package main

import (
	"caloteiros/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
	fmt.Println("Servidor rodando em localhost:8080")
}
