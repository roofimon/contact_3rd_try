package contact

import "github.com/ant0ine/go-json-rest/rest"

var All = &rest.Route{HttpMethod: "GET", PathExp: "/contact", Func: allContact}
var Get = &rest.Route{HttpMethod: "GET", PathExp: "/contact/:id", Func: contact}

func allContact(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Add": "Added"})
}

func contact(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Add": "Added"})
}
