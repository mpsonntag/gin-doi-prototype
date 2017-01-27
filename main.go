// Copyright (c) 2017, German Neuroinformatics Node (G-Node)
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted under the terms of the BSD License. See
// LICENSE file in the root of the Project.

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/gorilla/mux"
)

const (
	major  = 0
	minor  = 1
	status = "Alpha"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server running")
}

func versionString() string {
	return fmt.Sprintf("gin-doi-prototpye %d.%d %s", major, minor, status)
}

var port = ":8083"

const doc = `Prototype for the G-Node Infrastructure DOI service

Usage:
  gin-doi-prototype [--conf <dir>] [--listen <address>]
  gin-doi-prototype -h | --help
  gin-doi-prototype --version

Options:
  -h --help           Show this screen.
  --version           Print version.
  --conf <dir>        Path to the configuration files directory. default: ./resources/conf
  --listen <address>  Address to listen at [default: :8083]
  `

func main() {
	args, err := docopt.Parse(doc, nil, true, versionString(), false)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] Parsing command line: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("[Starting server] arguments: %v\n", args)

	if config, ok := args["--conf"]; ok && config != nil {
		fmt.Printf("[Starting server] Use external config file: %v\n", config)
	}

	if p, ok := args["--listen"]; ok {
		fmt.Printf("[Starting server] Use port: %s\n", p)
		port = args["--listen"].(string)
	}

	fmt.Println("[Starting server] Registering routes")
	router := mux.NewRouter()
	router.HandleFunc("/", handler)

	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Println("[Starting server] Listen and serve")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] Server startup: %v\n", err)
		os.Exit(-1)
	}

}
