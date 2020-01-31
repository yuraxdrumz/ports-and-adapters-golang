#!/bin/bash

# CHANGE PROJECT NAME HERE!
name=${PWD##*/}

# variables
SCRIPT_DIR="$( cd "$(dirname "$0")" ; pwd -P )"
BUILD_DIR="$(dirname "$SCRIPT_DIR")/build/build"
RELEASE_DIR="$(dirname "$SCRIPT_DIR")/build/release"

# init
echo "building project..."
rm -rf $BUILD_DIR
rm -rf $RELEASE_DIR
mkdir -p $BUILD_DIR
mkdir -p $RELEASE_DIR

# build
os_archs=(linux darwin windows)
for os in "${os_archs[@]}"
do
    if [[ $os == "windows" ]]; then
        echo "building: ${os}/amd64"
        GOOS=${os} GOARCH=amd64 go build -o "${BUILD_DIR}/${os}/${name}.exe"
    else 
        echo "building: ${os}/amd64"
        GOOS=${os} GOARCH=amd64 go build -o "${BUILD_DIR}/${os}/${name}"
    fi
done

# pack
list=$(find ${BUILD_DIR} -type f)
for file in ${list}
do
    os=$(echo ${file} | awk -F${BUILD_DIR}/ '{print $2}' | awk -F'/' '{print $1}')
    arch=$(echo ${file} | awk -F${BUILD_DIR}/ '{print $2}' | awk -F'/' '{print $2}' | awk -F'.exe' '{print $1}')

    echo "archiving: ${os}/${arch}"
    basename=$(basename "$file" | cut -d. -f1)
    mkdir -p ${RELEASE_DIR}/${os}
    if [[ $os == "windows" ]]; then 
        mv ${file} "${RELEASE_DIR}/${os}/${basename}.exe"
        zip -jq ${RELEASE_DIR}/${basename}-${os}.zip "${RELEASE_DIR}/${os}/${basename}.exe"
    else
        mv ${file} "${RELEASE_DIR}/${os}/${basename}"
        zip -jq ${RELEASE_DIR}/${basename}-${os}.zip "${RELEASE_DIR}/${os}/${basename}"
    fi
done

# init
rm -rf $BUILD_DIR

echo "finished!"