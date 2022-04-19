# dns-trace
A simple tool to recursively resolves IP for a given domain
It always starts at the root server, traversing until the A record for the domain is found
The tool also resolves if the domain has a CNAME record by applying the same logic to the CNAME target

## Simple example
```go
$ go run trace.go www.google.com.

Resolving domain -> IP query for domain=www.google.com. with nameserver=199.9.14.201:53
NS = a.gtld-servers.net.  NS IP = 192.5.6.30
Resolving domain -> IP query for domain=www.google.com. with nameserver=192.5.6.30:53
NS = ns2.google.com.  NS IP = 216.239.34.10
Resolving domain -> IP query for domain=www.google.com. with nameserver=216.239.34.10:53
Go A record 74.125.24.103

Domain www.google.com. resolved to IP(V4) 74.125.24.103
```
## With CNAME Resolution
The below example has multiple CNAME mappings
```go
$ go run trace.go developer.api.autodesk.com.

Resolving domain -> IP query for domain=developer.api.autodesk.com. with nameserver=199.9.14.201:53
NS = a.gtld-servers.net.  NS IP = 192.5.6.30
Resolving domain -> IP query for domain=developer.api.autodesk.com. with nameserver=192.5.6.30:53
NS = dns1.p05.nsone.net.  NS IP = 157.53.234.1
Resolving domain -> IP query for domain=developer.api.autodesk.com. with nameserver=157.53.234.1:53
Looking for cname adsk-prod.apigee.net.

Resolving domain -> IP query for domain=adsk-prod.apigee.net. with nameserver=199.9.14.201:53
NS = a.gtld-servers.net.  NS IP = 192.5.6.30
Resolving domain -> IP query for domain=adsk-prod.apigee.net. with nameserver=192.5.6.30:53
NS = ns-116.awsdns-14.com.  NS IP = 205.251.194.158
Resolving domain -> IP query for domain=adsk-prod.apigee.net. with nameserver=205.251.194.158:53
Looking for cname adsk-00.dn.apigee.net.

Resolving domain -> IP query for domain=adsk-00.dn.apigee.net. with nameserver=199.9.14.201:53
NS = a.gtld-servers.net.  NS IP = 192.5.6.30
Resolving domain -> IP query for domain=adsk-00.dn.apigee.net. with nameserver=192.5.6.30:53
NS = ns-116.awsdns-14.com.  NS IP = 205.251.194.158
Resolving domain -> IP query for domain=adsk-00.dn.apigee.net. with nameserver=205.251.194.158:53
NS = ns-1346.awsdns-40.org. 
NS IP not found, Run Query for: ns-1346.awsdns-40.org. 
Resolving domain -> IP query for domain=ns-1346.awsdns-40.org. with nameserver=199.9.14.201:53
NS = a0.org.afilias-nst.info.  NS IP = 199.19.54.1
Resolving domain -> IP query for domain=ns-1346.awsdns-40.org. with nameserver=199.19.54.1:53
NS = g-ns-1640.awsdns-40.org.  NS IP = 205.251.192.168
Resolving domain -> IP query for domain=ns-1346.awsdns-40.org. with nameserver=205.251.192.168:53
Go A record 205.251.197.66

Resolving domain -> IP query for domain=adsk-00.dn.apigee.net. with nameserver=205.251.197.66:53
Go A record 54.171.246.133

Domain developer.api.autodesk.com. resolved to IP(V4) 54.171.246.133
```
