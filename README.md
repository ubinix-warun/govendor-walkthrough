# govendor-walkthrough

gvm use go1.6
export GOPATH=`pwd`

go get -u github.com/kardianos/govendor
export PATH=$PATH:`pwd`/bin

cd src/ubinix.com/warun/echo
govendor init
govendor add github.com/ant0ine/go-json-rest/rest