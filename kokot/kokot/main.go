package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/kokot/kokot/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 5000...")
	log.Fatal(http.ListenAndServe(":5000", r))
}
