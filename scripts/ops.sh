#!/bin/bash

CMD=$1

function usage(){
    echo "Usage: $0 <commands>"
    echo ""
    echo "commands:"
    echo "  gen     Generate codes"
    echo "  clean   Cleen project"
}


function javaClean(){
    if [ -d ./java/src ]; then
        rm -rf ./java/src
        mkdir -p ./java/src/main/java
    else
        mkdir -p ./java/src/main/java
    fi
}

function goClean(){
    if [ -d ./go/pkg ]; then
        rm -rf ./go/pkg
        mkdir -p ./go/pkg
    else
        mkdir -p ./go/pkg
    fi
}


function build(){
    goClean
    javaClean
    go run ./tools/protoc/main.go
    pushd ./go/pkg
        go mod tidy
    popd
}

if [ "$#" -ne 1 ]; then
    usage
    exit
fi

case $CMD in
    "gen")
        build
        ;;
    "clean")
        goClean
        javaClean
        ;;
    *)
        usage
        ;;
esac