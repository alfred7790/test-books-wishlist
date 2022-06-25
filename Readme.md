# Test API REST
```shell
Autor: I.S.C. Edgar Alfred Rodriguez Robles
E-mail: alfred.7790@gmail.com
```
# Table of Contents
[_TOC_]

# Overview
This project is an example of API REST.

# Requirements
To run this project it is necessary to have installed:
- docker
- docker-compose
- go
- swaggo

# GOPATH is exported?
Make sure that yout GOPATH is exported.
```shell
$ echo $GOPATH
```
> To create the API documentation, swagger will be installed in your go directory ($GOPATH).

# Quick start - Using Makefile
> IMPORTANT! If you have a linux distribution, you will be able to perform this procedure, otherwise I suggest you try another way to run the service.
1. Clone the repo:
```shell
$ git clone git@github.com:alfred7790/test-books-wishlist.git
```
2. Open the project
```shell
$ cd test-books-wishlist
```
5. Build and run the service:
```shell
$ make
```
6. If everything is ok, you should see something like this:
```shell
...
[GIN-debug] Listening and serving HTTP on :8080
connected to 'test' database
```
7. Go to [swagger docs](http:localhost:8080/v1/swagger/index.html) and have fun.

# Quick start - Running Binary
1. Clone the repo:
```shell
$ git clone git@github.com:alfred7790/test-books-wishlist.git
```
2. Open the project
```shell
$ cd test-books-wishlist
```
3. Running the DB service:
```shell
$ docker-compose up -d booksdb
```
4. Get dependencies:
```shell
$ go mod tidy && go get -u github.com/swaggo/swag/cmd/swag
```
5. Build the service:
```shell
$ go build -o ./build/bin/main cmd/main.go
```
6. Running the service:
```shell
$ ./build/bin/main
```
7. If everything is ok, you should see something like this:
```shell
...
[GIN-debug] Listening and serving HTTP on :8080
connected to 'test' database
```
8. Go to [swagger docs](http:localhost:8080/v1/swagger/index.html) and have fun.

# Custom Config
> If you need to change the default values of the configuration.
1. Open the project.
```shell
$ cd test-books-wishlist
```
2. Copy the template of the configuration to `config.yml`.
```shell
$ cp ./cmd/config/config_template.yml config.yml
```
3. Edit it with your own values.
> WARNING! Make sure that if you edit the values about the `DB service`, also you should modify the `docker-compose.yml` file.
4. Restart the service `using Makefile` or `running Binary`.
5. Go to [swagger docs](http:localhost:8080/v1/swagger/index.html) and have fun.

# Tests
- Using `Makefile`.
```shell
$ make test
```
- Or try this:
```shell
$ go test ./...
```

# Up and Down DB service
- Using `Makefile`.
```shell
  $ make db
```
```shell
  $ make db-down
```
- Using `docker-compose`
```shell
$ docker-compose up -d booksdb
```
```shell
$ docker-compose down
```

# Dependencies
> Make sure that your GOPATH is exported.

If you get an error with swagger or with another package, try this:
```shell
$ make dep
```
Or
```shell
$ go mod tidy && go get -u github.com/swaggo/swag/cmd/swag
```

# Comments:
- The Database is not persistent.
- I didn't cover the whole project with tests.
- I include default values in the config structure to run the service more easily,
  but it is not good practice including database credentials in public files,
  these values should be placed in an .env file.
- I didn't include comments about functions because I didn't have enough time.
- I decided to use GORM because the migrations of structures within the database are much easier to do.