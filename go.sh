#!/bin/bash
# . ./go.sh

if [[ ("$SHELL" != "$(which zsh)") && ("$0" != "bash" && "$0" != "-bash" && "$0" != "/bin/bash") ]]; then
	echo "Please run this script as: . ./go.sh  from your project directory"
else
	PROJECT_DIR=$(pwd)
	if [ -d ~/.go ]; then
		# Add ~/.go to our path if needed.
		if [[ ":$GOPATH:" == *":$HOME/.go:"* ]]; then
			echo "Path contains .go"
		else
			# is it empty?
			if [ -z "${GOPATH-}" ]; then
				export GOPATH=~/.go
			else
				export GOPATH=~/.go:$GOPATH
			fi
		fi
		export GOBIN=~/.go/bin
	else
		echo "NOTE: No ~/.go directory found. Not setting a baseline go path for modules"
		export GOBIN=$PROJECT_DIR/bin
	fi

	if [ -z "${GOPATH}" ]; then
		export GOPATH=$PROJECT_DIR
	else
		if [[ ":$GOPATH:" == *":$PROJECT_DIR:"* ]]; then
			echo "Project dir already in path"
		else
			export GOPATH=$GOPATH:$PROJECT_DIR
		fi
	fi

	if [ -d /usr/local/go/bin ]; then
		echo "Local go distribution detected - altering PATH and GOROOT"
		export GOROOT=/usr/local/go
		if [[ ":$PATH:" == *":/usr/local/go/bin:"* ]]; then
			true
		else
			export PATH=/usr/local/go/bin:$PATH
		fi
	fi

	if [ -d $PROJECT_DIR/conf/JSON_Modules ]; then
		export SNAREJSONCONFIGPATH=$PROJECT_DIR/conf/JSON_Modules
	fi

	export GO111MODULE=auto
fi
