FROM golang:1.19-alpine3.16

WORKDIR /usr/src/goAPI

COPY go.mod go.sum ./
RUN go mod download 
RUN go mod verify


COPY main.go ./
COPY endpoints ./endpoints/
COPY structs ./structs/
RUN go build -v -o ./main

CMD ["./main"]

