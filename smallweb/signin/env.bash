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
#第三方库和包,.gitignore 要忽略掉
export LIBDIR="$(dirname "$PWDDIR")/golibs"
#自己的 库和包
export LIB="$(dirname "$PWDDIR")/lib"
#gopath
export GOPATH=$LIBDIR:$LIB:$PWDDIR
echo $GOPATH

