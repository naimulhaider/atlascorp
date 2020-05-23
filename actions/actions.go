package actions

func CalculateLocation(sectorID int64, x, y, z, velocity float64) float64 {
	return x*float64(sectorID) + y*float64(sectorID) + z*float64(sectorID) + velocity
}
