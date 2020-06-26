test:
	docker-compose run --rm app go test ./src/... -coverpkg ./src/... -coverprofile cover.out
	docker-compose run --rm app go tool cover -html cover.out -o cover.html
