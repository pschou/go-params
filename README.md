go-params
-----

The package is a fork of the Go standard library flag and gnuflag.  As this
package is a rewrite to enable additional functionality and usability, it is
not backwards compatible.

Goals in mind for this re-write and some notable differences this package provides:

- `--longflag` and `-l` single-character flag syntax
- flag stacking `-abc` is the same as `-a -b -c` for present flags
- full unicode support and printing with alignment
- multiple flags for a single value `-i, --include`
- exemplifies the needed input type `--time DURATION`
- custom definable functions to handle parsing of value
- ability to allow more than one input per parameter
- allow interspersed parameters, if set `-a data -b` is the same as `-a -b data`

Full documentation can be found here: https://godoc.org/github.com/pschou/go-param.
