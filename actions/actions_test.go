package actions

import "testing"

func TestCalculateLocation(t *testing.T) {
	tests := []struct {
		sectorID int64
		x        float64
		y        float64
		z        float64
		velocity float64
		location float64
	}{
		{
			sectorID: 5,
			x:        10,
			y:        20,
			z:        30,
			velocity: 2,
			location: 302.0,
		},
		{
			sectorID: 10,
			x:        10,
			y:        20,
			z:        30,
			velocity: 100,
			location: 700.0,
		},
		{
			sectorID: 0,
			x:        10,
			y:        20,
			z:        30,
			velocity: 5000,
			location: 5000,
		},
	}

	for _, tc := range tests {
		loc := CalculateLocation(tc.sectorID, tc.x, tc.y, tc.z, tc.velocity)
		if tc.location != loc {
			t.Errorf("Expected: %v, got: %v", tc.location, loc)
		}
	}
}
