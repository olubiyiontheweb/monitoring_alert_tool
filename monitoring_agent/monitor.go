package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
  // load .env file from given path
  // we keep it empty it will load .env from current directory
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }
}