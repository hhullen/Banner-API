package main

import (
	"log"

	"main/infrastructure/controllers"
)

func main() {
	ctrl := controllers.NewControllerREST()
	log.Fatal(ctrl.Serve(":8080"))
}
