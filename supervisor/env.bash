PWDDIR=`pwd`

export LIBDIR="$(dirname "$PWDDIR")/tool"
GOPATH=$PWDDIR:$LIBDIR
echo "now gopath:"$GOPATH
