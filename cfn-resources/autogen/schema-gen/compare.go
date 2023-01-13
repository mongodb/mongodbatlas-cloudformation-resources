package main

import (
	"log"
	"os"

	"github.com/nsf/jsondiff"
)

// CompareJsonFiles Compares the JSON Content in given Files
func CompareJSONFiles(resourceName, existingFilePath, latestFilePath string) (diffJSON string, err error) {
	existingAPIContent, err := os.ReadFile(existingFilePath)
	if err != nil {
		return
	}

	latestAPIContent, err := os.ReadFile(latestFilePath)
	if err != nil {
		return
	}

	differences, diffJSON := jsondiff.Compare(existingAPIContent, latestAPIContent, &jsondiff.Options{SkipMatches: true})
	if differences > 0 {
		log.Printf("found difference in %s schema ", resourceName)
		err = os.WriteFile(resourceName+"-"+diffFile, []byte(diffJSON), 0600)
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}
