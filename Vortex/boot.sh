#!/bin/bash

if [ "$#" -eq  "0" ]
   then
     docker-compose up
 else
     docker-compose up $1
 fi
