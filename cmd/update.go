package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func UpdateCNTLMFile(p int) {
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
