FROM golang:1.20-alpine
WORKDIR /Momentum
COPY go.mod .
COPY go.sum .
COPY prisma ./prisma/
RUN go mod download
RUN go run github.com/steebchen/prisma-client-go generate
COPY . .
RUN go build -o ./out/dist .
CMD ./out/dist
