#! /bin/bash

# This script assumes that nifi-jre11 and Vortex are in the same directory 
#
# Builds nifi/jre-11 image  
# docker-compose Vortex

cd ../nifi-jre11/nifi-docker/dockerhub
docker build -t apache/nifi:jre-11 .

cd ../../../Vortex
docker-compose up
