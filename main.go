package main

import (
	"log"
	"net/http"
	"os"
	"github.com/Nanit20/project-go/storage"
	"github.com/Nanit20/project-go/router"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	db := storage.NewStorage()

	r := router.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Servidor corriendo en el puerto", port)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
