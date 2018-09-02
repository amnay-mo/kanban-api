FROM alpine:latest

LABEL maintainer="amnay.m@gmail.com"

ENTRYPOINT [ "/todos" ]

COPY ./todos /
