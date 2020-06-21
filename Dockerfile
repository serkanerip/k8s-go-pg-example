# build stage
FROM golang:alpine AS build-env
WORKDIR /src
ADD . .
RUN go get -d -v
RUN go build -o goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
ENTRYPOINT ./goapp