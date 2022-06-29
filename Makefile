PKG_LIST = `go list ./... | grep -v /vendor`
.PHONY: all
all: | db dep test build deploy

.PHONY: dep
dep:
	@echo "Download dependencies"
	@echo "GOPATH is:" $(GOPATH)
	go mod tidy
	go get -u github.com/swaggo/swag/cmd/swag

.PHONY: db
db:
	@echo "Running servicedb"
	docker-compose up -d servicedb
	docker-compose ps
	# need to make sure our db is up and available before we run tests.
	# 2 seconds
	sleep 2

.PHONY: build
build:
	@echo "Building image"
	docker build . -t googlebooksimg -f ./ias/docker/Dockerfile


.PHONY: test
test:
	@echo "Running tests"
	go test $(PKG_LIST)

.PHONY: deploy
deploy:
	@echo "Deploying google-books"
	docker-compose up -d service
	docker-compose ps

.PHONY: down
down:
	docker-compose down
