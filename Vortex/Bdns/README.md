# DNS Setup

## Test operation

'''
    nslookup -port=1234 -debug mail.vortex.mil 127.0.0.1
'''

## Enabling DNSSEC

'''
    systemd-resolve –status
    look for “DNSSEC setting: no”
    sudo mkdir -p /etc/systemd/resolved.conf.d
    sudo nano /etc/systemd/resolved.conf.d/dnssec.conf
    enter:
    [Resolve]
    DNSSEC=true
    sudo systemctl restart systemd-resolved
    systemd-resolve –status
    look for “DNSSEC setting: yes”
'''

## TCP Format Request

### JWT Token

'''
Test key : $C&F)J@NcRfUjXn2r5u8x/A%D*G-KaPd

'''
'''
{
    "Credentials": {
        "User": 0,
        "Token": ""
    },
    "Request": {
        "Type": "lookup",
        "Version": "v4",
        "Address": "vortex.mil",
        "Sec": 2
    }
}
'''
