FROM golang:1.8

WORKDIR /go/src/idea
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN touch test
RUN echo "hello world" > test
CMD go build *.go && ./binaryIO test fr 01101100 11111110 && cat out
