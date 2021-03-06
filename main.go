package main

import (
	"log"

	"github.com/andersondalmina/golang-rpc/cmd"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}
