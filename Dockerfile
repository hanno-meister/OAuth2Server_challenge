FROM golang

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build main.go

ENTRYPOINT ["./main"]