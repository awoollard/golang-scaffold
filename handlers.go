package main

import (
	"encoding/json"
	"golang-scaffold/model"
	"log"
	"net/http"
	"os"
)

// GetStatus Handler for the status endpoint
func GetStatus(w http.ResponseWriter, r *http.Request) {
	// Open the metadata file
	metadataFile, err := os.Open("/meta.json")
	if err != nil {
		log.Fatal(err)
	}
	defer metadataFile.Close()

	// Open the Git SHA file
	shaFile, err := os.Open("/git.sha")
	if err != nil {
		log.Fatal(err)
	}
	defer shaFile.Close()

	// Parse contents of the files
	metaItem, err := GetMetadata(metadataFile)
	if err != nil {
		log.Fatal(err)
	}

	sha, err := GetSha(shaFile)
	if err != nil {
		log.Fatal(err)
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(
		model.StatusResponse{
			Author: []model.StatusItem{{metaItem.Description, sha, metaItem.Version}},
		})

	// Handle errors
	if err != nil {
		log.Fatal(err)
	}
}

// GetRoot Handler for the root endpoint
func GetRoot(w http.ResponseWriter, r *http.Request) {
	// Write a simple Hello World response
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Fatal(err)
	}
}
