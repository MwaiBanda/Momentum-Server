FROM golang:1.26.2-alpine
WORKDIR /dir

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./out/dist .

CMD ["./out/dist"]
