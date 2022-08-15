package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

const (
	Version  = "0.1"
	Priority = 1
)

type Config struct {
	GeoIPDB string
}

func New() interface{} {
	return &Config{}
}

func main() {
	server.StartServer(New, Version, Priority)
}

func (Config) Access(kong *pdk.PDK) {

	// read GeoIP db

	// get country code

	// append country info in header

	// send to upstream
}
