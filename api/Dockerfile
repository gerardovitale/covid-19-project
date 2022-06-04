FROM golang:1.18-buster AS build

WORKDIR /app

COPY ./api/src .
RUN go mod download
RUN go build -o /covid-api

FROM gcr.io/distroless/base-debian10

ARG MONGO_PASS
ENV MONGO_PASS=${MONGO_PASS}

WORKDIR /

COPY --from=build /covid-api /covid-api

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/covid-api"]