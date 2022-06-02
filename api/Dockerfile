FROM golang:1.18-buster AS build

WORKDIR /app

COPY ./api/src .
RUN go mod download
RUN go build -o /docker-gs-ping

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-gs-ping /docker-gs-ping

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]