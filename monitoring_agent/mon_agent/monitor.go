package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/olubiyiontheweb/monitoring_alert_tool/pkgs/alert_sender"
	"github.com/olubiyiontheweb/monitoring_alert_tool/pkgs/api_caller"
	"github.com/olubiyiontheweb/monitoring_alert_tool/pkgs/database"
)

func main() {
  log.SetPrefix("Monitoring agent: ")
	log.SetFlags(0)

  // recording the start time of the program
  start := time.Now()
  fmt.Print("\n")
  log.Println("Starting monitoring agent...")

  // load .env file from given path
  // we keep it empty it will load .env from current directory
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  // credentials for aws elastic search
  es_creds := make(map[string]string)
  es_creds["aws_access_key"] = os.Getenv("AWS_ES_ACCESS_KEY_ID")
  es_creds["aws_secret_key"] = os.Getenv("AWS_ES_SECRET_ACCESS_KEY")
  es_creds["aws_region"] = os.Getenv("AWS_REGION")
  es_creds["aws_service"] = os.Getenv("AWS_SERVICE")
  es_creds["aws_es_endpoint"] = os.Getenv("AWS_ES_DOMAIN_ENDPOINT")

  // credentials for aws
  ses_creds := make(map[string]string)
  ses_creds["aws_access_key"] = os.Getenv("AWS_ACCESS_KEY_ID")
  ses_creds["aws_secret_key"] = os.Getenv("AWS_SECRET_ACCESS_KEY")
  ses_creds["aws_region"] = os.Getenv("AWS_REGION")
  ses_creds["charset"] = os.Getenv("CHARSET")
  ses_creds["from_address"] = fmt.Sprintln("\"Watched System Down Alert\" <", os.Getenv("FROM_EMAIL"), ">")

  recipient := "	olubiyiontheweb@gmail.com"

  // call the apis, if status code is 4XX or 5XX, then log the error in elastic search and send email to admin
  // recipients := []string{recipient}

  available_endpoints := []string{}
  for _, v := range strings.Split(os.Getenv("API_ENDPOINTS"), ",") {
    available_endpoints = append(available_endpoints, v)
  }

  for i := 0; i < len(available_endpoints); i++ {
    // check the status of the api
    log.Println("Checking endpoint: ", available_endpoints[i])
    status_code, message := api_caller.MakeRequest(os.Getenv("API_URL"), available_endpoints[i], os.Getenv("AUTH_TOKEN"))
    var api_status string 

    // if status code is 4XX or 5XX, then log the error in elastic search and send email to admin
    if (status_code >= 400 && status_code < 600) {
      api_status = "error"
    } else {
      api_status = "ok"
    }
    // log the error in elastic search
    event_details := `{ "event_type": "` + api_status + `", "status_code": "` + strconv.Itoa(status_code) + `", "message": "` + message[12:len(message)-2] + `", "endpoint": "` + available_endpoints[i] + `", "time": "` + time.Now().Format(time.RFC3339) + `"}`
    fmt.Print(event_details + "\n")
                          
    es_status_code, es_message := database.RecordEvent(event_details, os.Getenv("ES_INDEX"), es_creds)

    if (es_status_code >= 400 && es_status_code < 600) {

      log.Fatalf("Error: %s", es_message)

    }

    // only send the email if the status code is 4XX or 5XX
    if (status_code >= 400 && status_code < 600) {
      
      // delay for 5 seconds before sending the email
      log.Println("Sleeping for 5 seconds before sending the email alert")
      time.Sleep(time.Second * 5)

      // The HTML body for the email.
      html :=  
      "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
        "<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
        "<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

      //The email body for recipients with non-HTML email clients.
      text := "Quodity system is down. Please check the system and resolve the issue.\n\n" + message

      // send email to admin
      mail_res_code, mail_res_message := alert_sender.SendAlert(recipient, os.Getenv("SUBJECT"), text, html, ses_creds)

      if (mail_res_code >= 200 && mail_res_code < 300) {
        log.Println("Email sent to admin")
      } else {
        log.Fatalf("Error: %s", mail_res_message)
      }
    }

    // wait for 5 seconds before making the next request
    log.Println("Sleeping for 5 seconds before making the next request\n")
    time.Sleep(time.Second * 5)
  }

  time_elapsed := time.Since(start).Seconds()
  log.Printf("Monitoring operation completed successfully after %d seconds \n", int(time_elapsed))
}