package main

import (
	"fmt"
	"os"

	dir "github.com/ministryofjustice/cloud-platform-directory-hash/pkg/hashdir"
	"github.com/sethvargo/go-githubactions"
	create "golang.org/x/mod/sumdb/dirhash"
)

const fileName string = ".checksum"
const base string = "namespaces/live.cloud-platform.service.justice.gov.uk"

func main() {
	// DefaultHash is the default hash function used to hash a directory.
	var DefaultHash create.Hash = create.Hash1

	// Introspects a file found in the root of the cloud-platform-environments.
	prevHash, namespace := dir.ReadChecksum(fileName)

	// Sets up the relative path of a namespace.
	nsDir := base + "/" + namespace

	// Creates a new sha256 hash of a namespace.
	newHash, err := create.HashDir(nsDir, namespace, DefaultHash)
	if err != nil {
		fmt.Println("Failed to create hash, reason:", err)
	}

	// Confirms expected namespace is the only namespace changed in PR.
	_, err = dir.ValidateNamespace(namespace, base)
	if err != nil {
		fmt.Println("Failed to validate namespace, reason:", err)
		// Fail softly as to allow GitHub actions to report success.
		os.Exit(0)
	}

	// Compares two hashes and ensures only a single namespace has been modified.
	// If both conditions have been met a file will be created for a downstream GitHub
	// action. If the conditions fail, a message will be printed. The script deliberately
	// returns exit code 0 regardless of pass or fail to allow GitHub actions to report success.
	if dir.HashesMatch(prevHash, newHash) {
		fmt.Println("Checksums match. Approve PR.")
		githubactions.SetOutput("checksum_match", "true")
	} else {
		fmt.Println("Checksums do not match. Aborting.")
		githubactions.SetOutput("checksum_match", "false")
	}
}
