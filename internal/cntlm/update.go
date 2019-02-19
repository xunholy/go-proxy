package cntlm

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type CNTLMKeyPairValues struct {
	Key   string
	Value string
	Line  int
}

func UpdateFile(file, match string) {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}
	match = strings.TrimSpace(match)

	matches := strings.Split(match, "\n")
	lines := strings.Split(string(content), "\n")

	keyPairValues := ParseFileIntoKeyPairValues(lines)

	for i := 0; i <= len(matches)-1; i++ {
		matchFields := strings.Fields(matches[i])
		for _, i := range keyPairValues {
			if strings.Contains(i.Key, matchFields[0]) {
				UpdateValue(lines, i, file, matchFields)
			}
		}
	}
}

func ParseFileIntoKeyPairValues(lines []string) []CNTLMKeyPairValues {
	keyPairValues := []CNTLMKeyPairValues{}
	for i, l := range lines {
		if strings.HasPrefix(l, "#") {
			continue
		}
		if l == "" {
			continue
		}
		fields := strings.Fields(l)
		if len(fields) > 2 {
			continue
		}
		keyPairValues = append(keyPairValues, CNTLMKeyPairValues{Key: fields[0], Value: fields[1], Line: i})
	}
	return keyPairValues
}

func UpdateValue(lines []string, keyPairValue CNTLMKeyPairValues, file string, matchFields []string) {
	line := fmt.Sprintf("%v\t%v", matchFields[0], matchFields[1])
	lines[keyPairValue.Line] = line
	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
