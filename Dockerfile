FROM golang:1.19.2

RUN mkdir -p /usr/src/goAPI
WORKDIR /usr/src/goAPI

COPY src/go.mod src/go.sum ./
RUN go mod download 
RUN go mod verify

COPY src/ ./

RUN go build -v -o ./main

CMD ["./main"]

# docker build -t aleben/itlandfill-cas-go_api:1.0 .