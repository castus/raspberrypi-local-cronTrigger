FROM golang:latest

WORKDIR /data

RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install nano -y

COPY ./src /data

ENV LANG en_US.utf8

CMD sh run.sh