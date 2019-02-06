package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func UpdatePort(p int) {
	input, err := ioutil.ReadFile(cntlmFile)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "Listen") {
			l := fmt.Sprintf("Listen %v", p)
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
