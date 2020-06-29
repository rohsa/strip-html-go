// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*// +build ignore*/

package main

import (
	"Practice/Exercise1/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.ViewHandler)
	http.HandleFunc("/submit/", handlers.SubmitHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}