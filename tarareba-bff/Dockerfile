FROM golang:alpine AS build-go

WORKDIR /app

COPY ./go.* ./

COPY ./*.go ./

COPY ./tarareba_competition_history_pb ./tarareba_competition_history_pb

COPY ./tarareba_algorithms_pb ./tarareba_algorithms_pb

COPY ./graph ./graph

RUN CGO_ENABLED=0 go build -o server ./server.go

CMD ./server
