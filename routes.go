// Copyright (c) 2017, German Neuroinformatics Node (G-Node)
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted under the terms of the BSD License. See
// LICENSE file in the root of the Project.

package main

import "github.com/gorilla/mux"

// RegisterRoutes does exactly that.
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", handler)
	r.HandleFunc("/authenticate", authenticate)
}
