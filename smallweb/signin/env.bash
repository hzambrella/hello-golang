#/!bin/bash

# 设定git库地址转换, 以便解决部分包的库被墙的问题
#git config --global url."git@git.ot24.net:".insteadOf "https://git.ot24.net"
#git config --global url."https://github.com/golang/".insteadOf "https://go.googlesource.com/"

PWDDIR=`pwd`
#MVC架构
#view
mkdir -p src    
cd src
mkdir -p public
#control
mkdir -p routes
#model
mkdir -p model
cd ..
#库和包
export LIBDIR="$(dirname "$PWDDIR")/golibs"
#gopath
export GOPATH=$LIBDIR:$PWDDIR
echo $GOPATH

