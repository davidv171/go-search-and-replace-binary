FROM golang:1.8

WORKDIR /go/src/idea
COPY . .
RUN echo "Starting"
RUN touch test
RUN echo "hello world" > test
RUN ls
RUN go build binaryIO.go binarySearch.go linkedList.go osArgsToBinary.go
RUN ./binaryIO test fr 01101100 11111110
RUN cat out
