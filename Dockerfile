FROM golang:1.10.0-stretch
MAINTAINER Neuron Wade neuronwade@gmail.com"

RUN apt-get update && \
        apt-get install -y git

ENV TZ "Asia/Shanghai"

EXPOSE 21001

WORKDIR /go/src/alerter
COPY . .
RUN go install ./

RUN /go/bin/alerter