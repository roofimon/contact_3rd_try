package contact

import "github.com/ant0ine/go-json-rest/rest"

var All = &rest.Route{HttpMethod: "GET", PathExp: "/contact", Func: allContact}
var Get = &rest.Route{HttpMethod: "GET", PathExp: "/contact/:id", Func: contact}

var allContact = func(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Add": "Added"})
}

var contact = func(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Add": "Added"})
}
