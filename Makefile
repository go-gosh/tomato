package_dir=./output/pkg

# get os name
ifeq ($(OS),Windows_NT)
    detected_OS:=windows
    ext:=.exe
else
    detected_OS:=$(shell sh -c 'uname 2>/dev/null || echo Unknown' | tr 'A-Z' 'a-z')
    ext:=
endif

run: clear_output mkdir_package gen_model test build
	clear && ./output/bin/gomato_$(detected_OS)$(ext)

test: mkdir_package
	go test -v ./...
	go test -race -coverprofile=output/coverage.txt -covermode=atomic ./...

clear_output:
	rm -rf ./output

mkdir_package:
	if [ ! -d $(package_dir) ]; then mkdir -p $(package_dir) && echo created package_dir: "$(package_dir)"; fi

build:
	go build -o ./output/bin/gomato_$(detected_OS)$(ext) ./cmd/gomato

build_darwin:
	GOOS=darwin GOARCH=amd64 go build -o output/bin/gomato_darwin ./cmd/gomato

package_darwin: build_darwin mkdir_package
	tar -czvf ./output/pkg/gomato_darwin_v2_0_0.tar.gz ./output/bin/gomato_darwin

build_windows:
	GOOS=windows GOARCH=amd64 go build -o output/bin/gomato_windows.exe ./cmd/gomato

package_windows: build_windows mkdir_package
	zip -r ./output/pkg/gomato_windows_v2_0_0.zip ./output/bin/gomato_windows.exe

gen_changelog:
	conventional-changelog -p angular -i CHANGELOG.md -s

install:
	yarn global add conventional-changelog-cli # changelog cli

gen_model:
	go generate app/ent/generate.go

gen: gen_model gen_changelog
