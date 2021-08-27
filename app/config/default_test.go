package config

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_defaultConfigSource(t *testing.T) {
	fp, err := ioutil.ReadFile("./metadata/config.example.yaml")
	require.NoError(t, err)
	t.Logf("%s", fp)
	require.Equal(t, fp, defaultConfigSource)
	assert.NotPanics(t, func() {
		_ = LoadDefaultConfig()
	})
}
