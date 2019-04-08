package cmd

import "fmt"

func makeProxyURL(port int) string {
	return fmt.Sprintf("http://localhost:%d", port)
}
