package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListItems_Empty(t *testing.T) {
	// Reset items for test isolation
	items = make([]Item, 0)

	req := httptest.NewRequest("GET", "/api/items", nil)
	w := httptest.NewRecorder()

	ListItems(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result []Item
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("expected empty list, got %d items", len(result))
	}
}

func TestCreateItem_Success(t *testing.T) {
	items = make([]Item, 0)

	body := bytes.NewBufferString(`{"name": "Test Item"}`)
	req := httptest.NewRequest("POST", "/api/items", body)
	w := httptest.NewRecorder()

	CreateItem(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", w.Code)
	}

	var result Item
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result.Name != "Test Item" {
		t.Errorf("expected name 'Test Item', got '%s'", result.Name)
	}
}

func TestCreateItem_MissingName(t *testing.T) {
	body := bytes.NewBufferString(`{}`)
	req := httptest.NewRequest("POST", "/api/items", body)
	w := httptest.NewRecorder()

	CreateItem(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}
