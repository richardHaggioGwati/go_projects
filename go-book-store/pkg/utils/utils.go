package utils

import (
	"encoding/json"
	"io"
	"net/http"
)


//Read the body of an HTTP request and convert the JSON into a struct
func ParseBody(r *http.Request, x interface{}){
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}