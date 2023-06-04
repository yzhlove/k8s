FROM alpine:latest
LABEL Auth="yurisa"

RUN apk add --no-cache tree;
WORKDIR /yurisa
ADD ./1.txt /yurisa/a/
ADD ./2.txt /yurisa/b/

CMD pwd;tree;
