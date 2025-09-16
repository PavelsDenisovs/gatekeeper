package domain

import "testing"

func TestNewFlag_ValidInputs(t *testing.T) {
	want := struct{
		envID       int
		rollout     int
		key         string
		description string
		enabled     bool
	}{
		envID: 1,
		rollout: 50,
		key: "feature x",
		description: "some description",
		enabled: true,
	}
	f, err := NewFlag(want.envID, want.rollout, want.key, want.description, want.enabled)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if f.EnvID() != want.envID {
		t.Errorf("expected EnvID=%d, got %d", want.envID, f.EnvID())
	}
	if f.Rollout() != want.rollout {
		t.Errorf("expected Rollout=%d, got %d", want.rollout, f.Rollout())
	}
	if f.Key() != want.key {
		t.Errorf("expected Key=%q, got %q", want.key, f.Key())
	}
	if f.Description() != want.description {
		t.Errorf("expected Description=%q, got %q", want.description, f.Description())
	}
	if f.Enabled() != want.enabled {
    t.Errorf("expected Enabled=%v, got %v", want.enabled, f.Enabled())
  }
	if f.CreatedAt().IsZero() {
    t.Errorf("CreatedAt should not be zero")
	}
	if f.UpdatedAt().IsZero() {
    t.Errorf("UpdatedAt should not be zero")
	}
}