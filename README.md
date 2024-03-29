# rtsp-proxy
Imported from [aler9/rtsp-simple-proxy](https://github.com/aler9/rtsp-simple-proxy)
[![Go Report Card](https://goreportcard.com/badge/github.com/aler9/rtsp-simple-proxy)](https://goreportcard.com/report/github.com/aler9/rtsp-simple-proxy)
[![Build Status](https://travis-ci.org/aler9/rtsp-simple-proxy.svg?branch=master)](https://travis-ci.org/aler9/rtsp-simple-proxy)

_rtsp-simple-proxy_ is a simple, ready-to-use and zero-dependency RTSP proxy, a software that receives one or more existing RTSP streams and makes them available to other users. A proxy is usually deployed in one of these scenarios:
* when there are multiple users that are receiving a stream and the bandwidth is limited, so the proxy is used to receive the stream once. Users can then connect to the proxy instead of the original source.
* when there's a NAT / firewall between a stream and the users, in this case the proxy is installed in the NAT and makes the stream available to the outside world.

Features:
* Receive multiple streams in TCP or UDP
* Distribute streams in TCP or UDP
* Supports the RTP/RTCP streaming protocol
* Supports authentication (i.e. username and password)
* Compatible with Linux, Windows and Mac, does not require any dependency or interpreter, it's a single executable

## Installation

Precompiled binaries are available in the [release](https://github.com/aler9/rtsp-simple-proxy/releases) page. Just download and extract the executable.

## Usage

#### Basic usage

1. Create a configuration file named `conf.yml`, placed in the same folder of the executable, with the following content:
   ```yaml
   streams:
     # name of the stream
     mypath:
       # url of the source stream, in the format rtsp://user:pass@host:port/path
       url: rtsp://myhost:8554/mystream
   ```

2. Launch the proxy:
   ```
   ./rtsp-simple-proxy
   ```

3. Open any stream you have defined in the configuration file, by using the stream name as path, for instance with VLC:
   ```
   vlc rtsp://localhost:8554/mypath
   ```

#### Full configuration file

```yaml
# timeout of read operations
readTimeout: 5s
# timeout of write operations
writeTimeout: 5s

server:
  # supported protocols
  protocols: [ tcp, udp ]
  # port of the RTSP TCP listener
  rtspPort: 8554
  # port of the RTP UDP listener
  rtpPort: 8050
  # port of the RTCP UDP listener
  rtcpPort: 8051
  # optional username required to read
  readUser:
  # optional password required to read
  readPass:

streams:
  # name of the stream
  test1:
    # url of the source stream, in the format rtsp://user:pass@host:port/path
    url: rtsp://myhost:8554/mystream
    # whether to use tcp or udp
    protocol: udp
```

#### Full command-line usage

```
usage: rtsp-simple-proxy [<flags>] <confpath>

rtsp-simple-proxy v0.0.0

RTSP proxy.

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --version  print version

Args:
  [<confpath>]  path of a config file. The default is conf.yml. Use 'stdin' to
                read config from stdin
```

## Links

Related projects
* https://github.com/aler9/rtsp-simple-server
* https://github.com/aler9/gortsplib

IETF Standards
* RTSP 1.0 https://tools.ietf.org/html/rfc2326
* RTSP 2.0 https://tools.ietf.org/html/rfc7826
* HTTP 1.1 https://tools.ietf.org/html/rfc2616
