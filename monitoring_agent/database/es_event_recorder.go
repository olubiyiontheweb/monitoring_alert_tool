package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/joho/godotenv"
)

func main() {
  // load .env file from given path
  // we keep it empty it will load .env from current directory
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  // Basic information for the Amazon OpenSearch Service domain
  //domain := "quodity" // e.g. https://my-domain.region.es.amazonaws.com
  index := "my-index"
  id := "1"
  endpoint := os.Getenv("AWS_ES_DOMAIN_ENDPOINT") + "/" + index + "/" + "_doc" + "/" + id
  region := os.Getenv("AWS_REGION")
  service := "es"

  // Sample JSON document to be included as the request body
  json := `{ "title": "Thor: Ragnarok", "director": "Taika Waititi", "year": "2017" }`
  body := strings.NewReader(json)

  // Get credentials from environment variables and create the Signature Version 4 signer
  credentials := credentials.NewStaticCredentials(os.Getenv("AWS_ES_ACCESS_KEY_ID"), os.Getenv("AWS_ES_SECRET_ACCESS_KEY"), "")
  signer := v4.NewSigner(credentials)

  // An HTTP client for sending the request
  client := &http.Client{}

  // Form the HTTP request
  req, err := http.NewRequest(http.MethodPut, endpoint, body)
  if err != nil {
    fmt.Print(err)
  }

  // You can probably infer Content-Type programmatically, but here, we just say that it's JSON
  req.Header.Add("Content-Type", "application/json")

  // Sign the request, send it, and print the response
  signer.Sign(req, body, service, region, time.Now())
  resp, err := client.Do(req)
  if err != nil {
    fmt.Print(err)
  }
  fmt.Print(resp.Status + "\n")
}