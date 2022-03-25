FROM golang:latest

RUN mkdir /app
WORKDIR /app

RUN export GO111MODULE=on
RUN go get github.com/joho/godotenv

RUN cd /app && git clone 

RUN go get

RUN



