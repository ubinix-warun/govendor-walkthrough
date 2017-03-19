package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"ubinix.com/warun/echo"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": echo.Echo("helloworld!")})
	}))
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
