# govendor-walkthrough

```bash
gvm use go1.6
export GOPATH=`pwd`

go get -u github.com/kardianos/govendor
export PATH=$PATH:`pwd`/bin

cd src/ubinix.com/warun/echo
govendor init
govendor add github.com/ant0ine/go-json-rest/rest

create fn-echo() on ubinix.com/warun/echo

govendor install +local
```

Inspired: https://gocodecloud.com/blog/2016/03/29/go-vendoring-beginner-tutorial/