// Copyright 2024 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Get local version and GitHub version
package version

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"

	"github.com/google/go-github/v59/github"
)

// Fetch latest release in GitHub repo
func fetchLatestRelease(owner, repo string) (*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	ctx := context.Background()
	latest, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)

	return latest, err
}

// GetVersion() display fetch latest release and local version
func GetVersion(csvFile string) {
	file, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	version := "--version"

	for _, item := range records[1:] {
		if item[0] == "" {
			panic(err)
		}

		cmd := exec.Command(item[0], version)

		stdout, err := cmd.Output()
		if err != nil {
			panic(err)
		}

		latest, err := fetchLatestRelease(item[1], item[2])
		if err != nil {
			panic(err)
		}

		fmt.Printf("Local version of %s %slast version in GitHub repo is %s\nurl %s\n\n",
			item[0],
			string(stdout),
			latest.GetTagName(),
			latest.GetHTMLURL())
	}
}
