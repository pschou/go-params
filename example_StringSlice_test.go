// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params_test

import (
	"fmt"
	"github.com/pschou/go-params"
	"os"
)

func ExampleStringSlice() {
	fs := params.NewFlagSet("ExampleFunc", params.ContinueOnError)
	fs.SetOutput(os.Stdout)
	var install, remove []string

	fs.StringSliceVar(&install, "i install", "List of packages to install", "PACKAGES")
	fs.StringSliceVar(&remove, "r remove", "List of packages to install", "PACKAGES")
	fs.Parse([]string{"--install", "a", "b", "-r", "c"})

	fmt.Printf("{install: %#v, remove: %#v}\n\n", install, remove)

	// Output:
	// {install: []string{"a", "b"}, remove: []string{"c"}}
}
