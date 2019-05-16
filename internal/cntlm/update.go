package cntlm

import (
	"io/ioutil"
	"log"
	"runtime"
	"strings"

	"github.com/xUnholy/go-proxy/internal/os"
)

type KeyPairValues struct {
	Key   string
	Value string
	Line  int
}

func contains(replaceCommented []string, check string) bool {
	for _, val := range replaceCommented {
		if val == check {
			return true
		}
	}
	return false
}

func UpdateFile(cntlmValues map[string]string) error {
	replaceCommented := []string{"PassLM", "PassNT", "PassNTLMv2"}
	file, err := os.GetConfigurationPath(runtime.GOOS)
	if err != nil {
		return err
	}
	input, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")
	for field, val := range cntlmValues {
		for i, line := range lines {
			if strings.Contains(line, field) {
				if strings.HasPrefix(line, "#") {
					if contains(replaceCommented, field) {
						lines[i] = field + "  " + val
						goto getOut
					} else {
						continue
					}
				} else {
					lines[i] = field + "  " + val
					goto getOut
				}
			}
		}
	getOut:
		continue
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
