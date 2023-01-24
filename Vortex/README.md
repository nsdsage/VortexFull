
# README

## UUIDs

### UUID & System Serial Number

'''
sudo dmidecode -s system-uuid
sudo dmidecode -s system-serial-number
'''

## Baseboard SSN

'''
sudo dmidecode -s baseboard-serial-number
'''

### Chassis SSN

'''
sudo dmidecode -s chassis-serial-number
'''

### Processor ID

'''
sudo dmidecode -t PROCESSOR | grep "ID:" | sed "s/ID\://g" | sed "s/\s//g"
'''

### BlE Serial Number

'''
hcitool dev | grep -w hci0 | sed "s/^.hci0//g" | sed "s/\s//g" | sed "s/\://g
'''

### NETWORK SSN (WIFI & ETHERNET)

'''
sudo lshw -class network | grep serial: | sed "s/serial\://g" | sed "s/\s//g" | sed "s/\://g"
'''

## SNMP

### SNMP Test Setup

'''
sudo apt-get update
sudo apt-get install -y snmp 
sudo apt-get install -y snmp-mibs-downloader
sudo apt-get install -y snmpd
'''

### SNMP Brute Scan

'''
nmap -sU -p161 --script snmp-brute --script-args snmplist=community.lst 192.168.1.0/24
'''

### LOCALHOST TEST

'''
snmpwalk -v 2c -c public -O e 127.0.0.1
'''
