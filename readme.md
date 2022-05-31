
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

