FROM golang:1.10.0-stretch
MAINTAINER Neuron Wade neuronwade@gmail.com"

RUN apt-get update && \
        apt-get install -y git

ENV TZ "Asia/Shanghai"

WORKDIR /go/src/alerter
COPY . .
RUN go install ./

EXPOSE 21001

ENTRYPOINT ["/go/bin/alerter"]