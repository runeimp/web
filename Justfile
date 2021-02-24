
PROJECT_NAME := 'Web Lib'
DISTRO_NAME := 'web'
VERSION_FILE := 'web.go'

alias cover := test-coverage
alias ver := version

@_default:
	just _term-label "{{PROJECT_NAME}}"
	just _term-wipe
	just --list


# Test Suite
test +args='':
	@just _term-label "{{PROJECT_NAME}}"
	@just _term-wipe
	CGO_ENABLED=0 go test {{args}}

# Test Coverage
test-coverage +args='':
	#!/bin/sh
	just _term-wipe
	COVERFLAGS=''
	t=$(mktemp -t cover)
	echo "Go Test"
	CGO_ENABLED=0 go test ${COVERFLAGS} -coverprofile="${t}" {{args}}
	echo
	echo "Go Test Coverage"
	CGO_ENABLED=0 go tool cover -func="${t}"
	unlink "${t}"


# Helper recipe to change the terminal label
@_term-label label:
	printf "\033]0;{{label}}\007"

# Wipe the terminal buffer
@_term-wipe:
	#!/bin/sh
	if [[ ${#VISUAL_STUDIO_CODE} -gt 0 ]]; then
		clear
	elif [[ ${KITTY_WINDOW_ID} -gt 0 ]] || [[ ${#TMUX} -gt 0 ]] || [[ "${TERM_PROGRAM}" = 'vscode' ]]; then
		printf '\033c'
	elif [[ "$(uname)" == 'Darwin' ]] || [[ "${TERM_PROGRAM}" = 'Apple_Terminal' ]] || [[ "${TERM_PROGRAM}" = 'iTerm.app' ]]; then
		osascript -e 'tell application "System Events" to keystroke "k" using command down'
	elif [[ -x "$(which tput)" ]]; then
		tput reset
	elif [[ -x "$(which reset)" ]]; then
		reset
	else
		clear
	fi


# Prints the compiler or interpreter version(s)
@version:
	cat {{VERSION_FILE}} | grep -F 'const Version' | cut -d'"' -f2
	# go version
