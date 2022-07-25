package database

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestRecordEventSuccess(t *testing.T) {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		t.Errorf("Error loading .env file")
	}

	// credentials for aws elastic search
	es_creds := make(map[string]string)
	es_creds["aws_access_key"] = os.Getenv("AWS_ES_ACCESS_KEY_ID")
	es_creds["aws_secret_key"] = os.Getenv("AWS_ES_SECRET_ACCESS_KEY")
	es_creds["aws_region"] = os.Getenv("AWS_REGION")
	es_creds["aws_service"] = os.Getenv("AWS_SERVICE")
	es_creds["aws_es_endpoint"] = os.Getenv("AWS_ES_DOMAIN_ENDPOINT")

	event_data := `{ "event_type": "error_test", "status_code": "404", "message": "This is just a error test message for 404, relax", "endpoint": "server_error", "time": "` + time.Now().Format(time.RFC3339) + `"}`

	res_code, _ := RecordEvent(event_data, os.Getenv("ES_INDEX"), es_creds)

	// if res_code is between 4XX and 5XX, it is an error, so it should pass
	if res_code != http.StatusCreated {
		t.Errorf("Expected %d, got %d", http.StatusOK, res_code)
	}
}

func TestRecordEventFail(t *testing.T) {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		t.Errorf("Error loading .env file")
	}

	// credentials for aws elastic search
	es_creds := make(map[string]string)
	es_creds["aws_access_key"] = os.Getenv("AWS_ES_ACCESS_KEY_ID")
	es_creds["aws_secret_key"] = os.Getenv("AWS_ES_SECRET_ACCESS_KEY")
	es_creds["aws_region"] = os.Getenv("AWS_REGION")
	es_creds["aws_service"] = os.Getenv("AWS_SERVICE")
	es_creds["aws_es_endpoint"] = "http://localhost:9200"

	event_data := `{ "event_type": "error_test", "status_code": "404", "message": "This is just a error test message for 404, relax", "endpoint": "server_error", "time": "` + time.Now().Format(time.RFC3339) + `"}`

	res_code, _ := RecordEvent(event_data, os.Getenv("ES_INDEX"), es_creds)

	// if res_code is between 4XX and 5XX, it is an error, so it should pass
	if !(res_code >= 400 && res_code < 600) {
		t.Errorf("Did not expect %d, got %d", http.StatusOK, res_code)
	}	
}