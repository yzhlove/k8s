FROM golang:latest AS golang
LABEL IMAGE="GOLANG:Hello"
WORKDIR /yurisa
COPY ./main.go /yurisa
RUN go build -o hello ./main.go
CMD "./hello"

#FROM alpine:latest
FROM scratch
LABEL IMAGE="ALPINE:Hello"
#WORKDIR /yurisa
# --from=0 指的是第一个FROM
# --from=golang 构建过程使用别名，更好区分
#COPY --from=golang /yurisa/hello /yurisa/
COPY --from=golang /yurisa/hello /
#RUN mv /yurisa/hello /yurisa/alpine
#CMD "./alpine"
CMD "./hello"
