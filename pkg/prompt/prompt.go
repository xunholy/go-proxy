package prompt

import (
	"bufio"
	"io"
)

func GetInput(input io.Reader) (string, error) {
	reader := bufio.NewReader(input)
	output, err := reader.ReadString('\n')
	if err != nil {
		return output, err
	}
	return output, err
}
