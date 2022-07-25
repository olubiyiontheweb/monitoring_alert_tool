package database

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

// record event in AWS Elasticsearch Service, providing the index(table name), document ID, and the event data
func RecordEvent(event_data string, 
                index string,
                creds map[string]string) (int, string) {

  // Basic information for the Amazon OpenSearch Service domain
  endpoint := creds["aws_es_endpoint"] + "/" + index + "/" + "_doc"

  // document to saved in elastic search. Json format received here.
  body := strings.NewReader(event_data)

  // Get creds from environment variables and create the Signature Version 4 signer
  aws_creds := credentials.NewStaticCredentials(creds["aws_access_key"], creds["aws_secret_key"], "")
  signer := v4.NewSigner(aws_creds)

  // An HTTP client for sending the request
  client := &http.Client{}

  // Form the HTTP request  
  req, err := http.NewRequest(http.MethodPost, endpoint, body)
  if err != nil {
    return http.StatusUnprocessableEntity, fmt.Sprintf("Error creating request: %s", err)
  }

  // You can probably infer Content-Type programmatically, but here, we just say that it's JSON
  req.Header.Add("Content-Type", "application/json")

  // Sign the request, send it, and print the response
  signer.Sign(req, body, creds["aws_service"], creds["aws_region"], time.Now())
  resp, err := client.Do(req)
  if err != nil {
    return http.StatusUnprocessableEntity, fmt.Sprintf("Error sending request: %s", err)
  }

  return resp.StatusCode, resp.Status
}