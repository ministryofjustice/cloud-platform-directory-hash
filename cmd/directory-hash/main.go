package main

import (
	"fmt"

	dir "golang.org/x/mod/sumdb/dirhash"
)

const fileName string = ".checksum"
const base string = "namespaces/live-1.cloud-platform.service.justice.gov.uk"

func main() {
	// DefaultHash is the default hash function used in new go.sum entries.
	var DefaultHash dir.Hash = dir.Hash1

	original, namespace := read.readChecksum(fileName)

	nsDir := base + "/" + namespace

	hashNs, _ := dir.HashDir(nsDir, namespace, DefaultHash)

	if check.hashesMatch(original, hashNs) && check.singleNamespace(namespace) {
		fmt.Println("Checksums match. Approve PR.")
		err := create.createArtifact("pass")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Checksums do not match. Aborting.")
	}
}
