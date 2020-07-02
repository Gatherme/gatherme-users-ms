FROM golang:latest
RUN mkdir -p /go/src/github.com/Gatherme/gatherme-users-ms
ADD . /go/src/github.com/Gatherme/gatherme-users-ms
WORKDIR /go/src/github.com/Gatherme/gatherme-users-ms
RUN go get -v
RUN go install github.com/Gatherme/gatherme-users-ms
ENTRYPOINT /go/bin/gatherme-users-ms
EXPOSE 3000

