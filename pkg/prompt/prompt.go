package prompt

import (
	"bufio"
	"io"
	"os"
)

var input io.Reader = os.Stdin

func GetInput() (string, error) {
	reader := bufio.NewReader(input)
	output, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return output, nil
}
