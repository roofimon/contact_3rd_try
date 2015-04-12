package contact

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
)

func makeHandler() http.Handler {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, _ := rest.MakeRouter(All, Get)
	api.SetApp(router)
	return api.MakeHandler()
}

func TestGet(t *testing.T) {
	expected := Info{Name: "roof", Email: "roofimon@gmail.com"}
	body := Info{}
	handler := makeHandler()
	recorded := test.RunRequest(t, handler,
		test.MakeSimpleRequest("GET", "http://1.2.3.4/contact/1", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()

	recorded.DecodeJsonPayload(&body)
	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Body %v expected got %v", expected, body)
	}
}

func TestAll(t *testing.T) {
	expected := []Info{
		Info{Name: "roof", Email: "roofimon@gmail.com"},
		Info{Name: "foor", Email: "foorimon@gmail.com"},
	}
	body := []Info{}
	handler := makeHandler()
	recorded := test.RunRequest(t, handler,
		test.MakeSimpleRequest("GET", "http://1.2.3.4/contact", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.DecodeJsonPayload(&body)
	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Body %v expected got %v", expected, body)
	}
}
