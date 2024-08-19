package main

import (
	"testing"
	"time"

	"github.com/alikrugl/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC)
	hd := humanDate(tm)
	if hd != "17 Mar 2024 at 10:15" {
		t.Errorf("got %q; want %q", hd, "17 Mar 2024 at 10:15")
	}

	tests := []struct {
		name string
		time time.Time
		want string
	}{
		{
			name: "UTC",
			time: time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2024 at 10:15",
		},
		{
			name: "Empty",
			time: time.Time{},
			want: "",
		},
		{
			name: "CET",
			time: time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2024 at 09:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.time)

			assert.Equal(t, hd, tt.want)
		})
	}
}
