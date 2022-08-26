build-and-push-docker-image:
	docker buildx build --push --platform linux/arm64,linux/amd64 -t xuqingfeng/kong-go-plugin-geoip .

fmt:
	go fmt ./...

tidy:
	go mod tidy

cleanup:
	docker stop kong-go-plugin-geoip && docker rm kong-go-plugin-geoip