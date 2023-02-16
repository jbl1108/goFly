# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
WORKDIR /app
COPY *.go ./
COPY /driver/*.go ./driver/
COPY /usecase/*.go ./usecase/
COPY /restservice/*.go ./restservice/
COPY /util/*.go ./util/
COPY *.conf ./
COPY *.mod ./
RUN go mod download
RUN go get github.com/jbl1108/goFly/driver
RUN go get github.com/jbl1108/goFly/usecase
RUN go get github.com/jbl1108/goFly/restservice
RUN ls -al
RUN go build -o goFly
CMD [ "goFly" ]