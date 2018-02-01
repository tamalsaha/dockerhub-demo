package main

import (
	_ "crypto/sha256"
	_ "crypto/sha512"
	"encoding/json"
	"log"

	_ "github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/heroku/docker-registry-client/registry"
	"github.com/tamalsaha/go-oneliners"
	"k8s.io/kubernetes/pkg/util/parsers"
)

func main() {
	url := "https://registry-1.docker.io/"
	username := "" // anonymous
	password := "" // anonymous
	hub, err := registry.New(url, username, password)
	oneliners.FILE(err)

	// https://github.com/kubernetes/kubernetes/blob/a7a3dcfc527123b6cca15913fbb309172ef2c6e4/pkg/util/parsers/parsers.go#L33
	repo, tag, digestHash, err := parsers.ParseImageName("tigerworks/labels:latest")
	oneliners.FILE(err)
	oneliners.FILE("repo = ", repo)
	oneliners.FILE("tag = ", tag)
	oneliners.FILE("digest = ", digestHash)

	repositories, err := hub.Repositories()
	oneliners.FILE(repositories)

	tags, err := hub.Tags("tigerworks/labels")
	oneliners.FILE(tags, err)

	manifest, err := hub.ManifestV2("tigerworks/labels", "latest")
	//data, err := manifest.MarshalJSON()
	//oneliners.FILE(string(data), err)

	//digest, err := hub.ManifestDigest("tigerworks/labels", "latest")
	//oneliners.FILE(digest, err)

	oneliners.FILE(manifest.Config.Digest.Encoded())

	reader, err := hub.DownloadLayer("tigerworks/labels", manifest.Config.Digest)
	if err != nil {
		log.Fatalln(err)
	}

	var cfg types.ImageInspect
	err = json.NewDecoder(reader).Decode(&cfg)
	oneliners.FILE(err)
	defer reader.Close()

	oneliners.FILE(cfg.Config.Labels)
}
