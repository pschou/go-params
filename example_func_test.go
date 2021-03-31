// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gnuflag_test

import (
	"errors"
	"fmt"
	"gnuflag"
	"net"
	"os"
)

func ExampleFunc() {
	fs := gnuflag.NewFlagSet("ExampleFunc", gnuflag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	var ip net.IP
	fs.Func("ip", "`IP address` to parse", func(s string) error {
		ip = net.ParseIP(s)
		if ip == nil {
			return errors.New("could not parse IP")
		}
		return nil
	}, "ADDR")
	fs.Parse([]string{"--ip", "127.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())

	// 256 is not a valid IPv4 component
	fs.Parse([]string{"--ip", "256.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())

	// Output:
	// {ip: 127.0.0.1, loopback: true}
	//
	// invalid value "256.0.0.1" for flag --ip: could not parse IP
	// Usage of ExampleFunc:
	// --ip ADDR  `IP address` to parse  (Default: )
	// {ip: <nil>, loopback: false}
}
