package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleDNS(t *testing.T) {
	tests := []struct {
		Input        string
		OutputStatus int
		OutputData   string
	}{
		{
			Input:        `{"x":"123.5","y":"234.5","z":"456.5","vel":"1000.5"}`,
			OutputStatus: http.StatusOK,
			OutputData:   `{"loc":5073}`,
		},
		{
			Input:        `{"x":"0.0","y":"0.0","z":"0.0","vel":"500.5"}`,
			OutputStatus: http.StatusOK,
			OutputData:   `{"loc":500.5}`,
		},
		{
			Input:        `{"vel":"500.5"}`,
			OutputStatus: http.StatusBadRequest,
			OutputData:   `{"error":"invalid value for x: strconv.ParseFloat: parsing \"\": invalid syntax"}`,
		},
		{
			Input:        `{"x":"0.0","y":"0.0","z":"0.0","vel":"abcd"}`,
			OutputStatus: http.StatusBadRequest,
			OutputData:   `{"error":"invalid value for vel: strconv.ParseFloat: parsing \"abcd\": invalid syntax"}`,
		},
	}

	for _, tc := range tests {
		req, err := http.NewRequest("POST", "/api/v1/dns", bytes.NewBuffer([]byte(tc.Input)))
		if err != nil {
			t.Error(err)
		}

		rec := httptest.NewRecorder()
		hdlr := http.HandlerFunc(HandleDNS)

		hdlr.ServeHTTP(rec, req)

		if rec.Code != tc.OutputStatus {
			t.Errorf("Expected status code: %d, got %d", tc.OutputStatus, rec.Code)
		}

		if rec.Body.String() != tc.OutputData {
			t.Errorf("Expected response: %s, got %s", tc.OutputData, rec.Body.String())
		}
	}
}

func TestServeJSON(t *testing.T) {
	tests := []struct {
		InputStatus  int
		InputData    interface{}
		OutputStatus int
		OutputData   string
	}{
		{
			InputStatus: http.StatusOK,
			InputData: DNSResponse{
				Loc: 5073,
			},
			OutputStatus: http.StatusOK,
			OutputData:   `{"loc":5073}`,
		},
		{
			InputStatus: http.StatusBadRequest,
			InputData: ErrorResponse{
				Err: "abcd",
			},
			OutputStatus: http.StatusBadRequest,
			OutputData:   `{"error":"abcd"}`,
		},
	}

	for _, tc := range tests {
		rec := httptest.NewRecorder()

		ServeJSON(rec, tc.InputStatus, tc.InputData)

		if rec.Code != tc.OutputStatus {
			t.Errorf("Expected status code: %d, got %d", tc.OutputStatus, rec.Code)
		}

		if rec.Body.String() != tc.OutputData {
			t.Errorf("Expected response: %s, got %s", tc.OutputData, rec.Body.String())
		}
	}
}
