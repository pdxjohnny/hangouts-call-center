package api

import (
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/ant0ine/go-json-rest/rest/test"

	"github.com/pdxjohnny/hangouts-call-center/variables"
)

func TestCall(t *testing.T) {
	err := os.Chdir("../../")
	if err != nil {
		log.Fatal(err)
	}

	handler := MakeHandler()

	// Use BackendToken to avoid login
	err = variables.LoadTokenKeys()
	if err != nil {
		t.Fatal(err)
	}
	// Test that call works
	path := variables.APIPathCallServer
	path = strings.Replace(path, ":number", "test number", 1)
	testReq := test.MakeSimpleRequest(
		"GET",
		"http://localhost"+path,
		nil,
	)
	testReq.Header.Set("Authorization", "Bearer "+variables.BackendToken)
	// Do not gzip the response
	testReq.Header.Set("Accept-Encoding", "identity")
	testRes := test.RunRequest(t, *handler, testReq)
	testRes.CodeIs(http.StatusOK)
	testRes.ContentTypeIsJson()
}
