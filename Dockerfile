FROM golang:1.19.4


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]