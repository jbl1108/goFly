# syntax=docker/dockerfile:1
FROM golang:1.18-alpine AS build
# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
RUN apk add --no-cache git
RUN mkdir -p /go/src/github.com/jbl1108/gofly/
COPY go.mod /go/src/github.com/jbl1108/gofly/
COPY *.go /go/src/github.com/jbl1108/gofly/
COPY ./config/*.go /go/src/github.com/jbl1108/gofly/config/
COPY ./delivery/*.go /go/src/github.com/jbl1108/gofly/delivery/
COPY ./gateways/*.go /go/src/github.com/jbl1108/gofly/gateways/
COPY ./repositories/*.go /go/src/github.com/jbl1108/gofly/repositories/
COPY ./usecase/*.go /go/src/github.com/jbl1108/gofly/usecase/
COPY ./usecase/ports/*.go /go/src/github.com/jbl1108/gofly/usecase/ports/
COPY ./model/*.go /go/src/github.com/jbl1108/gofly/model/
COPY ./util/*.go /go/src/github.com/jbl1108/gofly/util/
COPY ./config.conf /go/src/github.com/jbl1108/gofly/
COPY ./*.html /go/src/github.com/jbl1108/gofly/

WORKDIR /go/src/github.com/jbl1108/gofly
RUN ls -s
RUN go mod download
RUN go get github.com/jbl1108/goFly/model
RUN go get github.com/jbl1108/goFly/usecase
RUN go get github.com/jbl1108/goFly/usecase/ports
RUN go get github.com/jbl1108/goFly/delivery
RUN go get github.com/jbl1108/goFly/gateways
RUN go get github.com/jbl1108/goFly/config
RUN go get github.com/jbl1108/goFly/util
#RUN dep ensure -vendor-only

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /root/
COPY --from=build /go/src/github.com/jbl1108/gofly/app ./
COPY --from=build /go/src/github.com/jbl1108/gofly/config.conf ./
COPY --from=build /go/src/github.com/jbl1108/gofly/index.html ./
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./app"]
