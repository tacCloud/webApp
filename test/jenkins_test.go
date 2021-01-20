package test

import (
	"github.com/tacCloud/webApp/blah"
	"fmt"
	"testing"

)

// Test for /movie/list
func TestListInventory(t *testing.T) {
	for i:= 0; i < len(blah.FakeItems); i++ {
		fmt.Printf("%v\n", blah.FakeItems[i]);
	}
	
	/*
	req, err := http.NewRequest("GET", "/v1/movie/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.MovieList)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	result := rr.Body.String()
	expected := `"message":"20 movie(s) have been discovered.","success":true`
	if !strings.Contains(result, expected) {
		t.Errorf("Handler returned unexpected body: got %v want string containing %v", result, expected)
	}
	*/
}
