FROM golang:1.22

WORKDIR /src
COPY . .

ENV GO111MODULE=on

RUN go build -o /bin/dir-hash cmd/directory-hash/main.go

ENTRYPOINT ["/bin/dir-hash"]
