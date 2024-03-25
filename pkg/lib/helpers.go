package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HttpRespErrorJson struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

// p := isProduction - checks if value passed is production
func GetBaseUrl(p bool) string {
	if p {
		return ProductionBaseURL
	}

	return SandboxBaseURL
}

// check http error occurance
func CheckHttpError(response *http.Response) error {
	if response.StatusCode < 200 || response.StatusCode > 299 {
		var errStr string

		var respErrorJson HttpRespErrorJson
		jsonErr := json.NewDecoder(response.Body).Decode(&respErrorJson)

		if jsonErr == nil {
			errStr = fmt.Sprintf("Error %d: %s", respErrorJson.StatusCode, respErrorJson.Error)
		}

		return errors.New(errStr)
	}
	return nil
}
