unit_test:
	go test ./...

integration-tests:
	go test ./tests/integration -v -coverpkg=./...

docker-integration-tests:
	DOCKER_BUILDKIT=1 docker-compose -f docker-compose.ci.yml run server sh -c "make integration-tests"

docker-unit-tests:
	DOCKER_BUILDKIT=1 docker-compose -f docker-compose.ci.yml run --build -e ENV="" --rm server make unit_test \
	&& docker-compose -f docker-compose.ci.yml  down

acceptance-tests:
	go test ./tests/acceptance -v -coverpkg=./...

docker-acceptance-tests:
	DOCKER_BUILDKIT=1 docker-compose -f docker-compose.ci.yml run --build --rm server make acceptance-tests \
	&& docker-compose -f docker-compose.ci.yml  down