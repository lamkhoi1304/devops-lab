# build stage
FROM golang:1.18-alpine as build-stage

WORKDIR /build

COPY .  .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

# production stage
FROM alpine

COPY --from=build-stage /build/app .

CMD ["./app"]