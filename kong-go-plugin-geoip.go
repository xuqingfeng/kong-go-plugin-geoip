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
	Db_path          string
	Echo_down_stream bool
}

func New() interface{} {
	return &Config{}
}

func main() {
	server.StartServer(New, Version, Priority)
}

func (conf Config) Access(kong *pdk.PDK) {

	// read GeoIP db
	ip, err := kong.Client.GetForwardedIp()
	if err != nil {
		kong.Log.Err(err.Error())
	}

	// get country code
	countryCode, err := lookupGeoInfoFromDB(ip, conf.Db_path)
	if err != nil {
		kong.Log.Err(err.Error())
	}
	// append country info in request header and send to upstream
	kong.ServiceRequest.SetHeader("X-Country-Code", countryCode)

	// check if echo back to client
	if conf.Echo_down_stream {
		kong.Response.SetHeader("X-Country-Code", countryCode)
	}
}

// FIXME: load from plugin init
func lookupGeoInfoFromDB(ip string, path string) (string, error) {

	db, err := geoip2.Open(path)
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
