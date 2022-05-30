
FROM golang:1.18-alpine

ENV REDIS_URL host.docker.internal:6379
ENV RABBIT_URL amqp://host.docker.internal
ENV MYSQL_URL host.docker.internal:3306

WORKDIR /go/src/walletgo
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...
EXPOSE 3010
CMD ["go" , "run" , "/go/src/statsv0"]