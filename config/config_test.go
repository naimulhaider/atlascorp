package config

import (
	"fmt"
	"os"
	"testing"
)

func TestReadENV(t *testing.T) {
	tests := []struct {
		Port     int
		SectorID int64
	}{
		{
			Port:     5050,
			SectorID: 123456,
		},
		{
			Port:     9090,
			SectorID: 0,
		},
	}

	for _, tc := range tests {
		os.Setenv(PORT, fmt.Sprintf("%d", tc.Port))
		os.Setenv(SECTOR_ID, fmt.Sprintf("%d", tc.SectorID))

		ReadENV()

		if Port != tc.Port {
			t.Errorf("Expected port: %d, got %d", tc.Port, Port)
		}

		if SectorID != tc.SectorID {
			t.Errorf("Expected sector id: %d, got %d", tc.SectorID, SectorID)
		}
	}
}
