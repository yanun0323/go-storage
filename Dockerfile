# build stage
FROM golang:1.22-alpine AS build

ADD . /go/build
WORKDIR /go/build

ADD config.yaml /go/build/
ADD go.mod go.sum /go/build/

RUN go mod download

# install gcc
RUN apk add build-base

RUN go build -o go-storage main.go

# final stage
FROM alpine:3.18

# install timezone data
RUN apk add --no-cache tzdata
ENV TZ Asia/Taipei
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY --from=build /go/build/go-storage /var/application/go-storage
COPY --from=build /go/build/config.yaml /var/application/config.yaml

EXPOSE 8001

WORKDIR /var/application
CMD [ "./go-storage" ]