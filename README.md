go-params
-----

The package is a fork of the Go standard library flag and gnuflag.  As this
package is a rewrite to enable additional functionality and usability.  The driving motivation was
to provide a solution to the missing toolbox for a good flag parser that is both simple and doesn't
differ from other gnu programs.  Some models used in the creation of this tool is the openldap and curl
help flags.  This is a personal project (aka: no funding), and thus my support time is limited!

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

# Example

Here is what it looks like when implemented:
```
var version = "0.0"
func main() {
  // Set a custom header,
  params.Usage = func() {
    fmt.Fprintf(os.Stderr, "My Sample, Version: %s\n\n" +
      "Usage: %s [options...]\n\n", version, os.Args[0])
    params.PrintDefaults()
  }

  // An example boolean flag, used like this: -tls true -tls false, or optionally: -tls=true -tls=false
  var tls_enabled = params.Bool("tls", true, "Enable listener TLS", "BOOL")

  // An example of a present flag, returns true if it was seen
  var verbose = params.Pres("debug", "Verbose output")

  // Start of a grouping set
  params.GroupingSet("Listener")
  var listen = params.String("listen", ":7443", "Listen address for forwarder", "HOST:PORT")
  var verify_server = params.Bool("verify-server", true, "Verify server, do certificate checks", "BOOL")
  var secure_server = params.Bool("secure-server", true, "Enforce minimum of TLS 1.2 on server side", "BOOL")

  // Start of another grouping set
  params.GroupingSet("Target")
  var target = params.String("target", "127.0.0.1:443", "Sending address for forwarder", "HOST:PORT")
  var verify_client = params.Bool("verify-client", true, "Verify client, do certificate checks", "BOOL")
  var secure_client = params.Bool("secure-client", true, "Enforce minimum of TLS 1.2 on client side", "BOOL")
  // using -H and --host as options, all one needs to do is add a space
  var tls_host = params.String("H host", "", "Hostname to verify outgoing connection with", "FQDN")

  // Start of our last grouping set
  params.GroupingSet("Certificate")
  var cert_file = params.String("cert", "/etc/pki/server.pem", "File to load with CERT - automatically reloaded every minute\n", "FILE")
  var key_file = params.String("key", "/etc/pki/server.pem", "File to load with KEY - automatically reloaded every minute\n", "FILE")
  var root_file = params.String("ca", "/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem", "File to load with ROOT CAs - reloaded every minute by adding any new entries\n", "FILE")

  // Indicate that we want all the flags indented for ease of reading
  params.CommandLine.Indent = 2

  // Let us parse everything!
  params.Parse()

  // ... Variables are ready for use now!
}
```
This example was taken directly from the SSL-Forwarder program (below) so one may compare the output and see what it looks like in the finished product.

# Real World Examples
Here are some examples which demonstrate the power of this paramber parsing tool.

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
Prometheus Collector, written by Paul Schou (github.com/pschou/prom-collector) in December 2020
Prsonal use only, provided AS-IS -- not responsible for loss.
Usage implies agreement.

Usage: ./prom-collector [options...]

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





Full documentation can be found here: https://godoc.org/github.com/pschou/go-param.
