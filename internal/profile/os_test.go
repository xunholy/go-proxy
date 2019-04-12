package profile_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xUnholy/go-proxy/internal/profile"
)

func TestGetConfigurationPath(t *testing.T) {
	_, err := profile.GetConfigurationPath("linux")
	assert.Equal(t, nil, err)
	_, err = profile.GetConfigurationPath("windows")
	assert.Equal(t, nil, err)
	_, err = profile.GetConfigurationPath("darwin")
	assert.Equal(t, nil, err)
	_, err = profile.GetConfigurationPath("unknown")
	assert.NotEqual(t, nil, err)
}
