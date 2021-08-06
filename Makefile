package_dir=./output/pkg

test:
	go test -v ./...
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

clear_output:
	rm -rf ./output

mkdir_package:
	if [ ! -d $(package_dir) ]; then mkdir -p $(package_dir) && echo created package_dir: "$(package_dir)"; fi

build_darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o output/bin/gomato_darwin ./cmd/gomato

package_darwin: build_darwin mkdir_package
	tar -czvf ./output/pkg/gomato_darwin_v2_0_0.tar.gz ./output/bin/gomato_darwin

gen_changelog:
	conventional-changelog -p angular -i CHANGELOG.md -s

install:
	yarn global add conventional-changelog-cli # changelog cli

gen_model:
	go generate app/ent/generate.go

gen: gen_model gen_changelog