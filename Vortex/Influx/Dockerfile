FROM quay.io/influxdb/influxdb:v2.0.8

COPY assets/ /tmp/
RUN chmod 770 /tmp/entrypoint.sh

ENTRYPOINT ["/tmp/entrypoint.sh"]