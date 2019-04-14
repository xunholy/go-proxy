package os_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xUnholy/go-proxy/internal/os"
)

func TestGetConfigurationPath(t *testing.T) {
	_, err := os.GetConfigurationPath("linux")
	assert.Equal(t, nil, err)
	_, err = os.GetConfigurationPath("windows")
	assert.Equal(t, nil, err)
	_, err = os.GetConfigurationPath("darwin")
	assert.Equal(t, nil, err)
	_, err = os.GetConfigurationPath("unknown")
	assert.NotEqual(t, nil, err)
}
