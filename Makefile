unit_test:
	go test ./...

integration-tests:
	go test ./tests/integrations -v -coverpkg=./...

docker-integration-tests:
	docker-compose -f docker-compose.ci.yml run --build --rm server make integration-tests

