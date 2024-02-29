# syntax=docker/dockerfile:1
FROM golang:1.18-alpine AS build
# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
RUN apk add --no-cache git
RUN mkdir -p /go/src/github.com/jbl1108/goFly/goFetch/
COPY go.mod /go/src/github.com/jbl1108/goFly/goFetch/
COPY *.go /go/src/github.com/jbl1108/goFly/goFetch/
COPY ./config/*.go /go/src/github.com/jbl1108/goFly/goFetch/config/
COPY ./delivery/*.go /go/src/github.com/jbl1108/goFly/goFetch/delivery/
COPY ./gateways/*.go /go/src/github.com/jbl1108/goFly/goFetch/gateways/
COPY ./repositories/*.go /go/src/github.com/jbl1108/goFly/goFetch/repositories/
COPY ./usecase/*.go /go/src/github.com/jbl1108/goFly/goFetch/usecase/
COPY ./usecase/ports/*.go /go/src/github.com/jbl1108/goFly/goFetch/usecase/ports/
COPY ./model/*.go /go/src/github.com/jbl1108/goFly/goFetch/model/
COPY ./util/*.go /go/src/github.com/jbl1108/goFly/goFetch/util/
COPY ./config.conf /go/src/github.com/jbl1108/goFly/goFetch/
COPY ./*.html /go/src/github.com/jbl1108/goFly/goFetch/

WORKDIR /go/src/github.com/jbl1108/goFly/goFetch
RUN ls -s
RUN go mod download
RUN go get github.com/jbl1108/goFly/goFetch/model
RUN go get github.com/jbl1108/goFly/goFetch/usecase
RUN go get github.com/jbl1108/goFly/goFetch/usecase/ports
RUN go get github.com/jbl1108/goFly/goFetch/delivery
RUN go get github.com/jbl1108/goFly/goFetch/gateways
RUN go get github.com/jbl1108/goFly/goFetch/config
RUN go get github.com/jbl1108/goFly/goFetch/util
#RUN dep ensure -vendor-only

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /root/
COPY --from=build /go/src/github.com/jbl1108/goFly/goFetch/app ./
COPY --from=build /go/src/github.com/jbl1108/goFly/goFetch/config.conf ./
COPY --from=build /go/src/github.com/jbl1108/goFly/goFetch/index.html ./
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./app"]
