package contact

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
)

func TestHandler(t *testing.T) {
	handler := rest.ResourceHandler{
		DisableJsonIndent: true,
		// make the test output less verbose by discarding the error log
		ErrorLogger: log.New(ioutil.Discard, "", 0),
	}
	handler.SetRoutes(get)
	recorded := test.RunRequest(t, &handler,
		test.MakeSimpleRequest("GET", "http://1.2.3.4/contact", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`{"Add":"Added"}`)
}
