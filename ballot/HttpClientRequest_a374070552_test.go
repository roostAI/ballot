package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func httpClientRequest(method, url, path string, body io.Reader) (int, []byte, error) {
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

func TestHttpClientRequest_a374070552(t *testing.T) {
	t.Run("Successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(rw, "Hello, client")
		}))
		defer server.Close()

		statusCode, body, err := httpClientRequest("GET", server.URL, "/", nil)
		if err != nil {
			t.Error(err)
		}
		if statusCode != 200 {
			t.Errorf("Expected status code 200, got %d", statusCode)
		}
		if !strings.Contains(string(body), "Hello, client") {
			t.Errorf("Unexpected body: got %v", string(body))
		}
	})

	t.Run("Unsuccessful request - invalid operation", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(rw, "Hello, client")
		}))
		defer server.Close()

		_, _, err := httpClientRequest("INVALID", server.URL, "/", nil)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("Unsuccessful request - invalid url", func(t *testing.T) {
		_, _, err := httpClientRequest("GET", "http://invalid-url", "/", nil)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("Unsuccessful request - non 2xx status code", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			http.Error(rw, "Bad request", http.StatusBadRequest)
		}))
		defer server.Close()

		statusCode, _, err := httpClientRequest("GET", server.URL, "/", nil)
		if err != nil {
			t.Error(err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, statusCode)
		}
	})

	t.Run("Successful POST request with params", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.Method != http.MethodPost {
				t.Errorf("Expected method 'POST', got '%s'", req.Method)
			}

			body := new(bytes.Buffer)
			_, err := io.Copy(body, req.Body)
			if err != nil {
				t.Error(err)
			}
			req.Body.Close()

			if body.String() != "param1=value1&param2=value2" {
				t.Errorf("Expected body to be 'param1=value1&param2=value2', got '%s'", body.String())
			}

			fmt.Fprintln(rw, "Hello, client")
		}))
		defer server.Close()

		params := strings.NewReader("param1=value1&param2=value2")
		statusCode, body, err := httpClientRequest("POST", server.URL, "/", params)
		if err != nil {
			t.Error(err)
		}
		if statusCode != 200 {
			t.Errorf("Expected status code 200, got %d", statusCode)
		}
		if !strings.Contains(string(body), "Hello, client") {
			t.Errorf("Unexpected body: got %v", string(body))
		}
	})
}
