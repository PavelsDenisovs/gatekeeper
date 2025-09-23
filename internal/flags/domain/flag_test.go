package domain

import (
	"errors"
	"strings"
	"testing"
	"time"
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
		{"valid input -> no error", 
		1, 50, "feature x", "some description", true, nil},
		{"invalid rollout -> returns ErrInvalidRollout",
		1, 101, "feature x", "some description", true, ErrInvalidRollout},
		{"too long key -> returns ErrInvalidKey",
		1, 50, strings.Repeat("x", 101), "some description", true, ErrInvalidKey},
		{"too long description -> returns ErrInvalidDescription",
		1, 50, "feature x", strings.Repeat("x", 501), true, ErrInvalidDescription},
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

func TestSetters(t *testing.T) {
	f := RehydrateFlag(1, 1, 50, "key", "desc", true, time.Now(), time.Now())

	t.Run("valid description -> updates field and UpdatedAt", func(t *testing.T) {
		assertUpdatedAtIncreased(t, f, func(t *testing.T, f *Flag) {
			checkValidSetter(t, f.SetDescription, f.Description, "description", "new decs")
		})
	})

	t.Run("invalid description -> returns ErrInvalidDescription", func(t *testing.T) {
		err := f.SetDescription(strings.Repeat("x", 501))
		if !errors.Is(err, ErrInvalidDescription) {
			t.Errorf("expected ErrInvalidDescription, got %v", err)
		}
	})

	t.Run("valid key -> updates field and UpdatedAt", func(t *testing.T) {
		assertUpdatedAtIncreased(t, f, func(t *testing.T, f *Flag) {
			checkValidSetter(t, f.SetKey, f.Key, "key", "new key")
		})
	})

	t.Run("invalid key -> returns ErrInvalidKey", func(t *testing.T) {
		err := f.SetKey(strings.Repeat("x", 101))
		if !errors.Is(err, ErrInvalidKey) {
			t.Errorf("expected ErrInvalidKey, got %v", err)
		}
	})

	t.Run("valid rollout -> updates field and UpdatedAt", func(t *testing.T) {
		assertUpdatedAtIncreased(t, f, func(t *testing.T, f *Flag) {
			checkValidSetter(t, f.SetRollout, f.Rollout, "rollout", 75)
		})
	})

	t.Run("rollout too high -> returns ErrInvalidRollout", func(t *testing.T) {
		err := f.SetRollout(101)
		if !errors.Is(err, ErrInvalidRollout) {
			t.Errorf("expected ErrInvalidRollout, got %v", err)
		}
	})

	t.Run("rollout too low -> returns ErrInvalidRollout", func(t *testing.T) {
		err := f.SetRollout(-1)
		if !errors.Is(err, ErrInvalidRollout) {
			t.Errorf("expected ErrInvalidRollout, got %v", err)
		}
	})
}

func TestEnableDisable(t *testing.T) {
	f := RehydrateFlag(1, 1, 50, "key", "desc", true, time.Now(), time.Now())
	t.Run("enable -> sets enabled=true and updates UpdatedAt", func(t *testing.T) {
		assertUpdatedAtIncreased(t, f, func(t *testing.T, f *Flag) {
			f.Enable()
			if !f.Enabled() {
				t.Errorf("expected enabled=true, got enabled=false")
			}
		})
	})
	t.Run("disable -> sets enabled=false and updates UpdatedAt", func(t *testing.T) {
		assertUpdatedAtIncreased(t, f, func(t *testing.T, f *Flag) {
			f.Disable()
			if f.Enabled() {
				t.Errorf("expected enabled=false, got enabled=true")
			}
		})
	})
}

func checkValidSetter[T comparable](
	t *testing.T,
	setter func(T) error,
	getter func() T,
	fieldName string,
	newValue T,
) {
	t.Helper()
	err := setter(newValue)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if getter() != newValue {
		t.Errorf("expected %s=%v, got %v", fieldName, newValue, getter())
	}
}

// fn must update UpdatedAt of the f
func assertUpdatedAtIncreased(t *testing.T, f *Flag, fn func(t *testing.T, f *Flag)) {
	t.Helper()
	before := f.UpdatedAt()
	time.Sleep(time.Millisecond)
	fn(t, f)
	after := f.UpdatedAt()
	if !after.After(before) {
		t.Errorf("expected UpdatedAt to increase, before=%v after=%v", before, after)
	}
}