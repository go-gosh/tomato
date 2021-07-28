//+build go1.16

package config

import _ "embed"

//go:embed metadata/config.example.yaml
var defaultConfigSource []byte
