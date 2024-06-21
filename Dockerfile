FROM golang:1.22

WORKDIR /usr/local/bin/

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o ./cetu ./cmd/...

EXPOSE 4000
CMD ["cetu"]
