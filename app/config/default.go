//+build !go1.16

package config

var defaultConfigSource = []byte(`application:
  port: 8080

logger:
  level: debug
  path: ./log

database:
  type: sqlite3
  file: file:database.sqlite?cache=shared&_fk=1
`)
