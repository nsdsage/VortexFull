FROM apache/nifi:jre-11
LABEL maintainer="nsdamuth@gmail.com"
ADD VERSION .
ARG DEB_SERVER
ARG DEB_PORT
ARG ARTIFACT_SERVER
ARG ARTIFACT_PORT
ARG PIP_REPOS
ARG PIP_PORT

USER root

RUN apt-get -y update && \
        apt-get -y dist-upgrade && \
        apt-get -y upgrade

RUN mkdir -p /etc/atlas/conf

# RUN apt-get install -y snmp && \
#         apt-get install -y snmpd git make gcc ca-certificates dos2unix curl nmap net-tools

COPY assets/* /tmp/
COPY ./configs/authorizers.xml /opt/nifi/nifi-current/conf/authorizers.xml
COPY ./configs/nifi.properties /opt/nifi/nifi-current/conf/nifi.properties

#COPY --chown=nifi:nifi assets/flow.xml.gz /opt/nifi/nifi-current/conf/
COPY assets/*.nar /opt/nifi/nifi-current/lib/
COPY --chown=nifi:nifi assets/xmltojson.xml /opt/nifi/nifi-current/conf/

EXPOSE 8080
