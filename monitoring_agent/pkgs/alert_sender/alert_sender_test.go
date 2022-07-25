package alert_sender

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestSendAlertSuccess(t *testing.T) {

	// load .env file from given path
  	// we keep it empty it will load .env from current directory
  	err := godotenv.Load(".env")

  	if err != nil {
    	t.Errorf("Error loading .env file")
  	}

	recipient := "	olubiyiontheweb@gmail.com"
	subject := "Test Alert"
	message := "This is a test alert"
	html := "<h1>This is a test alert</h1>"

	ses_creds := make(map[string]string)
  	ses_creds["aws_access_key"] = os.Getenv("AWS_ACCESS_KEY_ID")
  	ses_creds["aws_secret_key"] = os.Getenv("AWS_SECRET_ACCESS_KEY")
  	ses_creds["aws_region"] = os.Getenv("AWS_REGION")
  	ses_creds["charset"] = os.Getenv("CHARSET")
  	ses_creds["from_address"] = fmt.Sprintln("\"Watched System Down Alert\" <", os.Getenv("FROM_EMAIL"), ">")

	// send email to admin
	mail_res_code, _ := SendAlert(recipient, subject, message, html, ses_creds)

	time.Sleep(time.Second * 5)

	if mail_res_code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, mail_res_code)
	}
}

func TestSendAlertFail(t *testing.T) {

	// load .env file from given path
  	// we keep it empty it will load .env from current directory
  	err := godotenv.Load(".env")

  	if err != nil {
    	t.Errorf("Error loading .env file")
  	}

	recipient := "	fdvgfhdcvd-doesnotexist@gmail.com"
	subject := "Test Alert"
	message := "This is a test alert"
	html := "<h1>This is a test alert</h1>"

	ses_creds := make(map[string]string)
  	ses_creds["aws_access_key"] = os.Getenv("AWS_ACCESS_KEY_ID")
  	ses_creds["aws_secret_key"] = os.Getenv("AWS_SECRET_ACCESS_KEY")
  	ses_creds["aws_region"] = os.Getenv("AWS_REGION")
  	ses_creds["charset"] = os.Getenv("CHARSET")
  	ses_creds["from_address"] = fmt.Sprintln("\"Watched System Down Alert\" <", os.Getenv("FROM_EMAIL"), ">")

	// send email to admin
	mail_res_code, _ := SendAlert(recipient, subject, message, html, ses_creds)

	time.Sleep(time.Second * 5)

	if mail_res_code == http.StatusOK {
		t.Errorf("Did not expect %d, got %d", http.StatusOK, mail_res_code)
	}
}