# syntax=docker/dockerfile:1

FROM golang:1.18
WORKDIR /app
COPY *.go ./
COPY /config/*.go ./config/
COPY /delivery/*.go ./delivery/
COPY /gateways/*.go ./gateways/
COPY /model/*.go ./model/
COPY /repositories/*.go ./repositories/
COPY /usecase/*.go ./usecase/
COPY /usecase/ports/*.go ./usecase/ports/
COPY /util/*.go ./util/
COPY *.conf ./
COPY *.mod ./
RUN go mod download
RUN go get github.com/jbl1108/goFly/config
RUN go get github.com/jbl1108/goFly/delivery
RUN go get github.com/jbl1108/goFly/gateways
RUN go get github.com/jbl1108/goFly/model
RUN go get github.com/jbl1108/goFly/repositories
RUN go get github.com/jbl1108/goFly/usecase
RUN go get github.com/jbl1108/goFly/util
RUN go get github.com/jbl1108/goFly/usecase/ports

RUN ls -al
RUN go build -o goFly
CMD [ "goFly" ]