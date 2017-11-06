# Enteprise Connect SDK Usage

Below is the list of all the available parameters for EC-SDK.

```shellscript
./ecagent -h
Usage of ./ecagent:
 -aid string
    	Specify the agent Id assigned by the EC Service. You may find it in the Cloud Foundry VCAP_SERVICE
  -bkl string
    	Specify the ip(s) blocklist in the IPv4/IPv6 format. Concatenate ips by comma. E.g. 10.20.30.5, 2002:4559:1FE2::4559:1FE2
  -cid string
    	Specify the client Id to auto-refresh the OAuth2 token.
  -crt string
    	Specify the relative path of a digital certificate to operate the EC agent. (.pfx, .cer, .p7s, .der, .pem, .crt)
  -csc string
    	Specify the client secret to auto-refresh the OAuth2 token.
  -dbg
    	Turn on debug mode. This will introduce more error information. E.g. connection error.
  -dur int
    	Specify the duration for the next token refresh in seconds. (default 100 years)
  -gen
    	Generate a certificate request for the usage validation purpose.
  -hca string
    	Specify a port# to turn on the Healthcheck API. This flag is always on when in the "gateway mode" with the provisioned local port. Upon provisioned, the api is available at <agent_uri>/health.
  -hst string
    	Specify the EC Gateway URI. E.g. wss://<somedomain>:8989
  -inf
    	The Product Information.
  -lpt string
    	Specify the EC port# if the "client" mode is set. (default "7990")
  -mod string
    	Specify the EC Agent Mode in "client", "server", or "gateway". (default "agent")
  -oa2 string
    	Specify URL of the OAuth2 provisioner. E.g. https://<somedomain>/oauth/token
  -pct string
    	Specify the relative path to a TLS cert when operate as the gateway as desired. E.g. ./path/to/cert.pem.
  -pky string
    	Specify the relative path to a TLS key when operate as the gateway as desired. E.g. ./path/to/key.pem.
  -plg string
    	Enable plugin list. Available options "tls", "ip-route", etc. 
  -pxy string
    	Specify a local Proxy service. E.g. http://hello.world.com:8080
  -rht string
    	Specify the Resource Host if the "server" mode is set. E.g. <someip>, <somedomain>. value will be discard when TLS is specified.
  -rpt string
    	Specify the Resource Port# if the "server" mode is set. E.g. 8989, 38989
  -sgn
    	Start a CA Cert-Signing process.
  -sst string
    	Specify the EC Service URI. E.g. https://<service.of.predix.io>
  -tid string
    	Specify the Target EC Server Id if the "client" mode is set
  -tkn string
    	Specify the OAuth Token. The token may expire depending on your OAuth provisioner. This flag is ignored if OAuth2 Auto-Refresh were set.
  -ver
    	Show EC Agent's version.
  -vfy
    	Verify the legitimacy of a digital certificate.
  -wtl string
    	Specify the ip(s) whitelist in the cidr net format. Concatenate ips by comma. E.g. 89.24.9.0/24, 7.6.0.0/16 (default "0.0.0.0/0,::/0")
  -zon string
    	Specify the Zone/Service Inst. Id. required in the "gateway" mode.
```