go-params
-----

The package is a fork of the Go standard library flag and gnuflag.  As this
package is a rewrite to enable additional functionality and usability, it is
not backwards compatible.

# Goals

This re-write includes some notable differences:

- `--longflag` and `-l` single-character flag syntax
- flag stacking `-abc` is the same as `-a -b -c` for present flags
- full unicode support and printing with alignment
- multiple flags for a single value `-i, --include`
- exemplifies the needed input type `--time DURATION`
- custom definable functions to handle parsing of value
- ability to allow more than one input per parameter
- collect a dynamic number of strings per flag into a slice, like args after `--install`
- allow interspersed parameters, if set `-a data -b` is the same as `-a -b data`

# Background

As there are many example of a programs that handle parameters differently, let us choose one commonly used
package to build a toolbox from, `curl`.  This has been done; in this `param` package, GoLang can duplicate
the output of `curl`.

# Real World Examples
Here are some examples which demonstrate the power of this paramber parsing tool.

## jqURL -- https://github.com/pschou/jqURL
```
$ jqurl -h
jqURL - URL and JSON parser tool, Written by Paul Schou (github.com/pschou/jqURL)
Usage:
  ./jqurl [options] "JSON Parser" URLs

Options:
  -C, --cache          Use local cache to speed up static queries
      --cachedir DIR   Path for cache  (Default="/dev/shm")
      --debug          Debug / verbose output
      --flush          Force redownload, when using cache
  -i, --include        Include header in output
      --max-age DURATION  Max age for cache  (Default=4h0m0s)
  -o, --output FILE    Write output to <file> instead of stdout  (Default="")
  -P, --pretty         Pretty print JSON with indents
  -r, --raw-output     Raw output, no quotes for strings
Request options:
  -d, --data STRING    Data to use in POST (use @filename to read from file)  (Default="")
  -H, --header 'HEADER: VALUE'  Custom header to pass to server
                         (Default="content-type: application/json")
  -k, --insecure       Ignore certificate validation checks
  -L, --location       Follow redirects
  -m, --max-time DURATION  Timeout per request  (Default=15s)
      --max-tries TRIES  Maximum number of tries  (Default=30)
  -X, --request METHOD  Method to use for HTTP request (ie: POST/GET)  (Default="GET")
      --retry-delay DURATION  Delay between retries  (Default=7s)
Certificate options:
      --cacert FILE    Use certificate authorities, PEM encoded  (Default="")
  -E, --cert FILE      Use client cert in request, PEM encoded  (Default="")
      --key FILE       Key file for client cert, PEM encoded  (Default="")
```

## SSL-Forwarder -- https://github.com/pschou/ssl-forwarder
```
$ ./ssl-forwarder -h
Simple SSL forwarder, written by Paul Schou (github.com/pschou/ssl-forwarder) in December 2020
All rights reserved, personal use only, provided AS-IS -- not responsible for loss.
Usage implies agreement.

Usage: ./ssl-forwarder [options...]

Options:
  --debug                 Verbose output
  --tls BOOL              Enable listener TLS  (Default: true)
Listener options:
  --listen HOST:PORT      Listen address for forwarder  (Default: ":7443")
  --secure-server BOOL    Enforce minimum of TLS 1.2 on server side  (Default: true)
  --verify-server BOOL    Verify server, do certificate checks  (Default: true)
Target options:
  --host FQDN             Hostname to verify outgoing connection with  (Default: "")
  --secure-client BOOL    Enforce minimum of TLS 1.2 on client side  (Default: true)
  --target HOST:PORT      Sending address for forwarder  (Default: "127.0.0.1:443")
  --verify-client BOOL    Verify client, do certificate checks  (Default: true)
Certificate options:
  --ca FILE               File to load with ROOT CAs - reloaded every minute by adding any new entries
                            (Default: "/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem")
  --cert FILE             File to load with CERT - automatically reloaded every minute
                            (Default: "/etc/pki/server.pem")
  --key FILE              File to load with KEY - automatically reloaded every minute
                            (Default: "/etc/pki/server.pem")
```

## Prom-collector -- https://github.com/pschou/prom-collector
```
$ ./prom-collector -h
Prometheus Collector - written by Paul Schou (github.com/pschou/prom-collector) in December 2020
Prsonal use only, provided AS-IS -- not responsible for loss.
Usage implies agreement.

 Usage of ./prom-collector:
Options:
--ca FILE             File to load with ROOT CAs - reloaded every minute by adding any new entries
                        (Default: "/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem")
--cert FILE           File to load with CERT - automatically reloaded every minute
                        (Default: "/etc/pki/server.pem")
--debug               Verbose output
--json JSON_FILE      Path into which to put all the prometheus endpoints for polling
                        (Default: "/dev/shm/metrics.json")
--key FILE            File to load with KEY - automatically reloaded every minute
                        (Default: "/etc/pki/server.pem")
--listen HOST:PORT    Listen address for metrics  (Default: ":9550")
--path DIRECTORY      Path into which to put the prometheus data  (Default: "/dev/shm/collector")
--prefix URL_PREFIX   Used for all incoming requests, useful for a reverse proxy endpoint
                        (Default: "/collector")
--secure-server BOOL  Enforce TLS 1.2+ on server side  (Default: true)
--tls BOOL            Enable listener TLS  (Default: false)
--verify-server BOOL  Verify or disable server certificate check  (Default: true)
```


Full documentation can be found here: https://godoc.org/github.com/pschou/go-param.
