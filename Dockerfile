FROM node:lts-alpine AS build-js

WORKDIR /app

COPY tarareba-frontend/package*.json ./

RUN yarn

COPY ./tarareba-frontend ./

RUN yarn build

FROM golang:alpine AS build-go

WORKDIR /app

COPY ./go.* ./
COPY ./main.go ./

COPY tarareba-algorithms ./tarareba-algorithms
COPY tarareba-competition-history ./tarareba-competition-history
COPY tarareba-bff ./tarareba-bff

RUN CGO_ENABLED=0 go build -o tarareba-algorithms tarareba-algorithms/server/main.go
RUN CGO_ENABLED=0 go build -o tarareba-competition-history tarareba-competition-history/server/main.go
RUN CGO_ENABLED=0 go build -o tarareba-bff tarareba-bff/server.go
RUN CGO_ENABLED=0 go build -o main main.go

FROM busybox

WORKDIR /app

COPY --from=build-js /app/build ./tarareba-frontend/build

COPY --from=build-go /app/tarareba-algorithms ./tarareba-algorithms
COPY --from=build-go /app/tarareba-competition-history ./tarareba-competition-history
COPY --from=build-go /app/tarareba-bff ./tarareba-bff
COPY --from=build-go /app/main ./main

EXPOSE 1213

CMD ./tarareba-algorithms
CMD ./tarareba-competition-history
CMD ./tarareba-bff
# CMD ./main
