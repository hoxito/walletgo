
FROM golang:1.18-alpine

ENV REDIS_URL host.docker.internal:49154
ENV MYSQL_URL root:jose1@tcp(host.docker.internal)/wallet

WORKDIR /go/src/walletgo
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...
EXPOSE 3010
CMD ["go" , "run" , "/go/src/walletgo"]