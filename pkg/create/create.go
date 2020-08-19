package create

import (
	"fmt"
	"io/ioutil"
)

func createArtifact(s string) error {
	m := []byte(s)
	err := ioutil.WriteFile(".approved", m, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
