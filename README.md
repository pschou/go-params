go-params
-----

# Introduction

What would the module look like if GoLang provided a full-featured flag/parameter parsing package?  ... One which offers flexibility, simplicity while also maintaining the familiar look-and-feel other open-source packages provide.  We must clone the best we can the look-and-feel of other commonly used Linux packages to lower our users' learning curve and blend in with the rest of the technologies available.  Welcome to the solution, go-param.

As there are many examples of programs that handle parameters differently, let us choose two commonly used
packages of which to model and build a generic module to mimic.  The two selected are `ldapsearch` and `curl`.  The first has been around for over two decades, and the second is fairly new.  Both are well used and understood by the Linux community as a whole.  The goal here is to lower the bar of learning and make the flags operate as close as other linux tools operate to ease user's learning curve.  The goal is done!  With this `param` package, GoLang can output the same help and parse the same parameter inputs.

# About

This package is a fork of the Go standard library flag and gnuflag.  As this
package is a rewrite to enable additional functionality and usability, one will find it is significantly different from the source.  The driving motivation was
to provide a solution to the missing toolbox, an excellent flag parser that is simple and is very similar to other gnu programs.  Being very similar to other gnu programs lowers the learning curve for users to use flags in go-built-tools.  Modeled gnu programs used in the creation of this tool are the openldap and curl
help flags.  As this is a personal project (aka no funding), my support time is limited!

# Goals

This re-write includes some notable differences:

- Support for both `--longflag` and `-l` single-character flag syntax.
- Addition of "present" flag with no parameters needed.
- Boolean flags always require a boolean input, true, t, 1, false, f, or 0 with either space ' ' or '=' separator.
- Flag stacking `-abc` is the same as `-a -b -c` for present flags.
- Unicode support for inputs and printing with alignment.
- Multiple flags for a single target value `-i, --include`.
- Custom exemplars demonstrating the needed input type `--time DURATION   How long to wait for a reply.`
- Custom definable functions to handle the parsing of parameters.
- Ability to allow more than one input per parameter `--port-range 1080 1090`, by using the custom var and the needed count.
- Collect a dynamic number of strings per flag into a slice, like args after `--install pkgA pkgB pkgC --remove pkgX`.
- Allow interspersed parameters if set `-a data -b` is the same as `-a -b data`.

# Example

Here is what it looks like when implemented:
```
import (
  ...
  "github.com/pschou/go-params"
  ...
)

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
  // To enable both -H and --host as options, all one needs to do is add a space "host" -> "host H"
  var tls_host = params.String("host", "", "Hostname to verify outgoing connection with", "FQDN")

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
