package rackspace

import (
	"testing"

	"github.com/rancher/machine/libmachine/drivers"
	"github.com/stretchr/testify/assert"
)

func TestSetConfigFromFlags(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"rackspace-region":        "REGION",
			"rackspace-username":      "user",
			"rackspace-api-key":       "KEY",
			"rackspace-endpoint-type": "publicURL",
		},
		CreateFlags: driver.GetFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)

	assert.NoError(t, err)
	assert.Empty(t, checkFlags.InvalidFlags)
}
