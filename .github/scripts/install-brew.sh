#!/usr/bin/env bash
set eu

/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
# shellcheck disable=SC2016
(
	echo
	echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"'
) >>/home/runner/.bash_profile

sudo apt-get install build-essential
brew install gcc
