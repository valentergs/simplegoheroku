package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT not defined")
	}

	address := os.Getenv("ADDRESS")
	if port == "" {
		panic("ADDRESS not defined")
	}

	http.HandleFunc("/", home)
	fmt.Printf("Servindo com GO na porta:%s:%s\n", address, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Bem-vindo à nossa HOMPAGE!"))
}
