FROM golang:alpine AS build-go

WORKDIR /app

COPY ./go.* ./

COPY ./server ./server

COPY ./service ./service

COPY ./tarareba_competition_history_pb ./tarareba_competition_history_pb

COPY ./proto ./proto

RUN CGO_ENABLED=0 go build -o main ./server/main.go

CMD ./main
