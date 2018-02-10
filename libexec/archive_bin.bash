#!/bin/bash

PROJECT=$1
shift
TARGETDIR=$1
shift
ARCHIVEDIR=$1

SRTPATH=$(cd "$(dirname "$0")"; pwd)

if [ -z "$PROJECT" ]; then
    PROJECT=nuxim
fi

if [ -z "$TARGETDIR" ]; then
    TARGETDIR=/vobs/cache/tmp
fi

if [ -z "$ARCHIVEDIR" ]; then
    ARCHIVEDIR=/innova/archived/
fi

cd $TARGETDIR

if [ -d $TARGETDIR/$PROJECT ]; then
    rm -rf $TARGETDIR/$PROJECT
fi

mkdir $PROJECT && cd $PROJECT

mkdir data
if [ -d $SRTPATH/../data/$PROJECT ]; then
    rsync -a $SRTPATH/../data/$PROJECT data/
fi

mkdir conf
if [ -d $SRTPATH/../conf/$PROJECT ]; then
    rsync -a $SRTPATH/../conf/$PROJECT conf/
fi

mkdir -p bin

rsync -a $SRTPATH/../bin/$PROJECT bin/

cd $TARGETDIR && tar czf ${PROJECT}_bin.tar.gz ${PROJECT}

cp ${PROJECT}_bin.tar.gz $ARCHIVEDIR
