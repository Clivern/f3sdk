// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package f3sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// Output type
type Output struct {
	Response   string              `json:"response,omitempty"`
	StatusCode int                 `json:"statusCode,omitempty"`
	Headers    map[string][]string `json:"headers,omitempty"`
	Method     string              `json:"method,omitempty"`
	RawQuery   string              `json:"rawQuery,omitempty"`
}

// ConvertToJSON convert object to json
func (o *Output) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&o)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// ServerMock mocks http server for testing purposes
func ServerMock(uri, response string, statusCode int, debug bool) *httptest.Server {
	handler := http.NewServeMux()

	handler.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		output := &Output{
			Response:   response,
			StatusCode: statusCode,
			Headers:    r.Header,
			Method:     r.Method,
			RawQuery:   r.URL.RawQuery,
		}

		result, _ := output.ConvertToJSON()

		w.WriteHeader(statusCode)

		if debug {
			w.Write([]byte(result))
		} else {
			w.Write([]byte(response))
		}
	})

	srv := httptest.NewServer(handler)

	return srv
}
