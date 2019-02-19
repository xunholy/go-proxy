package cntlm

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type CNTLMKeyPairValues struct {
	Key    string
	Values []string
}

func ReadFile(file string) ([]byte, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return []byte{0}, err
	}
	return content, nil
}

func ParseContentIntoStruct(content []byte) {
	var filteredContent []string
	lines := strings.Split(string(content), "\n")
	for _, l := range lines {
		if !strings.HasPrefix(l, "#") {
			filteredContent = append(filteredContent, l)
		}
	}
	fmt.Println(filteredContent)
}

func ParseKeyPairs(line string) {
	keyPairValues := []CNTLMKeyPairValues{}
	fields := strings.Fields(line)
	for _, f := range fields {
		if len(f) > 2 {
			// Handle extra values
		}
		keyPairValues = append(keyPairValues, CNTLMKeyPairValues{Key: string(f[0]), Values: string(f[1])})
	}
}
