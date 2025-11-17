package storage

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/azhinu/kinklist/internal/model"
)

func newTestDB(t *testing.T) *DB {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "test.db")

	db, err := NewDB(path)
	if err != nil {
		t.Fatalf("failed to create test DB: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
		os.Remove(path)
	})

	return db
}

func TestSaveAndGetKinkList(t *testing.T) {
	db := newTestDB(t)

	now := time.Now().UTC().Truncate(time.Second)
	expected := &model.KinkList{
		ID:        "test-id",
		Nickname:  "Tester",
		CreatedAt: now,
		UpdatedAt: now,
		Ratings: []model.Rating{
			{ID: "1", Label: "Like", Color: "green"},
		},
	}

	if err := db.SaveKinkList(expected); err != nil {
		t.Fatalf("SaveKinkList error: %v", err)
	}

	got, err := db.GetKinkList(expected.ID)
	if err != nil {
		t.Fatalf("GetKinkList error: %v", err)
	}

	if got.Nickname != expected.Nickname {
		t.Fatalf("expected nickname %q, got %q", expected.Nickname, got.Nickname)
	}

	if got.CreatedAt != expected.CreatedAt || got.UpdatedAt != expected.UpdatedAt {
		t.Fatalf("timestamps mismatch: expected %v/%v got %v/%v", expected.CreatedAt, expected.UpdatedAt, got.CreatedAt, got.UpdatedAt)
	}

	if len(got.Ratings) != len(expected.Ratings) {
		t.Fatalf("expected %d ratings, got %d", len(expected.Ratings), len(got.Ratings))
	}
}

func TestGetKinkListNotFound(t *testing.T) {
	db := newTestDB(t)

	_, err := db.GetKinkList("missing")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
