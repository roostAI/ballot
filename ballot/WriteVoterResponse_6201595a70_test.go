package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Status struct {
	Code    int
	Message string
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(status)
	if err != nil {
		log.Println("error marshaling response to vote request. error: ", err)
	}
	w.Write(resp)
}

func TestWriteVoterResponse_6201595a70(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rr := httptest.NewRecorder()
		writeVoterResponse(rr, Status{200, "Vote received"})

		result := rr.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, result.StatusCode)
		}

		var status Status
		json.NewDecoder(result.Body).Decode(&status)

		if status.Code != 200 || status.Message != "Vote received" {
			t.Errorf("Unexpected body: got %+v", status)
		}
	})

	t.Run("marshal_error", func(t *testing.T) {
		rr := httptest.NewRecorder()
		writeVoterResponse(rr, Status{200, string([]byte{0x80})})

		result := rr.Result()
		defer result.Body.Close()

		if result.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, result.StatusCode)
		}

		if rr.Body.String() != "" {
			t.Errorf("Expected empty body, got %s", rr.Body.String())
		}
	})
}
