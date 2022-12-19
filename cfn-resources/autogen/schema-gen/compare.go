package main

import (
	"fmt"
	"github.com/nsf/jsondiff"
	"os"
)

// CompareJsonFiles Compares the JSON Content in given Files
func CompareJsonFiles(resourceName, existingFilePath, latestFilePath string) (diffJson string, err error) {

	existingAPIContent, err := os.ReadFile(existingFilePath)
	if err != nil {
		return
	}

	latestAPIContent, err := os.ReadFile(latestFilePath)
	if err != nil {
		return
	}

	differences, diffJson := jsondiff.Compare(existingAPIContent, latestAPIContent, &jsondiff.Options{SkipMatches: true})
	if differences > 0 {
		fmt.Printf("diff %+v, val : %s", differences, diffJson)
		sendMail(resourceName, diffJson)
	}
	return
}
