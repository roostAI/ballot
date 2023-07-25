package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpClientRequest_8fc45b1eff(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != "/test" {
			t.Errorf("request url is wrong, got: %s want: %s", req.URL.String(), "/test")
		}
		rw.Write([]byte(`OK`))
	}))
	defer server.Close()

	statusCode, body, err := httpClientRequest("GET", server.URL, "/test", nil)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if statusCode != http.StatusOK {
		t.Errorf("Unexpected status code, got: %d, want: %d", statusCode, http.StatusOK)
	}
	if string(body) != "OK" {
		t.Errorf("Unexpected body, got: %s, want: %s", string(body), "OK")
	}
}

func TestHttpClientRequest_8fc45b1eff_Failure(t *testing.T) {
	statusCode, _, err := httpClientRequest("GET", "wrong address", "/test", nil)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !errors.Is(err, err.(*url.Error).Err) {
		t.Errorf("Unexpected error, got: %s, want: %s", err.Error(), "Failed to create HTTP request.")
	}
	if statusCode != http.StatusBadRequest {
		t.Errorf("Unexpected status code, got: %d, want: %d", statusCode, http.StatusBadRequest)
	}
}
