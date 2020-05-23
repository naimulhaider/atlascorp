package main

import (
	"fmt"
	"net/http"

	"github.com/naimulhaider/atlascorp/api"
	"github.com/naimulhaider/atlascorp/config"
)

func main() {
	config.Init()

	http.HandleFunc("/api/v1/dns", api.HandleDNS)

	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
