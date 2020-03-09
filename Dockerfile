FROM golang:latest as build

COPY . /go/src/trivago-BackendSoftwareEngineer-UserFlow
WORKDIR /go/src/trivago-BackendSoftwareEngineer-UserFlow
#COPY . .
RUN go get gopkg.in/yaml.v2
RUN go build -o main .
##--mount source=/go/src/trivago-BackendSoftwareEngineer-UserFlow/output,destination=./output
CMD ["./main", "json", "xml", "yaml"]

 #syntax = docker/dockerfile:experimental
#FROM golang:latest


#COPY  "/go/src/trivago-BackendSoftwareEngineer-UserFlow/output" .
