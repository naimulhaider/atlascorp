package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/naimulhaider/atlascorp/model"
)

func TestDNSRequest_ToModel(t *testing.T) {
	tests := []struct {
		Input       DNSRequest
		OutputData  model.DNSRequest
		OutputError error
	}{
		{
			Input: DNSRequest{
				X:        "111.1",
				Y:        "222.2",
				Z:        "333.3",
				Velocity: "444.4",
			},
			OutputData: model.DNSRequest{
				X:        111.1,
				Y:        222.2,
				Z:        333.3,
				Velocity: 444.4,
			},
			OutputError: nil,
		},
		{
			Input: DNSRequest{
				X:        "abcd",
				Y:        "222.2",
				Z:        "333.3",
				Velocity: "444.4",
			},
			OutputData:  model.DNSRequest{},
			OutputError: fmt.Errorf("invalid value for x: strconv.ParseFloat: parsing \"abcd\": invalid syntax"),
		},
	}

	for _, tc := range tests {
		res, err := tc.Input.ToModel()

		if !reflect.DeepEqual(tc.OutputError, err) {
			t.Errorf("Expected error: %v, got: %v", tc.OutputError, err)
		}

		if !reflect.DeepEqual(tc.OutputData, res) {
			t.Errorf("Expected output: %v, got: %v", tc.OutputData, res)
		}
	}
}
