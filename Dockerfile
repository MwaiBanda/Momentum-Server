FROM node:24-alpine AS cms
WORKDIR /cms

RUN corepack enable && corepack prepare yarn@1.22.22 --activate

COPY cms/package.json cms/yarn.lock ./
RUN yarn install --frozen-lockfile --ignore-scripts

COPY cms/ ./
RUN yarn build

FROM golang:1.26.2-alpine AS backend-build
WORKDIR /dir

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY --from=cms /cms/dist ./cms/dist

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/momentum-server .

FROM alpine:3.22 AS runtime

RUN apk add --no-cache ca-certificates \
	&& addgroup -S momentum \
	&& adduser -S -G momentum momentum

WORKDIR /app
COPY --from=backend-build /out/momentum-server ./momentum-server
COPY --from=backend-build /dir/cms/dist ./cms/dist

USER momentum
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=5s --start-period=15s --retries=3 \
	CMD wget --spider --quiet http://127.0.0.1:8080/healthz || exit 1

ENTRYPOINT ["./momentum-server"]
