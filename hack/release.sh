#!/bin/bash

set -e

SELF=$( realpath $0 )
BASEPATH=$( dirname $SELF )

TAG=${1:-devel}

if [ -z "${GITHUB_TOKEN}" ] || [ -z "${GITHUB_USER}" ]; then
	echo "make sure to set GITHUB_TOKEN and GITHUB_USER env vars"
	exit 2
fi

${BASEPATH}/build.sh

git add binaries && git ci -s -m "binaries: rebuild for tag ${TAG}"
git tag -a -m "hvinfo ${TAG}" ${TAG}
git push origin --tags
if  which github-release 2> /dev/null; then
	github-release release -t ${TAG} -r hvinfo
	github-release upload -t ${TAG} -r hvinfo \
		-n hvinfo-${TAG}-linux-amd64 -f binaries/hvinfo-${TAG}-linux-amd64
	github-release upload -t ${TAG} -r hvinfo \
		-n hvinfo-${TAG}-windows-amd64.exe -f binaries/hvinfo-${TAG}-windows-amd64.exe
fi
