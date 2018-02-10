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

rsync -a $SRTPATH/../libexec .

mkdir data
if [ -d $SRTPATH/../data/$PROJECT ]; then
    rsync -a $SRTPATH/../data/$PROJECT data/
fi

mkdir conf
if [ -d $SRTPATH/../conf/$PROJECT ]; then
    rsync -a $SRTPATH/../conf/$PROJECT conf/
fi

mkdir -p src

rsync -a $SRTPATH/../src/$PROJECT src/

cd $TARGETDIR && tar czf ${PROJECT}_src.tar.gz ${PROJECT}

cp ${PROJECT}_src.tar.gz $ARCHIVEDIR
