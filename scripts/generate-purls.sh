#!/usr/bin/env bash
set -euo pipefail

# Loop over all resource directories in cfn-resources, skipping autogen
for resource_dir in cfn-resources/*; do
	if [ "$resource_dir" = "cfn-resources/autogen" ]; then
		continue
	fi
	if [ -d "$resource_dir" ] && [ -f "$resource_dir/Makefile" ]; then
		resource=$(basename "$resource_dir")
		compliance_dir="$resource_dir/compliance"
		bin_dir="$resource_dir/bin"
		binary="$bin_dir/bootstrap"

		echo "==> Building $resource"
		mkdir -p "$compliance_dir"
		(cd "$resource_dir" && make build)

		if [ ! -f "$binary" ]; then
			echo "No built binary found at $binary for $resource. Skipping purl generation."
			continue
		fi

		purl_file="$compliance_dir/purls.txt"
		echo "==> Generating purls for $resource"
		go version -m "$binary" | awk '$1 == "dep" || $1 == "=>" { print "pkg:golang/" $2 "@" $3 }' | LC_ALL=C sort | uniq >"$purl_file"
		echo "Generated $purl_file"
	fi

done
