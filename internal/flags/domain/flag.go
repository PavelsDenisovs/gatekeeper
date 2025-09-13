package domain

import (
	"errors"
	"time"
	"unicode/utf8"
)

var ErrInvalidDescription = errors.New("description must be maximum 500 characters long")
var ErrInvalidKey =         errors.New("key must be maximum 100 characters long")
var ErrInvalidRollout =     errors.New("rollout is out of range")
var ErrIDAlreadySet =       errors.New("id already set")

type Flag struct {
	id          int
	envID       int
	key         string
	description string
	enabled     bool
	rollout     int
	createdAt   time.Time
	updatedAt   time.Time
}

func NewFlag(envID, rollout int, key, description string, enabled bool) (*Flag, error) {
	var err error
	err = validateRollout(rollout)
	if err != nil {
		return nil, err
	}

	err = validateDescription(description)
	if err != nil {
		return nil, err
	}

	err = validateKey(key)
	if err != nil {
		return nil, err
	}

	return &Flag{
		envID: envID,
		key: key,
		description: description,
		enabled: enabled,
		rollout: rollout,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func RehydrateFlag(id, envID, rollout int, key, description string, enabled bool, createdAt, updatedAt time.Time) *Flag {
	return &Flag{
			id:          id,
			envID:       envID,
			key:         key,
			description: description,
			enabled:     enabled,
			rollout:     rollout,
			createdAt:   createdAt,
			updatedAt:   updatedAt,
	}
}


// Setters
func (f *Flag) Enable() {
	f.enabled = true
	f.updatedAt = time.Now()
}

func (f *Flag) Disable() {
	f.enabled = false
	f.updatedAt = time.Now()
}

func (f *Flag) SetDescription(newDescription string) error {
	err := validateDescription(newDescription)
	if err != nil {
		return err
	}
	f.description = newDescription
	f.updatedAt = time.Now()
	return nil
}

func (f *Flag) SetKey(newKey string) error {
	err := validateKey(newKey)
	if err != nil {
		return err
	}
	f.key = newKey
	f.updatedAt = time.Now()
	return nil
}

func (f *Flag) SetRollout(newRollout int) error {
	err := validateRollout(newRollout)
	if err != nil {
		return err
	}
	f.rollout = newRollout
	f.updatedAt = time.Now()
	return nil
}

// Validators
func validateDescription(v string) error {
	if utf8.RuneCountInString(v) > 500 {
		return ErrInvalidDescription
	}
	return nil
}

func validateKey(v string) error {
	if utf8.RuneCountInString(v) > 100 {
		return ErrInvalidKey
	}
	return nil
}

func validateRollout(v int) error {
	if v > 100 || v < 0 {
		return ErrInvalidRollout
	}
	return nil
}

// Simple getters
func (f *Flag) ID() int              { return f.id }
func (f *Flag) EnvID() int           { return f.envID }
func (f *Flag) Key() string          { return f.key }
func (f *Flag) Description() string  { return f.description }
func (f *Flag) Enabled() bool        { return f.enabled }
func (f *Flag) Rollout() int         { return f.rollout }
func (f *Flag) CreatedAt() time.Time { return f.createdAt }
func (f *Flag) UpdatedAt() time.Time { return f.updatedAt }