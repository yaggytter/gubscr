#!/bin/sh -x

set -eu

LOCALPATH=${1}
BRANCH=${2}
USER=${3}
REPO=${4}

TMPDIR="${LOCALPATH}/${USER}/${REPO}/${BRANCH}"

mkdir -p ${TMPDIR}

cd ${TMPDIR}

git clone https://github.com/${USER}/${REPO}.git
git checkout ${BRANCH}

ctags -R
