build-docker-image:
	docker build --no-cache -t xuqingfeng/kong-go-plugin-geoip .

push-docker-image:
	docker push xuqingfeng/kong-go-plugin-geoip

fmt:
	go fmt ./...

tidy:
	go mod tidy

cleanup:
	docker stop kong-go-plugin-geoip && docker rm kong-go-plugin-geoip