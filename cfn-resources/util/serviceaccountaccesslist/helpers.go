// Copyright 2026 MongoDB Inc
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

package serviceaccountaccesslist

import (
	"context"
	"net/http"

	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

const ItemsPerPage = 500

// ListPageFunc is a function type that retrieves a page of access list entries.
type ListPageFunc func(ctx context.Context, pageNum int) (*admin.PaginatedServiceAccountIPAccessEntry, *http.Response, error)

// FindAccessListEntry iterates through access list pages looking for a specific entry.
// The firstPage can be provided to skip an API call (useful for Create operation which returns the first page).
func FindAccessListEntry(
	ctx context.Context,
	listPageFunc ListPageFunc,
	cidrOrIP string,
	firstPage *admin.PaginatedServiceAccountIPAccessEntry,
) (*admin.ServiceAccountIPAccessListEntry, error) {
	currentPage := 1
	page := firstPage

	for {
		if page == nil {
			var apiResp *http.Response
			var err error
			page, apiResp, err = listPageFunc(ctx, currentPage)
			if err != nil {
				if apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
					return nil, nil
				}
				return nil, err
			}
		}

		if page.Results != nil {
			for i := range *page.Results {
				entry := &(*page.Results)[i]
				entryIP := ""
				entryCIDR := ""
				if entry.IpAddress != nil {
					entryIP = *entry.IpAddress
				}
				if entry.CidrBlock != nil {
					entryCIDR = *entry.CidrBlock
				}

				if entryIP == cidrOrIP || entryCIDR == cidrOrIP {
					return entry, nil
				}
			}
		}

		if page.Results == nil || len(*page.Results) == 0 || len(*page.Results) < ItemsPerPage {
			break
		}

		currentPage++
		page = nil
	}

	return nil, nil
}

// ListAllAccessListEntries retrieves all access list entries across all pages.
func ListAllAccessListEntries(
	ctx context.Context,
	listPageFunc ListPageFunc,
) ([]admin.ServiceAccountIPAccessListEntry, error) {
	var allEntries []admin.ServiceAccountIPAccessListEntry
	currentPage := 1

	for {
		page, apiResp, err := listPageFunc(ctx, currentPage)
		if err != nil {
			if apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
				return allEntries, nil
			}
			return nil, err
		}

		if page.Results != nil {
			allEntries = append(allEntries, *page.Results...)
		}

		if page.Results == nil || len(*page.Results) == 0 || len(*page.Results) < ItemsPerPage {
			break
		}

		currentPage++
	}

	return allEntries, nil
}
