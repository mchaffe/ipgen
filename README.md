
`ipgen` is a simple tool to help both offensive and defensive security folk to quickly generate extended format IPv4 addresses. 

While not explicitly allowed by an RFC, these are supported through functions like `inet_aton` and `inet_addr` in the C standard library ([inet(3)](https://man7.org/linux/man-pages/man3/inet.3.html)). These functions are used widely across network programs, accept these extended formats as a matter of convenience and backward compatibility.

While writing `ipgen` I came across an excellent blog [Filters and Bypasses - Rare IPv4 Formats for SSRF](https://dominicbreuker.com/post/filters_bypasses_rare_ipv4_formats_for_ssrf/) which does a great job at explaining the concepts used

# Usage
```
$ ./ipgen 

▪   ▄▄▄· ▄▄ • ▄▄▄ . ▐ ▄ 
██ ▐█ ▄█▐█ ▀ ▪▀▄.▀·•█▌▐█
▐█· ██▀·▄█ ▀█▄▐▀▀▪▄▐█▐▐▌
▐█▌▐█▪·•▐█▄▪▐█▐█▄▄▌██▐█▌
▀▀▀.▀   ·▀▀▀▀  ▀▀▀ ▀▀ █▪

Usage
  ipgen [options] <IPv4 address>

Options:
  -format string
    	specify formats: dec, oct, hex, all (default "all")
  -mix
    	all mixed combinations
  -pad int
    	number of 0s to pad hex and oct numbers
```

Default output
```
$ ipgen 192.0.2.1
192.0.2.1
192.0.513
192.513
3221225985
0300.00.02.01
0300.00.01001
0300.01001
030000001001
0xc0.0x0.0x2.0x1
0xc0.0x0.0x201
0xc0.0x201
0xc0000201
```

padding hexadecimal and octal with leading 0s
```
$ ipgen -format hex,oct -pad 2 192.168.0.1
0x00c0.0x00a8.0x000.0x001
0x00c0.0x00a8.0x001
0x00c0.0x00a80001
0x00c0a80001
000300.000250.0000.0001
000300.000250.0001
000300.00052000001
00030052000001
```

Verifying these actually work
```
$ ipgen -format hex -pad 2 8.8.4.4 | shuf | head -1 | xargs ping -c 1
PING 0x008.0x008.0x004.0x004 (8.8.4.4) 56(84) bytes of data.
64 bytes from 8.8.4.4: icmp_seq=1 ttl=128 time=22.0 ms

--- 0x008.0x008.0x004.0x004 ping statistics ---
1 packets transmitted, 1 received, 0% packet loss, time 0ms
rtt min/avg/max/mdev = 21.975/21.975/21.975/0.000 ms
```

# Server Side Request Forgery (SSRF)

Example usage of `ipgen` is testing SSRF ([MITRE ATT&CK T1190](https://attack.mitre.org/techniques/T1190/)) to access AWS instance metadata service (IMDS)

There may be a simple IPv4 check to prevent an attacker from [retrieving instance metadata](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-data-retrieval.html) via SSRF.

The AWS IMDS URL is http://169.254.169.254/latest/meta-data/

Using `ipgen` to generate a number of different URLs to try:
```bash
$ echo "http://$(ipgen 169.254.169.254 | shuf | head -1)/latest/meta-data"
http://169.16689662/latest/meta-data
```

or

```bash
$ for i in $(ipgen 169.254.169.254); do echo "http://${i}/latest/meta-data"; done
http://169.254.169.254/latest/meta-data
http://169.254.43518/latest/meta-data
http://169.16689662/latest/meta-data
http://2852039166/latest/meta-data
http://0251.0376.0251.0376/latest/meta-data
http://0251.0376.0124776/latest/meta-data
http://0251.077524776/latest/meta-data
http://025177524776/latest/meta-data
http://0xa9.0xfe.0xa9.0xfe/latest/meta-data
http://0xa9.0xfe.0xa9fe/latest/meta-data
http://0xa9.0xfea9fe/latest/meta-data
http://0xa9fea9fe/latest/meta-data
```

Feeling lucky? Everybody shuffling...
```bash
$ echo "http://$(ipgen -mix 169.254.169.254 | shuf | head -1)/latest/meta-data"
http://169.16689662/latest/meta-data
```

To defend against this attack, enable IMDSv2. IMDSv1 is a request/response method while IMDSv2 is session-orientated. More detail on this can be found in the AWS blog [Add defense in depth against open firewalls, reverse proxies, and SSRF vulnerabilities with enhancements to the EC2 Instance Metadata Service](https://aws.amazon.com/blogs/security/defense-in-depth-open-firewalls-reverse-proxies-ssrf-vulnerabilities-ec2-instance-metadata-service/)


# License

This project is licensed under the GPLv3 License - see the LICENSE.md file for details

# Acknowledgements  
- Dominic Breuker

Shout outs: 
- Catalyst (RIP)
- IRATE