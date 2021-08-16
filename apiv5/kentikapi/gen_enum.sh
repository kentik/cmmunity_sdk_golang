#!/usr/bin/env bash

function stage() {
    COLOR="\e[95m"
    RESET="\e[0m"
    msg="$1"

    echo
    echo -e "$COLOR$msg$RESET"
}

function checkPrerequsites() {
    stage "Checking prerequisites"

    which enumer > /dev/null 2>&1
    # shellcheck disable=SC2181
    [[ $? != 0 ]] && echo "You need to install enumer with: go get github.com/alvaroloes/enumer" && exit 1

    echo "OK"
}

function genEnums() {
    stage "Generating enums"

    cd models/ || exit 1
    enumer -output enum_myenumtype.go  -type MyEnumType # update output and type to your enum

    echo "OK"
}

checkPrerequsites
genEnums
