package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	// godotenv package
	dotenv := goDotEnvVariable("SQL_PASSWORD")

	fmt.Println("Nice program running in a container")
	fmt.Printf("passwort:  %s", dotenv)
}
