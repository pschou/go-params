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

Examples of parameters flags which can be built using this module:
## Single value
```
-A               for bootstrapping, allow 'any' type  (Default: false)
--Alongflagname  disable bounds checking  (Default: false)
-C               a boolean defaulting to true  (Default: true)
-D               set relative path for local imports  (Default: "")
-E               issue 23543  (Default: "0")
-F STR           issue 23543  (Default: "0")
-I               a non-zero number  (Default: 2.7)
-K               a float that defaults to zero  (Default: 0)
-M               a multiline
                 help
                 string  (Default: "")
-N               a non-zero int  (Default: 27)
-O               a flag
                 multiline help string  (Default: true)
-Z               an int that defaults to zero  (Default: 0)
--maxT           set timeout for dial  (Default: 0s)
-世              a present flag
--世界           unicode string  (Default: "hello")
```

## Multiple value with indentation set
```
-A          for bootstrapping, allow 'any' type  (Default: false)
    --Alongflagname  disable bounds checking  (Default: false)
-C          a boolean defaulting to true  (Default: true)
-D          set relative path for local imports  (Default: "")
-E          issue 23543  (Default: "0")
-F STR      issue 23543  (Default: "0")
-G, --grind STR  issue 23543  (Default: "0")
-I          a non-zero number  (Default: 2.7)
-K          a float that defaults to zero  (Default: 0)
-M          a multiline
            help
            string  (Default: "")
-N          a non-zero int  (Default: 27)
-O          a flag
            multiline help string  (Default: true)
-Z          an int that defaults to zero  (Default: 0)
    --maxT  set timeout for dial  (Default: 0s)
-世         a present flag
    --世界  unicode string  (Default: "hello")
```

Full documentation can be found here: https://godoc.org/github.com/pschou/go-param.
