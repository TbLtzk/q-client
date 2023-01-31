#!/bin/sh

QVERSION=$1
QVERSION_PATH=params/version.go

QMAJOR=$(cut -d '.' -f 1 <<< $QVERSION)
QMINOR=$(cut -d '.' -f 2 <<< $QVERSION)
QPATCH=$(cut -d '.' -f 3 <<< $QVERSION)

if [ -z "$QMAJOR" ] || [ -z "$QMINOR" ] || [ -z "$QPATCH" ]
then
    echo "invalid version: $QVERSION"
    exit 1
fi

echo "Major: $QMAJOR"
echo "Minor: $QMINOR"
echo "Path: $QPATCH"

sed -r -i -E "s/QVersionMajor = [0-9]+/QVersionMajor = $QMAJOR/g" $QVERSION_PATH
sed -r -i -E "s/QVersionMinor = [0-9]+/QVersionMinor = $QMINOR/g" $QVERSION_PATH
sed -r -i -E "s/QVersionPatch = [0-9]+/QVersionPatch = $QPATCH/g" $QVERSION_PATH

git add -- .
git commit -m "Update q-client version"
git push origin $(git branch --show-current)
git tag -f v$QVERSION
git push -f origin v$QVERSION