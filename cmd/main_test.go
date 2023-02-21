package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jollej/item-catalog/pgm/models"
)

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

// checkResponseCode is a simple utility to check the response code
// of the response
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetItems(t *testing.T) {
	// Create a New Server Struct
	s := CreateNewServer()
	// Mount Handlers
	s.MountHandlers()

	// Create a New Request
	req, _ := http.NewRequest("GET", "/items", nil)

	// Execute Request
	response := executeRequest(req, s)

	// Check the response code
	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestGetItem(t *testing.T) {
	// Create a New Server Struct
	s := CreateNewServer()
	// Mount Handlers
	s.MountHandlers()

	// Create a New Request
	req, _ := http.NewRequest("GET", "/items/1", nil)

	// Execute Request
	response := executeRequest(req, s)
	// Check the response code
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCreateItem(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	item := &models.Item{Id: "3", Description: "Item from test", Category: "Test_category"}
	itemJson, err := json.Marshal(item)
	jsonRequest := bytes.NewReader(itemJson)

	if err != nil {
		log.Panic(err)
	}
	req, _ := http.NewRequest("POST", "/items", jsonRequest)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestDeleteItem(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	req, _ := http.NewRequest(http.MethodDelete, "/items/2", nil)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
}
