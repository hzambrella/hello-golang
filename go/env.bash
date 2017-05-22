#/!bin/bash
PWDDIR=`pwd`
export LIBDIR="$(dirname "$PWDDIR")/tool"
export GOPATH=$PWDDIR/go:$LIBDIR
echo $GOPATH
