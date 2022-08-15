package main

import (
	"net"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	"github.com/oschwald/geoip2-golang"
)

const (
	Version  = "0.1"
	Priority = 1
)

type Config struct {
	geo_ip_db_path   string
	echo_down_stream bool
}

func New() interface{} {
	return &Config{}
}

func main() {
	server.StartServer(New, Version, Priority)
}

func (Config) Access(kong *pdk.PDK) {

	// read GeoIP db
	ip, err := kong.Client.GetIp()
	if err != nil {
		kong.Log.Err(err.Error())
	}

	// get country code
	countryCode, err := lookupGeoInfoFromDB(ip)
	if err != nil {
		kong.Log.Err(err.Error())
	}
	// append country info in request header and send to upstream
	kong.ServiceRequest.SetHeader("X-Country-Code", countryCode)

	// check if echo back to client
}

func lookupGeoInfoFromDB(ip string) (string, error) {

	db, err := geoip2.Open("/data/geoip.dat")
	if err != nil {
		return "", err
	}
	defer db.Close()

	parsedIP := net.ParseIP(ip)
	record, err := db.City(parsedIP)
	if err != nil {
		return "", err
	}
	return record.Country.IsoCode, nil
}
