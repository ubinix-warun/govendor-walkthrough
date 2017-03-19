# govendor-walkthrough

```bash
gvm use go1.6
export GOPATH=`pwd`

go get -u github.com/kardianos/govendor
export PATH=$PATH:`pwd`/bin

cd src/ubinix.com/warun/echo
govendor init
govendor add github.com/ant0ine/go-json-rest/rest

# create fn-echo() on ubinix.com/warun/echo

govendor install +local

# create test-echo() on package

govendor test +local

# remove vendor on echo and create echoServer

rm vendor -Rf
cd ..
mkdir echoServer

# create main(), call echo() from package

govendor install +local
echoServer
"helloworld!"

# recreate "github.com/ant0ine/go-json-rest/rest"

govendor init
govendor add github.com/ant0ine/go-json-rest/rest

# create Helloworld JSON with "go-json-rest" for run!

govendor install +local
echoServer

open 127.0.0.1:8080

```

# BONUS! heroku
```bash

# edit .env and Procfile for new main()!

PORT=8080 heroku local

# create app.json and push vendor src to git!

heroku create
heroku git:remote -a __app_name__

# FIX dep.!

govendor fetch github.com/ant0ine/go-json-rest/rest


```


Inspired: 
* https://gocodecloud.com/blog/2016/03/29/go-vendoring-beginner-tutorial/
* https://github.com/github/gitignore/blob/master/Go.gitignore
* https://github.com/ubinix-warun/prj-golang-template