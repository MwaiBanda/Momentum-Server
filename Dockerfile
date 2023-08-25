FROM golang:1.20-alpine
WORKDIR /dir

COPY ./cms/. ./cms/.
COPY  ./cms/dist/. ./cms/dist/.
COPY ./cms/dist/index.html ./cms/dist/index.html


COPY ./prisma/schema.prisma ./prisma/schema.prisma
COPY ./controller/. ./controller/.
COPY ./docs/. ./docs/.
COPY ./httputil/. ./httputil/.
COPY ./model/. ./model/.


COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go run github.com/steebchen/prisma-client-go generate
COPY . .

RUN go build -o ./out/dist .

CMD ./out/dist
