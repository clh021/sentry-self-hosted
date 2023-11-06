.PHONY: build
gitTime=$(shell date +%Y%m%d%H%M%S)
gitCID=$(shell git rev-parse HEAD | cut -c1-8)
gitTag=$(shell git tag --list --sort=version:refname 'v*' | tail -1)
gitCount=$(shell git log --pretty=format:'' | wc -l)/$(shell git rev-list --all --count)
buildStr=${gitTime}.${gitCID}.${gitTag}.${gitCount}
build:
	CGO_ENABLED=1 go build -ldflags "-s -w -X main.build=${buildStr}" -o tmp/app *.go