FROM golang:latest

RUN apt-get update && \
    apt-get dist-upgrade -y && \
    apt-get upgrade -y && \
    apt autoremove -y

RUN echo "resolvconf resolvconf/linkify-resolvconf boolean false" | debconf-set-selections

 RUN apt-get install -y nano dialog apt-utils dnsutils resolvconf