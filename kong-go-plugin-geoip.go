package main

import (
	"net"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	"github.com/oschwald/maxminddb-golang"
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

	ip, err := kong.Client.GetForwardedIp()
	if err != nil {
		kong.Log.Err(err.Error())
	}

	// get country code
	geoIPHeaders, err := lookupGeoInfoFromDB(ip, conf.Db_path)
	if err != nil {
		kong.Log.Err(err.Error())
	}
	// append country info in request header and send to upstream
	kong.ServiceRequest.SetHeader("X-Country-Code", geoIPHeaders.Country.ISOCode)
	kong.ServiceRequest.SetHeader("X-City-Name", geoIPHeaders.City.Names["en"])

	// check if echo back to client
	if conf.Echo_down_stream {
		kong.Response.SetHeader("X-Country-Code", geoIPHeaders.Country.ISOCode)
		kong.Response.SetHeader("X-City-Name", geoIPHeaders.City.Names["en"])
	}
}

type GeoIPHeaders struct {
	City struct {
		Names map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Country struct {
		ISOCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`
}

// FIXME: load from plugin init
func lookupGeoInfoFromDB(ip string, path string) (GeoIPHeaders, error) {

	var record GeoIPHeaders

	db, err := maxminddb.Open(path)
	if err != nil {
		return record, err
	}
	defer db.Close()

	parsedIP := net.ParseIP(ip)

	err = db.Lookup(parsedIP, &record)
	if err != nil {
		return record, err
	}
	return record, nil
}
