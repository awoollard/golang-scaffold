package main

import (
	"encoding/json"
	"golang-scaffold/model"
	"io"
	"strings"
)

// GetSha Gets the current git revision
func GetSha(r io.Reader) (string, error) {
	shaFileBytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(shaFileBytes)), nil
}

// GetMetadata Parses the meta.json file
func GetMetadata(r io.Reader) (model.MetaItem, error) {
	fileBytes, _ := io.ReadAll(r)
	var metaItem model.MetaItem
	err := json.Unmarshal(fileBytes, &metaItem)
	if err != nil {
		return model.MetaItem{}, err
	}
	return metaItem, nil
}
