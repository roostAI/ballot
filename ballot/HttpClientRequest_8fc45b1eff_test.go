package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func httpClientRequest(method, url, path string, body *bytes.Buffer) (int, []byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url+path, body)
	if err != nil {
		return 0, nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, respBody, nil
}

func TestHttpClientRequest_8fc45b1eff(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte(`OK`))
		}))
		defer server.Close()

		statusCode, body, err := httpClientRequest("GET", server.URL, "/", nil)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		if statusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got: %d", http.StatusOK, statusCode)
		}

		if string(body) != "OK" {
			t.Errorf("Expected body %s, but got: %s", "OK", string(body))
		}
	})

	t.Run("failure case", func(t *testing.T) {
		_, _, err := httpClientRequest("GET", "http://invalid-url", "/", nil)
		if err == nil {
			t.Error("Expected error, but got none")
		}

		if !strings.Contains(err.Error(), "no such host") {
			t.Errorf("Expected error message to contain %q, but got: %v", "no such host", err)
		}
	})
}
