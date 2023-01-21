# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
WORKDIR /app
COPY *.go ./
COPY *.mod ./
RUN go mod download
RUN go build -o /fly-mock
EXPOSE 8000
CMD [ "/fly-mock" ]