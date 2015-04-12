package contact

import "github.com/ant0ine/go-json-rest/rest"

var get = &rest.Route{HttpMethod: "GET", PathExp: "/contact", Func: contact}

var contact = func(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Add": "Added"})
}
