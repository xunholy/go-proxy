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

func UpdateFile(cntlmValues map[string]string) error {
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
				lines[i] = field + "  " + val
			}
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
