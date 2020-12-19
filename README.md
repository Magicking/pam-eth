# pam-eth

Do not use this in production.

A pam security module backed by access control smart contract.

![](doc/pamette.png)

Forked from Uber's SSH certificate pam module.

This is a pam module that will authenticate a user based on them providing the valid challenge. 

Building:

1. clone the repo and run 'make'
```
  $ git clone github.com/Magicking/pam-eth

  ...

  $ make
go test -cover
PASS
coverage: 68.3% of statements
ok  	github.com/Magicking/pam-eth	0.111s
go build -buildmode=c-shared -o pam_eth.so
chmod +x pam_eth.so

  $
```

Usage:

1. put this pam module where ever pam modules live on your system, eg. `/lib/security`

2. add it as an authentication method, eg.

```
  $ grep auth /etc/pam.d/sudo
  auth [success=1 default=ignore] /lib/security/pam_eth.so
  auth requisite                  pam_deny.so
  auth required                   pam_permit.so
```

3. make sure you are ready to provide the challenge payload

## Testing

Modify your /etc/pam.d/chfn
```
auth⤇   sufficient⤇     /home/magicking/source/gocode/src/github.com/Magicking/pam-eth/pam_eth.so contract-address=0x000 rpc-endpoint=wss://goerli.eth.6120.eu/ws
account⤇sufficient⤇     /home/magicking/source/gocode/src/github.com/Magicking/pam-eth/pam_eth.so contract-address=0x000 rpc-endpoint=wss://goerli.eth.6120.eu/ws
session⤇sufficient⤇     /home/magicking/source/gocode/src/github.com/Magicking/pam-eth/pam_eth.so contract-address=0x000 rpc-endpoint=wss://goerli.eth.6120.eu/ws
```

And try using the command `chfn`

Watch build result with `while true; do inotifywait -e modify *.go ; sleep 0.3 ; make; done`

## Example configuration

TODO(1)
```
auth [success=1 default=ignore] /lib/security/pam_eth.so contract-address=0x6120... rpc-endpoint=wss://goerli.eth.6120.eu/wss
```

## Information
 - See Uber's https://medium.com/uber-security-privacy/introducing-the-uber-ssh-certificate-authority-4f840839c5cc
