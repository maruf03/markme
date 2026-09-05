FROM golang:1.27 as base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# generate files here

FROM base as dev-build
RUN CGO_ENABLED=0 go build -o /app/api ./cmd/api/main.go

FROM base as prod-build
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app/api ./cmd/api/main.go

FROM dev-build as dev
EXPOSE 8080
ENTRYPOINT [ "/app/api" ]

FROM gcr.io/distroless/static-debian13
COPY --from=prod-build /app/api /api
EXPOSE 8080
ENTRYPOINT [ "api" ]
