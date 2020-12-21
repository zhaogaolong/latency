# latency

latency is sets the network delay http server, Implementation principle based on [tc](https://tldp.org/HOWTO/Traffic-Control-HOWTO/intro.html).

---

## install

```shell
$ git clone git@github.com:zhaogaolong/latency.git
$ cd latency
$ kubectl apply -f deploy/pod.yml
```

## how to use

```shell
# curl -X GET podIP:8080/latency/20ms
$ curl -X GET 192.168.234.124:8080/latency/20ms
set latency 20ms success

# ping pod ip
$ ping -w 3 192.168.234.124
PING 192.168.234.124 (192.168.234.124) 56(84) bytes of data.
64 bytes from 192.168.234.124: icmp_seq=1 ttl=64 time=20.0 ms
64 bytes from 192.168.234.124: icmp_seq=2 ttl=64 time=20.1 ms
64 bytes from 192.168.234.124: icmp_seq=3 ttl=64 time=20.5 ms

--- 192.168.234.124 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2002ms
rtt min/avg/max/mdev = 20.076/20.248/20.517/0.224 ms

# reset latency time
$ curl -X GET 192.168.234.124:8080/latency/50ms
set latency 50ms success

# ping pod ip
$ ping -w 3 192.168.234.124
PING 192.168.234.124 (192.168.234.124) 56(84) bytes of data.
64 bytes from 192.168.234.124: icmp_seq=1 ttl=64 time=50.0 ms
64 bytes from 192.168.234.124: icmp_seq=2 ttl=64 time=50.8 ms
64 bytes from 192.168.234.124: icmp_seq=3 ttl=64 time=50.9 ms

--- 192.168.234.124 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2005ms
rtt min/avg/max/mdev = 50.093/50.603/50.905/0.406 ms
```
