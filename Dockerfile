FROM golang:1.26 AS build-stage

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY go.sum ./

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /latihan_rest_api ./cmd/api

FROM gcr.io/distroless/base-debian11 AS build-realese-stage

WORKDIR /

COPY --from=build-stage /latihan_rest_api /latihan_rest_api

EXPOSE 1323

USER nonroot:nonroot

ENTRYPOINT [ "./latihan_rest_api" ]