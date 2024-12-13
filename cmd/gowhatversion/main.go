// Copyright 2024 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/nullzeiger/gowhatversion/internal/utility"
	"github.com/nullzeiger/gowhatversion/internal/version"
)

const filename = "/.apps.csv"

func main() {
	home, err := utility.GetHome()
	if err != nil {
		log.Fatal(err)
	}

	err = version.GetVersion(home + filename)
	if err != nil {
		log.Fatal(err)
	}
}
