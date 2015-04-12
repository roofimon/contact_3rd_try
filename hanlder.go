package contact

import "github.com/ant0ine/go-json-rest/rest"

var All = &rest.Route{HttpMethod: "GET", PathExp: "/contact", Func: all}
var Get = &rest.Route{HttpMethod: "GET", PathExp: "/contact/:id", Func: get}

func all(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Add": "Added"})
}

func get(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Add": "Added"})
}
