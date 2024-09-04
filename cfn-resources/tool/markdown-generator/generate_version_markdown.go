// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const mongoDBPublisherID = "bb989456c78c398a858fef18f2ca1bfc1fbba082"
const ouputMarkdownFile = "resource-versions.md"

func main() {
	var regions = getAWSRegionsFromEnvVar()
	var resourceTypes = GetTypeNamesFromRpdkJSONFiles()

	// map of resource type name to list of values where each index corresponds to a region
	resourceVersions := getVersions(regions, resourceTypes)
	_ = generateMarkdown(ouputMarkdownFile, regions, resourceTypes, resourceVersions)
}

func getAWSRegionsFromEnvVar() []string {
	// Get the environment variable
	envVar := os.Getenv("AWS_REGIONS")
	if envVar == "" {
		log.Fatal("missing env variable AWS_REGIONS")
	}
	return strings.Split(envVar, ",")
}

func getVersions(regions, resourceTypes []string) map[string][]string {
	resourceValues := map[string][]string{}

	for _, typeName := range resourceTypes {
		resourceValues[typeName] = make([]string, len(regions))
	}

	for i, region := range regions {
		typeSummaries := getTypeSummariesOfRegion(region)

		for j := range typeSummaries {
			typeSummary := &typeSummaries[j]
			fmt.Printf("Region: %s, type name: %s, version: %s, date: %s\n", region, *typeSummary.TypeName, *typeSummary.LatestPublicVersion, typeSummary.LastUpdated.String())
			if _, exists := resourceValues[*typeSummary.TypeName]; exists {
				if typeSummary.LatestPublicVersion != nil && typeSummary.LastUpdated != nil {
					date := *typeSummary.LastUpdated
					resourceValues[*typeSummary.TypeName][i] = fmt.Sprintf("%s (%s)", *typeSummary.LatestPublicVersion, date.Format("2006-01-02"))
				}
			}
		}
	}
	return resourceValues
}

func generateMarkdown(filename string, regions, resourceTypes []string, resourceValues map[string][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, r := range regions {
		regions[i] = useNoBreakChar(r)
	}

	headers := []string{"Latest version"}
	headers = append(headers, regions...)
	_, _ = fmt.Fprintf(file, "%s\n", strings.Join(headers, "|"))

	divider := make([]string, len(headers))
	for i := range divider {
		divider[i] = "-"
	}
	_, _ = fmt.Fprintf(file, "%s\n", strings.Join(divider, "|"))

	for _, typeName := range resourceTypes {
		row := []string{fmt.Sprintf("**%s**", typeName)}
		for regionIndex := range regions {
			value := resourceValues[typeName][regionIndex]
			row = append(row, useNoBreakChar(value))
		}
		_, _ = fmt.Fprintf(file, "%s\n", strings.Join(row, "|"))
	}
	return nil
}

func getTypeSummariesOfRegion(region string) []types.TypeSummary {
	cfg, err := config.LoadDefaultConfig(context.Background())
	cfg.Region = region

	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}

	client := cloudformation.NewFromConfig(cfg)

	input := &cloudformation.ListTypesInput{
		Visibility: types.VisibilityPublic,
		Type:       types.RegistryTypeResource,
		Filters: &types.TypeFilters{
			PublisherId: aws.String(mongoDBPublisherID),
		},
	}

	paginator := cloudformation.NewListTypesPaginator(client, input)

	results := []types.TypeSummary{}
	// Iterate through paginated results
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			log.Fatalf("failed to fetch page, %v", err)
		}
		results = append(results, page.TypeSummaries...)
	}
	return results
}

// define no break '-' char for improved view of markdown
func useNoBreakChar(s string) string {
	return strings.ReplaceAll(s, "-", "&#x2011;")
}
