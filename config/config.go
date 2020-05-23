package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	PORT      = "PORT"
	SECTOR_ID = "SECTOR_ID"
)

var (
	Port     int   = 9090
	SectorID int64 = 5
)

func Init() {
	ReadENV()
}

func ReadENV() {
	var err error

	portStr := os.Getenv(PORT)
	if portStr != "" {
		Port, err = strconv.Atoi(portStr)
		if err != nil {
			panic(fmt.Errorf("invalid port: %s", err.Error()))
		}
	}

	sectorStr := os.Getenv(SECTOR_ID)
	if sectorStr != "" {
		SectorID, err = strconv.ParseInt(sectorStr, 10, 64)
		if err != nil {
			panic(fmt.Errorf("invalid sector id: %s", err.Error()))
		}
	}
}
