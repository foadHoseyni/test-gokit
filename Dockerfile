FROM golang:1.17

WORKDIR /go/src/test-gokit

COPY . .
RUN go get .
RUN go build -o main .
CMD ["./main"]
EXPOSE 8000