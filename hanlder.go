package contact

import "github.com/ant0ine/go-json-rest/rest"

type Contact struct {
	Name  string
	Email string
}

var All = &rest.Route{HttpMethod: "GET", PathExp: "/contact", Func: all}
var Get = &rest.Route{HttpMethod: "GET", PathExp: "/contact/:id", Func: get}

func all(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(Contact{Name: "roof", Email: "roofimon@gmail.com"})
}

func get(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(Contact{Name: "roof", Email: "roofimon@gmail.com"})
}
