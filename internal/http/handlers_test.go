package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/azhinu/kinklist/internal/model"
	"github.com/azhinu/kinklist/internal/storage"
)

func newTestHandler(t *testing.T) (*storage.DB, *Handler) {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "handler.db")

	db, err := storage.NewDB(path)
	if err != nil {
		t.Fatalf("failed to create test DB: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
		os.Remove(path)
	})

	return db, NewHandler(db)
}

func requestWithID(t *testing.T, method, target, id string, body io.Reader) *http.Request {
	t.Helper()

	var req *http.Request
	if body != nil {
		req = httptest.NewRequest(method, target, body)
	} else {
		req = httptest.NewRequest(method, target, nil)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id)
	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
}

func TestGetKinkList(t *testing.T) {
	db, handler := newTestHandler(t)

	now := time.Now()
	kl := &model.KinkList{
		ID:        "abc",
		Nickname:  "Tester",
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := db.SaveKinkList(kl); err != nil {
		t.Fatalf("failed to seed kink list: %v", err)
	}

	req := requestWithID(t, http.MethodGet, "/api/kinklist/abc", "abc", nil)
	rr := httptest.NewRecorder()

	handler.GetKinkList(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp model.KinkList
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.ID != kl.ID || resp.Nickname != kl.Nickname {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestGetKinkListNotFound(t *testing.T) {
	_, handler := newTestHandler(t)

	req := requestWithID(t, http.MethodGet, "/api/kinklist/missing", "missing", nil)
	rr := httptest.NewRecorder()

	handler.GetKinkList(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr.Code)
	}
}

func TestUpdateKinkList(t *testing.T) {
	db, handler := newTestHandler(t)

	body := map[string]any{
		"nickname": "Updated",
		"ratings": []map[string]string{
			{"id": "r1", "label": "Yes", "color": "green"},
		},
	}

	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}

	req := requestWithID(t, http.MethodPut, "/api/kinklist/xyz", "xyz", bytes.NewReader(raw))
	rr := httptest.NewRecorder()

	handler.UpdateKinkList(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp model.KinkList
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.ID != "xyz" {
		t.Fatalf("expected ID xyz, got %s", resp.ID)
	}

	if resp.Nickname != "Updated" {
		t.Fatalf("expected nickname Updated, got %s", resp.Nickname)
	}

	if resp.UpdatedAt.IsZero() {
		t.Fatal("expected UpdatedAt to be set")
	}

	stored, err := db.GetKinkList("xyz")
	if err != nil {
		t.Fatalf("failed to fetch saved kink list: %v", err)
	}
	if stored.Nickname != "Updated" {
		t.Fatalf("expected stored nickname Updated, got %s", stored.Nickname)
	}
}

func TestUpdateKinkListBadJSON(t *testing.T) {
	_, handler := newTestHandler(t)

	req := requestWithID(t, http.MethodPut, "/api/kinklist/xyz", "xyz", bytes.NewReader([]byte(`{"nickname":`)))
	rr := httptest.NewRecorder()

	handler.UpdateKinkList(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}
