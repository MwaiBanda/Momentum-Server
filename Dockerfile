FROM node:24-alpine AS cms
WORKDIR /cms

RUN corepack enable && corepack prepare yarn@1.22.22 --activate

COPY cms/package.json cms/yarn.lock ./
RUN yarn install --frozen-lockfile --ignore-scripts

COPY cms/ ./
RUN yarn build

FROM golang:1.26.2-alpine AS backend
WORKDIR /dir

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY --from=cms /cms/dist ./cms/dist

RUN go build -o ./out/dist .

CMD ["./out/dist"]
