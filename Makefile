test:
	go test -v ./...
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

gen_changelog:
	conventional-changelog -p angular -i CHANGELOG.md -s

install:
	yarn global add conventional-changelog-cli # changelog cli

gen_model:
	go generate app/ent/generate.go

gen: gen_model gen_changelog