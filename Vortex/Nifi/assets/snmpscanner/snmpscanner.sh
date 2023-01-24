nmap -P0 -v -sU -p 161 -oA snmp_scan 192.168.0.1/24
for i in *.gnmap
do
  for j in $(grep ‘161/open/’ $i | awk ‘{ print $2 }’)
  do
    echo $j
    snmpwalk -v2c -c public $j &> snmpwalk_${j}_public.txt
    if [ "$?" = "0" ]; then echo "$j accepts SNMP community string public"; fi
    snmpwalk -v2c -c private $j &> snmpwalk_${j}_private.txt
    if [ "$?" = "0" ]; then echo "$j accepts SNMP community string private"; fi
  done
done