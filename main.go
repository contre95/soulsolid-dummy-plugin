package main

import (
	"github.com/contre95/soulsolid-dummy-plugin/dummy"
	"github.com/contre95/soulsolid/src/features/downloading"
)

// NewDownloader creates a new dummy downloader instance
func NewDownloader(cfg map[string]interface{}) (downloading.Downloader, error) {
	return dummy.NewDummyDownloader(), nil
}
