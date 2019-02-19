package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func UpdateFile(match string, value string) {
	input, err := ioutil.ReadFile(cntlmFile)
	if err != nil {
		log.Fatalln(err)
	}
	v := strings.TrimSpace(value)

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, match) {
			l := fmt.Sprintf("%v\t%v", match, v)
			lines[i] = l
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(cntlmFile, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func UpdatePassword(hashes string) {
	input, err := ioutil.ReadFile(cntlmFile)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	hashlines := strings.Split(hashes, "\n")

	for x, hash := range hashlines {
		v := strings.Fields(hash)
		if len(v) > 0 && x >= 1 {
			for i, line := range lines {
				if strings.Contains(line, v[0]) {
					lines[i] = hash
				}
			}
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(cntlmFile, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

}
