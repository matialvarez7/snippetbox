package main

import (
	"testing"
	"time"

	"github.com/matialvarez7/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2025, 5, 7, 19, 15, 0, 0, time.UTC),
			want: "07 May 2025 at 19:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2025, 5, 7, 19, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "07 May 2025 at 18:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
