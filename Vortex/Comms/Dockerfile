FROM bitnami/rabbitmq:latest
LABEL maintainer="damuth.nicholas2@navy.mil"
ADD VERSION .
ENV RABBITMQ_PASSWORD=${RABBITMQ_PASSWORD}
ENV RABBITMQ_USERNAME=${RABBITMQ_USERNAME}
COPY assets/ /tmp/

# RUN apt-get install -y erlang-eldap
COPY assets/definitions.json /tmp/definitions.json

RUN rabbitmq-plugins enable rabbitmq_web_mqtt rabbitmq_mqtt rabbitmq_auth_backend_ldap rabbitmq_management
#ENTRYPOINT ["tail", "-f", "/dev/null"]