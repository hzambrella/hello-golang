#/!bin/bash
PWDDIR=`pwd`
export LIBDIR="$(dirname "$PWDDIR")/tool"
export GOPATH=$PWDDIR/studydb:$PWDDIR/vendor:$LIBDIR
echo $GOPATH
