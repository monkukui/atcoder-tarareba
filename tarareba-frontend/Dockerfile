FROM mhart/alpine-node:10 AS builder
WORKDIR /app
COPY package.json .
COPY package-lock.json .
COPY src .
COPY tsconfig.json .
COPY yarn.lock .

RUN yarn install
COPY . .
RUN yarn global add serve

EXPOSE 5000

CMD serve -s build
