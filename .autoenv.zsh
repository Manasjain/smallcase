if [[ -z $ORIG_PATH ]]; then
    export ORIG_PATH=$PATH
    export GOPATH=`pwd`/../..
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    export AppConf="./config.json"
    echo "Welcome to smallcase project"
else
    export ORIG_PATH=$PATH
    echo "Not set paths yet"
fi
