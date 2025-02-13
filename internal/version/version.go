// Copyright 2025 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Get local version and GitHub version
package version

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

type Release struct {
	TagName string `json:"tag_name"`
	URL     string `json:"url"`
}

// Fetch latest release in GitHub repo
func (r Release) fetchLatestRelease(owner, repo string) ([]string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	release := []string{r.TagName, r.URL}

	return release, nil
}

// Return local version
func localRelease(item string) (string, error) {
	flag := "--version"

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

		r := Release{"", ""}

		latest, err := r.fetchLatestRelease(item[1], item[2])
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
