.PHONY: update-atlas-sdk
update-atlas-sdk: ## Update the atlas-sdk dependency
	go install github.com/icholy/gomajor@v0.9.5
	## Fetch the latest major version and update imports.
	gomajor get go.mongodb.org/atlas-sdk/v20230201001@latest