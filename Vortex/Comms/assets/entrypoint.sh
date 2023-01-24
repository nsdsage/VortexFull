#!/bin/bash

rabbitmq-plugins enable rabbitmq_web_stomp

rabbitmqctl import_definitions /tmp/definitions.json

tail -F /dev/null