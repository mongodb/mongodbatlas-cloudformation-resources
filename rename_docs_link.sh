#!/bin/bash

# ---
# A script to update old MongoDB Atlas documentation links in .md files.
#
# It searches for URLs like:
# https://www.mongodb.com/docs/atlas/reference/api-resources-spec/[v2/]#tag/...
#
# And updates them to the new format:
# https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/...
#
# WARNING: This script modifies files directly in-place WITHOUT creating backups.
# It is highly recommended to have your changes committed to Git before running.
# ---

# Exit immediately if a command exits with a non-zero status.
set -e

echo "🔍 Starting search for .md files to update..."
echo "⚠️ WARNING: Files will be modified in-place without backups."

# Find all markdown files and process them with a Perl one-liner.
find . -type f -name "*.md" | while IFS= read -r file; do
  # Use perl for in-place editing. The -i flag without an extension modifies the file directly.
  # The regex is updated to match the new base URL and hash format inside markdown links.
  perl -i -pE '
    s{
        (https://www.mongodb.com/docs/atlas/reference/api-resources-spec/) # $1: The old base URL
        (?:v2/)?                                                          # Optional, non-capturing "v2/" segment
        (\#tag/[^)]+)                                                     # $2: The hash fragment, stopping at the markdown link parenthesis
    }
    {
        # --- Start of replacement logic ---
        my $new_base = "https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/";
        my $hash = $2;
        my $target_url;

        # Case 1: Match #tag/Resource/operation/OperationId
        if ($hash =~ m{^#tag/([^/]+)/operation/([A-Za-z0-9_-]+)$}) {
            my $op_id = lc($2); # Lowercase the operation ID
            $target_url = $new_base . "operation/operation-" . $op_id;
        }
        # Case 2: Match #tag/Resource
        elsif ($hash =~ m{^#tag/([^/]+)$}) {
            my $resource = lc($1); # Lowercase the resource name
            $resource =~ s/[^a-z0-9]+/-/g; # Replace non-alphanumeric chars with a dash
            $target_url = $new_base . "group/endpoint-" . $resource;
        }
        # Case 3: Fallback if hash format is unexpected
        else {
            $target_url = $new_base; # Fallback to the new base URL
        }

        # Return the calculated target URL for the replacement
        $target_url;
        # --- End of replacement logic ---
    }gex; # g: global, e: execute code, x: ignore whitespace in regex
  ' "$file"
done

echo "✅ Update process complete."
echo "Files have been modified directly."
echo "Please review the changes with 'git diff' before committing."
