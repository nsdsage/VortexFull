FROM telegraf:latest

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y systemd

COPY assets/telegraf.d /etc/telegraf/telegraf.d/