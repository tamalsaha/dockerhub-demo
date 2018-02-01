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
	// https://github.com/kubernetes/kubernetes/blob/master/pkg/util/parsers/parsers_test.go
	repo, tag, digestHash, err := parsers.ParseImageName("tigerworks/labels:latest")
	oneliners.FILE(err)
	oneliners.FILE("repo = ", repo)
	oneliners.FILE("tag = ", tag)
	oneliners.FILE("digest = ", digestHash)

	tags, err := hub.Tags("tigerworks/labels")
	oneliners.FILE(tags, err)

	//m2, err := hub.Manifest("tigerworks/labels", "latest")
	//oneliners.FILE(m2.Name, m2.Tag)
	//d2, err := m2.MarshalJSON()
	//oneliners.FILE(string(d2))

	manifest, err := hub.ManifestV2("tigerworks/labels", "latest")
	oneliners.FILE(manifest.Config.Digest.Encoded())
	oneliners.FILE(manifest.Layers[0].Digest.Encoded())

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
