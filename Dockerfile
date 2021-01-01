FROM node:lts-alpine AS build-js

COPY tarareba-frontend/package*.json ./

RUN yarn

COPY ./tarareba-frontend ./

RUN yarn build

FROM golang:alpine AS build-go

COPY ./go.* ./
COPY ./main.go ./

COPY tarareba-algorithms ./tarareba-algorithms
COPY tarareba-competition-history ./tarareba-competition-history
COPY tarareba-bff ./tarareba-bff

RUN CGO_ENABLED=0 go build -o tarareba-algorithms tarareba-algorithms/server/main.go
RUN CGO_ENABLED=0 go build -o tarareba-competition-history tarareba-competition-history/server/main.go
RUN CGO_ENABLED=0 go build -o tarareba-bff tarareba-bff/server.go
RUN CGO_ENABLED=0 go build -o main main.go

FROM ubuntu:latest

COPY --from=build-js build ./tarareba-frontend/build
COPY --from=build-go tarareba-algorithms ./tarareba-algorithms
COPY --from=build-go tarareba-competition-history ./tarareba-competition-history
COPY --from=build-go tarareba-bff ./tarareba-bff
COPY --from=build-go main ./main
COPY start.sh start.sh

EXPOSE 1213

RUN chmod +x start.sh

ENTRYPOINT ["./start.sh"]

# CMD ./tarareba-algorithms & ./tarareba-competition-history & ./tarareba-bff & ./main
