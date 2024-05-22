package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

// TestHomeHandler tests the homeHandler function
func TestHomeHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(homeHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check if the response body contains the expected title and content
    body := rr.Body.String()
    if !strings.Contains(body, "<title>Home</title>") {
        t.Errorf("handler returned unexpected body: title not found")
    }
    if !strings.Contains(body, "Home") {
        t.Errorf("handler returned unexpected body: Home not found")
    }
}

// TestAboutHandler tests the aboutHandler function
func TestAboutHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/about", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(aboutHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check if the response body contains the expected title and content
    body := rr.Body.String()
    if !strings.Contains(body, "<title>About</title>") {
        t.Errorf("handler returned unexpected body: title not found")
    }
    if !strings.Contains(body, "About") {
        t.Errorf("handler returned unexpected body: About not found")
    }
}

// contains checks if a string is contained in another string
func contains(str, substr string) bool {
    return strings.Contains(str, substr)
}

