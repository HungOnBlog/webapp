FROM golang:1.19.4


WORKDIR /app

COPY go.mod go.sum ./

RUN go install -v ./...

COPY . .

RUN go build -o main.sh .

RUN chmod +x main.sh

CMD ["go", "run", "main.go"]