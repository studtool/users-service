FROM golang:1.12 as base
WORKDIR /tmp/users-service
COPY . .
RUN go build -mod vendor -o /tmp/service .

FROM ubuntu:18.04
WORKDIR /tmp
COPY --from=base /tmp/service ./service
ENTRYPOINT ./service
EXPOSE 80
