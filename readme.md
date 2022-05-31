
# Wallet Go

This is a simple wallet application that stores users, their wallets, and transactions between wallets. This app uses a mysql database with the schema and some data sql script inside its folders. It is also monitored with prometheus and containerized with docker.


### Auth

Currently, the auth service is not implemented, you can login sending a request to the login url and it will set the default user for you.
## Requirements

Go 1.18  [golang.org](https://golang.org/doc/install)

## Initial Config
Set Environment Variables

    export GO111MODULE=on
    export GOFLAGS=-mod=vendor

To download the app correctly you must run:

    go get github.com/hoxito/walletgo

Once downloaded you will have the code in folder

    cd $GOPATH/src/github.com/hoxito/walletgo

Or you could just run

    https://github.com/hoxito/walletgo.git

To have the repository

# Installation

To install the application in your local machine:

1- Install required libs

    go mod download
	go mod vendor


2- build and execution

    go install
    walletgo


## Mysql

Users, wallets and transactions are stored in a mysql database [mysql](https://www.mysql.com/downloads/)
The schema is described in the next image:
```
![walletgo schema](https://github.com/hoxito/walletgo/blob/main/docker/mysql/dbschema.jpg?raw=true)
```
To load the schema you must run your local instance of mysql and execute the sql file located in /walletgo/docker/mysql/schema.sql


## API-DOCS
To load api docs you have to run the application and go to

> (http://localhost:3010/swagger/index.html)


## Run Docker Containers
If you wish to run the application in docker containers you can do so with:

    cd /walletgo
    docker-compose build
    docker-compose up
This will create 3 docker images with their respective containers and run the containers. This containers are Mysqlc, golangApp and PrometheusAPI.


## Instrumentation
To get the aplication scrapped metrics you can start the prometheus container located in walletgo/docker/prometheus/. and visit

> (http://localhost:9090)

You should see the prometheus panel and interact with it. The walletgo App exposes metrics at
> (http://localhost:3010/metrics)


## Testing
To run tests, you must cd into the desired test folder and run 

    go test -v
