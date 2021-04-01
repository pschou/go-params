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
- allow interspersed parameters, if set `-a data -b` is the same as `-a -b data`

# Background

As there are many example of a programs that handle parameters differently, let us choose one commonly used
package to build a toolbox from, `curl`.  This has been done; in this `param` package, GoLang can duplicate
the output of `curl`.

Example parameters which can be built using this module:
```
$ curl --help
Usage: curl [options...] <url>
Options: (H) means HTTP/HTTPS only, (F) means FTP only
     --anyauth       Pick "any" authentication method (H)
 -a, --append        Append to target file when uploading (F/SFTP)
     --basic         Use HTTP Basic Authentication (H)
     --cacert FILE   CA certificate to verify peer against (SSL)
     --capath DIR    CA directory to verify peer against (SSL)
 -E, --cert CERT[:PASSWD] Client certificate file and password (SSL)
     --cert-type TYPE Certificate file type (DER/PEM/ENG) (SSL)
     --ciphers LIST  SSL ciphers to use (SSL)
     --compressed    Request compressed response (using deflate or gzip)
 -K, --config FILE   Specify which config file to read
     --connect-timeout SECONDS  Maximum time allowed for connection
 -C, --continue-at OFFSET  Resumed transfer offset
 -b, --cookie STRING/FILE  String or file to read cookies from (H)
 -c, --cookie-jar FILE  Write cookies to this file after operation (H)
     --create-dirs   Create necessary local directory hierarchy
```

Full documentation can be found here: https://godoc.org/github.com/pschou/go-param.
