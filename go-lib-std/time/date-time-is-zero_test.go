package datetime

import (
	"errors"
	"testing"
	"time"
)

func ValidateTime(t time.Time) error {
	if t.IsZero() {
		return errors.New("time is zero value")
	}
	return nil
}

func TestValidateTime(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		err  bool
	}{
		{
			name: "Valid time",
			t:    time.Date(2022, 10, 5, 12, 0, 0, 0, time.UTC),
			err:  false,
		},
		{
			name: "Zero time",
			t:    time.Time{},
			err:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTime(tt.t)
			if tt.err && err == nil {
				t.Errorf("expected error but got nil")
			} else if !tt.err && err != nil {
				t.Errorf("didn't expect an error but got: %v", err)
			}
		})
	}
}
