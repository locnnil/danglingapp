## Build
FROM golang:alpine AS build

RUN apk update && apk add --no-cache git

WORKDIR /usr/local/go/src/refusekills

COPY go.mod ./

RUN go mod download && go mod verify && go mod tidy

COPY main.go/ ./main.go

RUN go build -mod=mod -o /refusekills main.go

## Deploy
FROM scratch

WORKDIR /

COPY --from=build /refusekills /refusekills

ENTRYPOINT ["/refusekills"]