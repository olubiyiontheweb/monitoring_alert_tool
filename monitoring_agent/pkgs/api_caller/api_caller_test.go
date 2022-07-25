package api_caller

import (
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMakeRequestServerUp(t *testing.T) {
	// load .env file from given path
  	// we keep it empty it will load .env from current directory
  	err := godotenv.Load(".env")

  	if err != nil {
		t.Errorf("Error loading .env file")
  	}

	res_code, _ := MakeRequest(os.Getenv("API_URL"), "active_status", os.Getenv("AUTH_TOKEN"))
	if res_code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, res_code)
	}
}

func TestMakeRequestServerError(t *testing.T) {
	// load .env file from given path
  	// we keep it empty it will load .env from current directory
  	err := godotenv.Load(".env")

  	if err != nil {
		t.Errorf("Error loading .env file")
  	}

	res_code, _ := MakeRequest(os.Getenv("API_URL"), "inactive_status", os.Getenv("AUTH_TOKEN"))

	// if res_code is between 4XX and 5XX, it is an error, so it should pass
	if res_code >= 400 && res_code < 600 {
		t.Errorf("Did not expect %d, got %d", http.StatusOK, res_code)
	}
}

func TestMakeRequestServerESQuery(t *testing.T) {
	// load .env file from given path
  	// we keep it empty it will load .env from current directory
  	err := godotenv.Load(".env")

  	if err != nil {
		t.Errorf("Error loading .env file")
  	}

	res_code, _ := MakeRequest(os.Getenv("API_URL"), "elasticsearch_query", os.Getenv("AUTH_TOKEN"))

	// if elasticsearch query is successful, it should pass
	if res_code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, res_code)
	}
}

func TestMakeRequestServerDown(t *testing.T) {
	// load .env file from given path
  	// we keep it empty it will load .env from current directory
  	err := godotenv.Load(".env")

  	if err != nil {
		t.Errorf("Error loading .env file")
  	}

	res_code, _ := MakeRequest(os.Getenv("API_URL"), "active_status", os.Getenv("AUTH_TOKEN"))

	// is server is down, this is should to pass
	if !(res_code >= 400 && res_code < 600) {
		t.Errorf("Did not expect %d, got %d", http.StatusOK, res_code)
	}
}

