package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eranamarante/go-react-expense-tracker/server/router"
)


func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	
	r := router.Router()

	fmt.Printf("Starting server on the port %v ...", port)
	log.Fatal(http.ListenAndServe(":" + port, r))
}