FROM golang:latest
LABEL ImageName="golang:hello"

WORKDIR /yurisa
COPY ./main.go /yurisa
RUN go build -o hello /yurisa/main.go
CMD ./hello
