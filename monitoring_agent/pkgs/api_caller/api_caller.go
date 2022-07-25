package api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeRequest(api_url string, 
                path string, 
                auth_token string) (int, string) {
    
    // Make the request
	req, err := http.NewRequest(
        http.MethodGet,
        api_url + path,
        nil,
    )
    if err != nil {
        // return http unprocessed 422 error
        return http.StatusUnprocessableEntity, fmt.Sprintf("Error making request: %s", path)
    }

    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization", "Bearer "+auth_token)
    req.Header.Add("Content-Type", "application/json")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return http.StatusUnprocessableEntity, fmt.Sprintf("Error sending request: %s", path)
    }
    responseBytes, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return http.StatusUnprocessableEntity, fmt.Sprintf("Error reading response: %s", path)
    }
    
	return res.StatusCode, string(responseBytes)
}