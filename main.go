package main

import (
	_ "crypto/sha256"
	_ "crypto/sha512"
	"github.com/tamalsaha/go-oneliners"
	_ "github.com/docker/distribution/reference"
	"github.com/heroku/docker-registry-client/registry"
)

func main() {
	url      := "https://registry-1.docker.io/"
	username := "" // anonymous
	password := "" // anonymous
	hub, err := registry.New(url, username, password)
	oneliners.FILE(err)

	repositories, err := hub.Repositories()
	oneliners.FILE(repositories)

	tags, err := hub.Tags("tigerworks/labels")
	oneliners.FILE(tags, err)

	manifest, err := hub.Manifest("tigerworks/labels", "latest")
	data, err := manifest.MarshalJSON()
	oneliners.FILE(string(data), err)

	digest, err := hub.ManifestDigest("tigerworks/labels", "latest")
	oneliners.FILE(digest, err)
}