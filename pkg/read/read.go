package read

import (
	"bufio"
	"fmt"
	"os"
)

func readChecksum(f string) (userHash, namespace string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	namespace = lines[1]
	userHash = lines[2]

	return
}
