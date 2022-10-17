package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cassiusbessa/goCarShop/routes"
)

func main() {
	routes := routes.CarRoutes()
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}
