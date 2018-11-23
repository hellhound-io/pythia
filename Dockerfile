# build stage
FROM golang:1.11-alpine AS build-stage
RUN apk update
RUN apk add --no-cache git gcc musl-dev
WORKDIR /bin
COPY . .
RUN go build -o /pythia


# final stage
FROM alpine
COPY --from=build-stage /pythia /pythia
EXPOSE 8080
CMD ["/pythia"]
