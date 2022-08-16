build-docker-image:
	docker build --no-cache -t kong-go-plugin-geoip .

fmt:
	go fmt ./...

tidy:
	go mod tidy