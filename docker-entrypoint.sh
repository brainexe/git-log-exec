#!/bin/sh

set -e

REPO_DIRECTORY=/repo

if [ -z "$@" ]
then
  echo "Please pass more parameters, at least a -command which gets executed on each matched commit."
  /git-log-exec --help
  exit 1
elif [ ! -d "/repo" ]; then
  REPO=$1
  echo "Clone ${REPO} into local docker container"
  git clone ${REPO} $REPO_DIRECTORY
  shift 1
fi

/git-log-exec --directory ${REPO_DIRECTORY} "$@"
cat out.csv