package api

import (
	"fmt"
	"strconv"

	"github.com/naimulhaider/atlascorp/model"
)

type DNSRequest struct {
	X        string `json:"x"`
	Y        string `json:"y"`
	Z        string `json:"z"`
	Velocity string `json:"vel"`
}

func (d DNSRequest) ToModel() (model.DNSRequest, error) {
	m := model.DNSRequest{}

	x, err := strconv.ParseFloat(d.X, 64)
	if err != nil {
		return m, fmt.Errorf("invalid value for x: %s", err.Error())
	}

	y, err := strconv.ParseFloat(d.Y, 64)
	if err != nil {
		return m, fmt.Errorf("invalid value for y: %s", err.Error())
	}

	z, err := strconv.ParseFloat(d.Z, 64)
	if err != nil {
		return m, fmt.Errorf("invalid value for z: %s", err.Error())
	}

	vel, err := strconv.ParseFloat(d.Velocity, 64)
	if err != nil {
		return m, fmt.Errorf("invalid value for vel: %s", err.Error())
	}

	m.X, m.Y, m.Z, m.Velocity = x, y, z, vel

	return m, nil
}

type DNSResponse struct {
	Loc float64 `json:"loc"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Err: err.Error(),
	}
}
