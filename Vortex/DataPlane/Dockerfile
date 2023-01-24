FROM golang:latest

COPY assets/GraphPlane /opt/

RUN /bin/sh -c "apt-get update"
RUN /bin/sh -c "apt-get install nano"

ENTRYPOINT ["tail", "-f", "/dev/null"]