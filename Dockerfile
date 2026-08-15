FROM golang:1.24 AS test

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN go test -race ./...
RUN LIBRARY_PACKAGES=$(go list ./... | grep -Ev '/(cmd|docs|examples)(/|$)') && \
    go test -race -coverprofile=coverage.out -covermode=atomic $LIBRARY_PACKAGES
