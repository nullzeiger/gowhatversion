// Copyright 2024 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Get local version and GitHub version
package version

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

type Release struct {
	TagName string `json:"tag_name"`
	Url     string `json:"url"`
}

// Fetch latest release in GitHub repo
func fetchLatestRelease(owner, repo string) ([]string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.Status)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err

	}

	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	a := []string{release.TagName, release.Url}

	return a, nil
}

// Return local version
func localRelease(item string) (string, error) {
	flag := "--version"

	if item == "zig" {
		flag = "version"
	}

	out, err := exec.Command(item, flag).Output()
	if err != nil {
		return "", fmt.Errorf("%w %s does not have a --version flag for printing the version", err, item)
	}

	return string(out), nil
}

// GetVersion() display fetch latest release and local version
func GetVersion(csvFile string) error {
	file, err := os.Open(csvFile)
	if err != nil {
		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, item := range records {
		if item[0] == "" {
			fmt.Println("Not element in file")
		}

		local, err := localRelease(item[0])
		if err != nil {
			return err
		}

		latest, err := fetchLatestRelease(item[1], item[2])
		if err != nil {
			return err
		}

		fmt.Printf("Local version of %s %slast version in GitHub repo is %s\nurl %s\n\n",
			item[0],
			local,
			latest[0],
			latest[1])
	}

	return nil
}
