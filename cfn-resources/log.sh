#!/usr/bin/env bash

# Log the given message at the given level. All logs are written to stderr with a timestamp.
function log {
  local -r level="$1"
  local -r message="$2"
  local -r timestamp=$(date +"%Y-%m-%d %H:%M:%S")
  local -r script_name="$(basename "$0")"
  >&2 echo -e "${timestamp} [${level}] [$script_name] ${message}"
}

function log_multiline {
  local -r level="$1"
  local -r message="$2"
  while read i ; do log "$level" "$i" ; done <<< "$message"
}


# Log the given message at INFO level. All logs are written to stderr with a timestamp.
function log_info {
  local -r message="$1"
  log_multiline "INFO" "$message"
}

# Log the given message at WARN level. All logs are written to stderr with a timestamp.
function log_warn {
  local -r message="$1"
  log_multiline "WARN" "$message"
}

# Log the given message at ERROR level. All logs are written to stderr with a timestamp.
function log_error {
  local -r message="$1"
  log_multiline "ERROR" "$message"
}
