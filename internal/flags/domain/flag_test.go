package domain

import (
	"errors"
	"strings"
	"testing"
)

func TestNewFlag(t *testing.T) {
	cases := []struct{
		name        string
		envID       int
		rollout     int
		key         string
		description string
		enabled     bool
		wantErr     error
	}{
		{"valid input", 1, 50, "feature x", "some description", true, nil},
		{"invalid rollout", 1, 101, "feature x", "some description", true, ErrInvalidRollout},
		{"too long key", 1, 50, strings.Repeat("x", 101), "some description", true, ErrInvalidKey},
		{"too long description", 1, 50, "feature x", strings.Repeat("x", 501), true, ErrInvalidDescription},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f, err := NewFlag(c.envID, c.rollout, c.key, c.description, c.enabled)
			if c.wantErr != nil {
				if !errors.Is(err, c.wantErr) {
					t.Errorf("expected error %v, got %v", c.wantErr, err)
				}
				return
			}

			if f.EnvID() != c.envID {
				t.Errorf("expected EnvID=%d, got %d", c.envID, f.EnvID())
			}
			if f.Rollout() != c.rollout {
				t.Errorf("expected Rollout=%d, got %d", c.rollout, f.Rollout())
			}
			if f.Key() != c.key {
				t.Errorf("expected Key=%q, got %q", c.key, f.Key())
			}
			if f.Description() != c.description {
				t.Errorf("expected Description=%q, got %q", c.description, f.Description())
			}
			if f.Enabled() != c.enabled {
  		  t.Errorf("expected Enabled=%v, got %v", c.enabled, f.Enabled())
  		}
			if f.CreatedAt().IsZero() {
  		  t.Errorf("CreatedAt should not be zero")
			}
			if f.UpdatedAt().IsZero() {
    		t.Errorf("UpdatedAt should not be zero")
			}
		})
	}
}