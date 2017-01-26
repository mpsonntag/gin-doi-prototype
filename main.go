// Copyright (c) 2017, German Neuroinformatics Node (G-Node),
//                     Michael Sonntag <dev@g-node.org>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted under the terms of the BSD License. See
// LICENSE file in the root of the Project.

package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

const (
	major  = 0
	minor  = 1
	status = "Alpha"
)

func versionString() string {
	return fmt.Sprintf("gin-doi %d.%d %s", major, minor, status)
}

const doc = `Prototype for the G-Node Infrastructure DOI service

Usage:
  gin-doi [--conf <dir>]
  gin-doi -h | --help
  gin-doi --version

Options:
  --conf <dir>    Path to the configuration files directory. By default
                  gin-doi will use the resources/conf directory.
  -h --help       Show this screen.
  --version       Print gin-doi version`

func main() {
	args, _ := docopt.Parse(doc, nil, true, versionString(), false)

	if config, ok := args["--conf"]; ok && config != nil {
		fmt.Printf("Use external config file: %v\n", config)
	}

	fmt.Println("Start server")
}
