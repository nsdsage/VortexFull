#!/bin/bash
influx setup --bucket trades -t mytoken -o demo.net --username=telegraf --password=telegraf123 --host=http://influx:8086 -f
influx bucket create --name snmpdata -t mytoken -o demo.net --host=http://influx:8086