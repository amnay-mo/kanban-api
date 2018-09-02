FROM alpine:latest

LABEL maintainer="amnay.m@gmail.com"

ENTRYPOINT [ "/kanban-api" ]

COPY ./kanban-api /
