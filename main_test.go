package contact

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
)

func setup() rest.ResourceHandler {
	handler := rest.ResourceHandler{
		DisableJsonIndent: true,
		ErrorLogger:       log.New(ioutil.Discard, "", 0),
	}
	handler.SetRoutes(All, Get)
	return handler
}

func TestGetAContact(t *testing.T) {
	handler := setup()
	recorded := test.RunRequest(t, &handler,
		test.MakeSimpleRequest("GET", "http://1.2.3.4/contact/1", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`{"Add":"Added"}`)
}

func TestGet(t *testing.T) {
	handler := setup()
	recorded := test.RunRequest(t, &handler,
		test.MakeSimpleRequest("GET", "http://1.2.3.4/contact", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`{"Add":"Added"}`)
}
