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
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DOI Server running")
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate request")
}
