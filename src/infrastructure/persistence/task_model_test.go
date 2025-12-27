package persistence

import (
	"testing"
	"time"

	"clean-architecture-golang/domain/entities"
)

func TestModelRoundTrip(t *testing.T) {
	orig, err := entities.NewTask("t", "d")
	if err != nil {
		t.Fatalf("new task failed: %v", err)
	}
	// set deterministic createdAt
	now := time.Now().UTC().Truncate(time.Second)
	orig.CreatedAt = now

	m := FromDomain(orig)
	if m.ID != string(orig.ID) || m.Title != orig.Title || m.Status != orig.Status.String() {
		t.Fatalf("model mismatch after FromDomain")
	}

	d := m.ToDomain()
	if d.ID != orig.ID || d.Title != orig.Title || d.Status != orig.Status {
		t.Fatalf("domain mismatch after ToDomain")
	}
	if !d.CreatedAt.Equal(orig.CreatedAt) {
		t.Fatalf("CreatedAt mismatch")
	}
}
