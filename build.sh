#!/bin/bash

#  The MIT License (MIT)
# 
#  Copyright (c) 2021 Micke Sommar <me@mickesommar.com>
# 
#  Permission is hereby granted, free of charge, to any person obtaining a copy
#  of this software and associated documentation files (the "Software"), to deal
#  in the Software without restriction, including without limitation the rights
#  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
#  copies of the Software, and to permit persons to whom the Software is
#  furnished to do so, subject to the following conditions:
# 
#  The above copyright notice and this permission notice shall be included in all
#  copies or substantial portions of the Software.
# 
#  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
#  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
#  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
#  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
#  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
#  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
#  SOFTWARE.

# Name   : build.sh
# Comment: Build application with git version number, to release directory.
#          WARNING, this script will remove untracked files (if exists).
#          This script will checkout to the latest git tag, remove all untracked files and build releases.
# Target : Linux, Mac OS and Windows (amd64, arm64)

# Options
BINARY_NAME=certls
SOURCEPATH=cmd/certls/main.go

# Version
VERSION=DEVELOPMENT
COMMIT=""
DATE=$(date +"%Y-%m-%d")
TIME=$(date +"%T")
OS_LIST=(linux windows darwin)
ARCH_LIST=(amd64 arm64)

# Use git to set version and git-cimmit-hash.
# Get version from tag.
GIT_TAG=$(git describe --abbrev=0 2>/dev/null)
if [ -n "$GIT_TAG" ];then
  VERSION=$GIT_TAG
  
  # Remove untracked files.
  if [ -n "$(git status --porcelain)" ];then
    echo "Untracked files exists, ALL UNTRACKED FILES WILL BE DELETED"
    read -n1 -p "Continue? [y,n]" doit 
    case $doit in  
      n|N) git switch -;exit ;; 
    esac
    
    # Clean up untracked files.
    git clean -fd
  fi

  # Checkout to latest git tag.
  git checkout tags/$GIT_TAG
 
  # Get short commit from git.
  COMMIT=$(git rev-parse --short HEAD 2>/dev/null)
 
  # Create releases directory, if missing.
  if [ ! -d "./releases" ];then
    mkdir ./releases
  fi
  
  # Clean up priviously build.
  rm -rf ./releases/*

  # Build
  for os in ${OS_LIST[@]}
  do
    for arch in ${ARCH_LIST[@]}
    do
      # Build app, using go.
      # Using ldflags -w -s, for removing debug info from binary, make the binary smaler.
      echo "build: ./releases/${BINARY_NAME} for $os $arch"
      env GOOS=$os GOARCH=$arch go build -ldflags="-w -s -X main.version=${VERSION} -X main.commitHash=${COMMIT} -X main.buildDate=${DATE} -X main.buildTime=${TIME}" \
                                         -o ./releases/${BINARY_NAME}-$os-$arch ${SOURCEPATH}
    done
  done

  # Switch back.
  git switch -
else
  echo "This is not a git repository with git tags (requried)"
fi