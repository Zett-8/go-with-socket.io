FROM golang:latest

RUN mkdir /app
WORKDIR /app
ADD . /app

RUN go get github.com/pilu/fresh
RUN go mod download

COPY . .

EXPOSE 8060