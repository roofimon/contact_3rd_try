package contact

import "github.com/ant0ine/go-json-rest/rest"

type Info struct {
	Name  string
	Email string
}

var All = &rest.Route{HttpMethod: "GET", PathExp: "/contact", Func: all}
var Get = &rest.Route{HttpMethod: "GET", PathExp: "/contact/:id", Func: get}

func all(w rest.ResponseWriter, r *rest.Request) {
	infos := []Info{
		Info{Name: "roof", Email: "roofimon@gmail.com"},
		Info{Name: "foor", Email: "foorimon@gmail.com"},
	}
	w.WriteJson(infos)
}

func get(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(Info{Name: "roof", Email: "roofimon@gmail.com"})
}
